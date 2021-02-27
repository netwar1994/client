package main

import (
	"github.com/netwar1994/client/pkg/cbr"
	"log"
)

func main()  {
	cbrUrl := "https://raw.githubusercontent.com/netology-code/bgo-homeworks/master/10_client/assets/daily.xml"
	err := cbr.Extract(cbrUrl)
	if err != nil {
		log.Println(err)
	}
}