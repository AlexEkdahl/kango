package main

import (
	"fmt"

	"github.com/AlexEkdahl/kango/config"
	"github.com/AlexEkdahl/kango/internal/client"
)

func main() {
	conf := config.New()

	_, err := client.New(*conf)
	if err != nil {
		fmt.Println("err", err)
	}
}
