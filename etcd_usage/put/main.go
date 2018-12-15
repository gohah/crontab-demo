package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	var (
		err         error
		config      clientv3.Config
		client      *clientv3.Client
		kv          clientv3.KV
		putResponse *clientv3.PutResponse
	)

	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	//New 一个k-v对象
	//context.TODO() TODO returns a non-nil, empty Context. Code should use context.TODO when
	// it's unclear which Context to use or it is not yet available
	kv = clientv3.NewKV(client)

	if putResponse, err = kv.Put(context.TODO(), "/cron/job/job1", "hello,world", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Reversion: ", putResponse.Header.Revision)
		if putResponse.PrevKv != nil {
			fmt.Println("PreValue: ", string(putResponse.PrevKv.Value))
		}
	}
}
