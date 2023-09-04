package main

import (
	message "Dousheng_Backend/internal/mircoservice/message/kitex-gen/message/messageservice"
	"log"
)

func main() {
	svr := message.NewServer(new(MessageServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
