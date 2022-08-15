package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

//使用sarama 的kafka client
func main() {
	config := sarama.NewConfig()                              //生成配置
	config.Producer.RequiredAcks = sarama.WaitForAll          //设置选项--发送数据需要kafka的leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //设置选项--随机选出一个partition
	config.Producer.Return.Successes = true                   //选项--成功交付的数据消息会在success channel返回

	//构建一个消息（消息数据是一个结构体）
	msg := &sarama.ProducerMessage{}
	msg.Topic = "test_news"
	msg.Value = sarama.StringEncoder("这是一条测试消息----------------")

	//连接fakfa
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Printf("连接kafka失败,err:%v\n", err)
		return
	}
	defer client.Close()

	//向kafka发数据
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Printf("向kafka发送消息失败,err:%v\n", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
