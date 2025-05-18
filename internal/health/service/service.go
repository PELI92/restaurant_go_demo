package service

import (
	"fmt"
)

func GetPing() string {
	fmt.Printf("GetPing")
	return "pong"
}
