package goroutines

import (
	"fmt"
	"time"
)

func CiftSayilar() {
	for i := 0; i < 10; i = i + 2 {
		fmt.Println("Çift Sayı:", i)
		time.Sleep(1 * time.Second)
	}
}
func TekSayilar() {
	for i := 1; i < 10; i = i + 2 {
		fmt.Println("Tek Sayı:", i)
		time.Sleep(1 * time.Second)
	}
}
