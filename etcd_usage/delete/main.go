package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"time"
)

func main() {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		err     error
		kv      clientv3.KV
		delResp *clientv3.DeleteResponse
		kvPair  *mvccpb.KeyValue
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

	//删除key /cron/job/job1
	if delResp, err = kv.Delete(context.TODO(), "/cron/job/job1", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(delResp.PrevKvs)
	//获取删除前的信息
	if len(delResp.PrevKvs) > 0 {
		for _, kvPair = range delResp.PrevKvs {
			fmt.Println("删除了：", string(kvPair.Value), string(kvPair.Key))
		}
	}
}
