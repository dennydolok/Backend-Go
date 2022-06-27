package helper

import (
	"fmt"
	"time"
)

func GenerateOrderId(id, kategori string) string {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	time := fmt.Sprintf("%02d%02d%02d", now.Day(), now.Month(), now.Year())
	orderid := "INV/" + time + "/" + kategori + "/" + id
	return orderid
}
