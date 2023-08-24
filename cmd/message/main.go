package main

import (
	"log"
	message "tiktok-backend/kitex_gen/message/messageservice"
)

func main() {
	svr := message.NewServer(new(MessageServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
