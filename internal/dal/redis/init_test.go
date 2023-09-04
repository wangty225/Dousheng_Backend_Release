package redis

import (
	"context"
	"fmt"
	"testing"
)

func Connect() {
	pong, err := DBRedis.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)
}

func TestConnect(t *testing.T) {
	Connect()
}
