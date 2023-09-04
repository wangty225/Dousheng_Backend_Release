package main

import (
	"Dousheng_Backend/internal/mircoservice/user/kitex-gen/user"
	"Dousheng_Backend/internal/mircoservice/user/kitex-gen/user/userservice"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"log"
	"time"
)

func main() {
	client, err := userservice.NewClient("user", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("client : %v", client)

	for {
		req := &user.DouyinUserRequest{UserId: 225, Token: "test_token"}
		fmt.Printf("%v\n", req)

		resp, err := client.User(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second * 2)
	}
}
