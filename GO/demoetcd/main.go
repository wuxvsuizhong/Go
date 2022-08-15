package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/client/v3"
)

func main() {
	//启动一个etcd的连接
	etcdcli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Printf("初始化连接etcd出错,err:%v\n", err)
		return
	}
	defer etcdcli.Close()

	//启用一个context并设置超时时间是1秒钟
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//向etcd中put键值对
	_, err = etcdcli.Put(ctx, "testkey", "testval hello...")
	_, err = etcdcli.Put(ctx, "/logagent/log-collect/", `[{"path":"./my.log","topic":"testkey"},{"path":"./db.log","topic":"dblog"},{"path":"./serv.log","topic":"servlog"}]`)
	// _, err = etcdcli.Put(ctx, "/logagent/log-collect/", `[{"path":"./my.log","topic":"testkey"},{"path":"./serv.log","topic":"servlog"}]`)
	cancel() //手动触发cancel更完整
	if err != nil {
		fmt.Printf("向etcd put出错,err:%v\n", err)
		return
	}

	//从etcd中get键值对
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	etcdResp, err := etcdcli.Get(ctx, "testkey")
	cancel()
	if err != nil {
		fmt.Printf("从etcd get出错,err:%v\n", err)
		return
	}

	for _, ev := range etcdResp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}

}

func main1() {
	//启动一个etcd的连接
	etcdcli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Printf("初始化连接etcd出错,err:%v\n", err)
		return
	}
	defer etcdcli.Close()

	//watch监控键testkey的值的变化
	ch := etcdcli.Watch(context.Background(), "testkey")
	for resp := range ch {
		for _, event := range resp.Events {
			fmt.Printf("时间类型:%v key:%s value:%s\n", event.Type, event.Kv.Key, event.Kv.Value)
		}
	}
}
