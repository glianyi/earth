package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {

}

func etcd() {
	etcdHost := flag.String("etcdHost", "127.0.0.1:49155", "etcd host")
	etcdWatchKey := flag.String("etcdWatchKey", "foo", "etcd key to watch")

	flag.Parse()

	fmt.Println("connecting to etcd - " + *etcdHost)

	etcd, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://" + *etcdHost},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("connected to etcd - " + *etcdHost)
	fmt.Println("watch key: ", *etcdWatchKey)

	defer etcd.Close()

	var watchChan clientv3.WatchChan
	go func() {
		fmt.Println("set WATCH on " + *etcdWatchKey)
		watchChan = etcd.Watch(context.Background(), *etcdWatchKey)
		fmt.Println("et ended...")
		fmt.Println("--------")
		for watchResp := range watchChan {
			fmt.Println("handle...")
			for _, event := range watchResp.Events {
				fmt.Printf("Event received! %s executed on %q with value %q\n", event.Type, event.Kv.Key, event.Kv.Value)
			}
		}
	}()

	// go func() {
	// 	fmt.Println("started goroutine for PUT...")
	// 	for {
	// 		etcd.Put(context.Background(), *etcdWatchKey, time.Now().String())
	// 		fmt.Println("populated " + *etcdWatchKey + " with a value..")
	// 		time.Sleep(2 * time.Second)
	// 	}

	// }()
	select {}
}

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	clientv3 "go.etcd.io/etcd/client/v3"
// )

// func main() {
// 	client, err := clientv3.New(clientv3.Config{
// 		Endpoints:   []string{"http://127.0.0.1:49155"},
// 		DialTimeout: time.Second,
// 	})

// 	fmt.Println("-----------------------")
// 	if err != nil {
// 		fmt.Println("connect failed err : ", err)
// 		return
// 	}
// 	defer client.Close()

// 	fmt.Println("connect success...")

// 	resp, err := client.Put(context.Background(), "foo", "inner")
// 	if err != nil {
// 		fmt.Println("err:", err.Error())
// 	}

// 	fmt.Println("============================")

// 	fmt.Println("revision", resp.Header.Revision)
// 	go func() {
// 		//watch
// 		fmt.Println("start watch...")
// 		watchKey := client.Watch(context.Background(), "foo")
// 		for resp := range watchKey {
// 			for _, item := range resp.Events {
// 				fmt.Printf("%s %q : %q \n", item.Type, item.Kv.Key, item.Kv.Value)
// 			}
// 		}
// 	}()

// 	fmt.Println("put key")
// 	if resp, err := client.Put(context.TODO(), "foo", "inner2"); err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(resp)
// 	}

// 	select {}
// }
