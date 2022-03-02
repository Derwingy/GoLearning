package etcd

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func EtcdBaseOperation() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, error:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

	//put
	_, err = cli.Put(context.TODO(), "Young", "Derwing")
	if err != nil {
		log.Fatal(err)
	}

	//get
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	cancel()
	resp, err := cli.Get(ctx, "Young")
	if err != nil {
		fmt.Printf("error is %v", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}
