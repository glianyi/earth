package watcheretcd

import (
	"context"
	"fmt"
	"testing"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestWatch(t *testing.T) {
	etcd, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://127.0.0.1:49157"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(etcd.Username, etcd.Password)

	putResp, err := etcd.Put(context.Background(), "cloud", "vsphere")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_ = putResp
	getResp, err := etcd.Get(context.Background(), "cloud")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("len: ", len(getResp.Kvs))
	for _, v := range getResp.Kvs {
		fmt.Println(string(v.Key), string(v.Value))
	}

}
