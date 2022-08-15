package logtransfer

import (
	"context"
	"encoding/json"
	"fmt"
	"logtransfer/conf"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/olivere/elastic/v7"
	"go.etcd.io/etcd/client/v3"
	"gopkg.in/ini.v1"
)

func LoadIniConf(filepath string) (cfg *conf.LogTransConf, err error) {
	var cfgObj conf.LogTransConf
	err = ini.MapTo(&cfgObj, filepath)
	if err != nil {
		fmt.Printf("加载配置文件<%s>出错,err:%v\n", err)
		return nil, err
	}
	fmt.Printf("cfgObj:%v\n", cfgObj)

	return &cfgObj, nil
}

func GetKafkaConsumer(addrs []string) (consumer sarama.Consumer, err error) {
	consumer, err = sarama.NewConsumer(addrs, nil)
	if err != nil {
		fmt.Printf("failed to start consumer,err:%v\n", err)
		return nil, err
	}

	return consumer, nil
}

func GetESCli(url string, u string, p string) (esCli *elastic.Client, err error) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	fmt.Printf("u:%s p:%s\n", u, p)
	EsCli, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetBasicAuth(u, p),
	)
	if err != nil {
		fmt.Printf("创建ES连接客户端失败,err:%v\n", err)
		return nil, err
	}
	fmt.Println("连接ES成功!")
	return EsCli, nil
}

type LogBody struct {
	data string
}

func SendMsg2ES(esCli *elastic.Client, ESIndex string, msgBody string) error {
	// p1 := Person{Name: "shangsan", Age: 100}
	msg := LogBody{data: msgBody}
	fmt.Printf("send msg:%v\n", msg)
	put1, err := esCli.Index().Index(ESIndex).BodyJson(msg).Do(context.Background())
	if err != nil {
		fmt.Printf("往ES put数据出错,err:\n", err)
		return err
	}
	fmt.Printf("index %s,type %s\n", put1.Index, put1.Type)
	return nil
}

//统筹kafka消费者goroutine
func ArrangeConsumer(kafkaConsumer sarama.Consumer, esCli *elastic.Client, topicItems []string) (err error) {
	for _, topic := range topicItems {
		partitions, err := kafkaConsumer.Partitions(topic)
		if err != nil {
			fmt.Printf("从kafka获取topic<%s>分区信息出错,err:%v\n", topic, err)
			continue
		}
		fmt.Printf("get topic:%s\n", topic)
		for part := range partitions {
			fmt.Printf("topic: %s, part:%d\n", topic, part)
			pc, err := kafkaConsumer.ConsumePartition(topic, int32(part), sarama.OffsetNewest)
			if err != nil {
				fmt.Printf("消费topic<%s>分区%d 出错,err:%v\n", topic, part, err)
				continue
			}
			defer pc.AsyncClose()
			go func(sarama.PartitionConsumer) { //gotourine 异步消费分区数据
				for msg := range pc.Messages() {
					fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
					SendMsg2ES(esCli, "logmonitor", string(msg.Value))
				}
			}(pc)
		}
	}
	select {}
	return nil
}

func GetEtcdClient(addr string, timeoutSec int) (cli *clientv3.Client) {
	etcdcli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: time.Duration(timeoutSec) * time.Second,
	})

	if err != nil {
		fmt.Printf("初始化连接etcd出错,err:%v\n", err)
		return
	}
	// defer etcdcli.Close()
	return etcdcli
}

type Tailfileinfo struct {
	Path  string
	Topic string
}

func GetKeyFromEtcd(cli *clientv3.Client, key string) (items []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	etcdResp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("从etcd get键值<key:%s>出错,err:%v\n", key, err)
		return
	}

	var val []*Tailfileinfo
	for _, kv := range etcdResp.Kvs {
		fmt.Printf("========%s:%s\n", kv.Key, kv.Value)
		err := json.Unmarshal(kv.Value, &val)
		if err != nil {
			fmt.Printf("解析json出错,err:%v\n", err)
			return
		}
	}

	for _, key := range val {
		items = append(items, fmt.Sprintf("%s", key.Topic))
	}

	return items
}

func Start() error {
	// 1.加载配置
	cfg, err := LoadIniConf("./conf/config.ini")
	if err != nil {
		return err
	}
	// 2.创建消费kafka的消费者
	consumer, err := GetKafkaConsumer([]string{cfg.KafkaInfo.Address})
	if err != nil {
		return err
	}
	//3.连接ES
	esCli, err := GetESCli(cfg.ESInfo.Address, cfg.ESInfo.User, cfg.ESInfo.Password)
	if err != nil {
		return err
	}

	etcdCli := GetEtcdClient(cfg.EtcdInfo.Address, cfg.EtcdInfo.TimeoutSec)
	items := GetKeyFromEtcd(etcdCli, cfg.EtcdInfo.Key)

	ArrangeConsumer(consumer, esCli, items)
	return nil
}
