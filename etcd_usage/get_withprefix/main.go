package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		err     error
		kv      clientv3.KV
		getResp *clientv3.GetResponse
	)
	//客户端配置
	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	//建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	//New 一个k-v对象
	//context.TODO() TODO returns a non-nil, empty Context. Code should use context.TODO when
	// it's unclear which Context to use or it is not yet available
	kv = clientv3.NewKV(client)

	//读取/cron/job/为前缀的所有key
	if getResp, err = kv.Get(context.TODO(), "/cron/job/", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(getResp.Kvs)
	}
}
