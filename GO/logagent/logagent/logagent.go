package logagent

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/hpcloud/tail"
	"go.etcd.io/etcd/client/v3"
	"gopkg.in/ini.v1"
)

func GetKafkaClient(addrs []string) (client sarama.SyncProducer) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	//连接fakfa
	cli, err := sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Printf("连接kafka失败,err:%v\n", err)
		panic(err)
	}
	return cli
}

func SendMsg2Kafka(client sarama.SyncProducer, topic string, chars string) {
	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(chars),
	}
	fmt.Printf("向topic-%s发送消息:%s\n", topic, chars)
	_, _, err := client.SendMessage(&msg)
	if err != nil {
		fmt.Printf("向topic-%s发送消息<%s>出错,err:%v\n", topic, chars, err)
		//尝试重新连接kafka获取clien
		return
	}
	return
}

//返回一个通道
func GetFileTail(fileName string) (lines <-chan *tail.Line) {
	conf := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	tailObj, err := tail.TailFile(fileName, conf)
	if err != nil {
		fmt.Printf("TailFile 出错,err:%v\n", err)
		return
	}
	//按行读取日志
	// line, ok := <-tails.Lines
	return tailObj.Lines

}

//监控文件并把文件中刷新的新内容发送到kafka中
func TailFileWithKafka(kafkacli sarama.SyncProducer, fpath string, topic string, ctx *context.Context, wg *sync.WaitGroup) {
	var datachan = GetFileTail(fpath)
	defer wg.Done()
	for {
		select {
		case msg := <-datachan:
			SendMsg2Kafka(kafkacli, topic, msg.Text)
		case <-(*ctx).Done():
			fmt.Printf("收到结束消息,topic:%s task:%s\n", topic, fpath)
			return
		default:
			time.Sleep(time.Second)
		}
	}
}

type Kafkainfo struct {
	Address string `ini:"address"`
}

type Tailfileinfo struct {
	Path  string
	Topic string
}

//统筹tailfile 任务的结构体对象
type TailFileRoutine struct {
	cancelSignal context.CancelFunc
	tailfileinfo *Tailfileinfo
}

type EtcdInfo struct {
	Address    string `ini:"address"`
	TimeoutSec int    `ini:"timeout"`
	Key        string `ini:"key"`
	// /logagent/log-collect/:[{"path":"./my.log","topic":"testkey"},{"path":"./db.log","topic":"testkey"},{"path":"./serv.log","topic":"testkey"}]
}

type IniInfo struct {
	Kafkainfo    `ini:"kafka"`
	Tailfileinfo `ini:"tailfile"`
	EtcdInfo     `ini:"etcd"`
}

func LoadIniConf(confpath string) (cfg *IniInfo) {
	cfg = new(IniInfo)
	err := ini.MapTo(cfg, confpath)
	if err != nil {
		fmt.Printf("加载配置<%s>失败,err:%v\n", confpath, err)
		return
	}

	fmt.Printf("%#v\n", cfg)

	return cfg
}

func GetEtcdClient(addr string, timeoutSec int) (cli *clientv3.Client) {
	etcdcli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: time.Duration(timeoutSec) * time.Second,
	})

	if err != nil {
		fmt.Printf("初始化连接etcd出错,err:%v\n", err)
		panic(err)
	}
	// defer etcdcli.Close()
	return etcdcli
}

func GetKeyFromEtcd(cli *clientv3.Client, key string) (val []*Tailfileinfo) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	etcdResp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("从etcd get键值<key:%s>出错,err:%v\n", key, err)
		return
	}

	for _, kv := range etcdResp.Kvs {
		fmt.Printf("========%s:%s\n", kv.Key, kv.Value)
		err := json.Unmarshal(kv.Value, &val)
		if err != nil {
			fmt.Printf("解析json出错,err:%v\n", err)
			return
		}
	}

	return val
}

//统筹安排日志监控的goroutine
func ArrangeFileMoni(kafkacli sarama.SyncProducer, items []*Tailfileinfo, wg *sync.WaitGroup, routineRec *map[string]*TailFileRoutine) {
	//遍历历史表，检索每个key是否在最新的items中
	//1.历史记录有key但是最新的key中没有该取消历史任务
	//2.历史有key但是最新的items中也有一样的key，不更新操作
	//3.历史无key,但是最新的itms中有key，则新增task

	//遍历记录
	var iteksKey = []string{}
	for _, key := range items {
		iteksKey = append(iteksKey, fmt.Sprintf("%s:%s", key.Path, key.Topic))
	}

	for k, _ := range *routineRec {
		var delflag bool = true
		for _, ik := range iteksKey {
			if k == ik {
				fmt.Printf("%s == %s\n", k, ik)
				delflag = false
				break
			}
		}
		if delflag {
			(*routineRec)[k].cancelSignal()
			delete(*routineRec, k)
			//	从记录中删除
			fmt.Printf("after delete, record has:%v.%v\n", *routineRec, (*routineRec))
		}
	}
	//遍历最新的配置列表和历史记录对比,找出其中有的但是历史中没有的即是要新增的
	for _, v := range items {
		key := fmt.Sprintf("%s:%s", v.Path, v.Topic)
		//历史记录不能够查到key
		_, ok := (*routineRec)[key]
		if !ok {
			//历史无key,但是最新的itms中有key，则新增task
			ctx, cancel := context.WithCancel(context.Background())
			//根据配置条目循环启动goroutine去监控日志
			wg.Add(1)
			go TailFileWithKafka(kafkacli, v.Path, v.Topic, &ctx, wg)
			//记录中新增任务条目
			(*routineRec)[key] = &TailFileRoutine{
				tailfileinfo: v,
				cancelSignal: cancel,
			}
			fmt.Printf("after add,routineRec:%v\n", *routineRec)
		}
	}
}

func Start() {
	cfg := LoadIniConf("./conf/config.ini")
	//初始化kafka连接,得到客户端
	var client sarama.SyncProducer = GetKafkaClient([]string{cfg.Kafkainfo.Address})
	//获取etcd客户端
	var etcdcli = GetEtcdClient(cfg.EtcdInfo.Address, cfg.EtcdInfo.TimeoutSec)
	//根据key从etcd中获取配置值
	var items = GetKeyFromEtcd(etcdcli, cfg.EtcdInfo.Key)
	for i, v := range items {
		fmt.Printf("index:%v val.path:%v val.topic:%v\n", i, v.Path, v.Topic)
	}
	var wg sync.WaitGroup
	var routineRec = make(map[string]*TailFileRoutine, 1)
	//根据获取的配置条目去启动goroutine监控日志
	ArrangeFileMoni(client, items, &wg, &routineRec)

	//启用watch监控etcd中的配置变化
	ch := etcdcli.Watch(context.Background(), cfg.EtcdInfo.Key)
	for resp := range ch {
		for _, event := range resp.Events {
			fmt.Printf("时间类型:%v key:%s value:%s\n", event.Type, event.Kv.Key, event.Kv.Value)
			var val = []*Tailfileinfo{}
			json.Unmarshal(event.Kv.Value, &val)
			ArrangeFileMoni(client, val, &wg, &routineRec)
		}
	}

	wg.Wait()
}
