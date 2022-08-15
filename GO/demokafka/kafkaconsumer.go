package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("failed to start consumer,err:%v\n", err)
		return
	}

	partitionList, err := consumer.Partitions("servlog") //按照topic获取所有分区
	if err != nil {
		fmt.Printf("failed to get list of partition,err:%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { //遍历所有分区
		//每个分区建立一个对应的分区消费者
		pc, err := consumer.ConsumePartition("servlog", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()

		go func(sarama.PartitionConsumer) { //异步从每个分区获取消息
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
			}
		}(pc)
	}
	select {}
}
