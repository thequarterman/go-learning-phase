package main

import (
	"fmt"
	"goroutines/channels"
)

func main() {
	ciftSayiCn := make(chan int)
	tekSayiCn := make(chan int)
	go channels.CiftSayilar(ciftSayiCn)
	go channels.TekSayilar(tekSayiCn)

	ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn
	carpim := ciftSayiToplam * tekSayiToplam
	fmt.Println("Çarpım:", carpim)
}
