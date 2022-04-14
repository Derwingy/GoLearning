package main

import (
	"context"
	"fmt"
	"log"
	"time"

	ds "golearning/godatastructure"
	pl "golearning/goproglan"

	"go.etcd.io/etcd/clientv3"
)

func EtcdBaseOperation() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, error:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

	//put
	kv := clientv3.NewKV(cli)
	res, err := kv.Put(context.TODO(), "Young", "Derwing")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Header.Revision)

	// if putResp, err := kv.Put(context.TODO(), "/school/class/students", "helios0", clientv3.WithPrevKV()); err != nil {
	//     fmt.Println(err)
	//     return
	// }
	// fmt.Println(putResp.Header.Revision)
	// if putResp.PrevKv != nil {
	//     fmt.Printf("prev Value: %s \n CreateRevision : %d \n ModRevision: %d \n Version: %d \n",
	//         string(putResp.PrevKv.Value), putResp.PrevKv.CreateRevision, putResp.PrevKv.ModRevision, putResp.PrevKv.Version)
	// }

	// resp, err = kv.Put(context.TODO(), "Young", "Derwing")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp.Header.Revision)

	//get
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
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

func main() {
	EtcdBaseOperation()
	pl.NewRequestID()
	ds.MapUsage()
}
