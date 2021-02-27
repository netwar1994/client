package main

import (
	"context"
	"github.com/netwar1994/client/pkg/qr"
	"log"
	"time"
)

func main()  {
	qrCodeServiceUrl := "http://api.qrserver.com/v1/create-qr-code/"
	content := "https://netology.ru"
	size := 1000
	timeout := 10 * time.Second
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	svc := qr.NewService(qrCodeServiceUrl)

	data, err := svc.Encode(ctx, content, size)
	if err != nil {
		log.Println(err)
	}

	err = svc.Download(data)
	if err != nil {
		log.Println(err)
	}
}
