package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {
	fmt.Println("hello, Go SDK!")
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("------------Error Founded------------")
		fmt.Println(err)
	}

	fmt.Println("------------Config Loaded------------")
	fmt.Println(cfg)
	fmt.Println("------------------------------------")

	cfg2, err2 := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile("myProfile"),
	)
	if err2 != nil {
		panic(fmt.Sprintf("failed loading config, %v", err2))
	}

	fmt.Println("------------Profile Config Loaded------------")
	fmt.Println(cfg2)

}
