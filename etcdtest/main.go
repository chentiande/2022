package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
		Username:    "root",
		Password:    "root",
	})
	if err != nil {
		fmt.Println("open ", err)
		// handle error!
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	resp, err := cli.Get(ctx, "chentiande")
	if err != nil {

	} else {

		fmt.Println(string(resp.Kvs[0].Value))
	}

	cancel()

	if err != nil {
		// handle error!
	}
	// use the response
}
