package main

import (
	"fmt"

	"github.com/GeorgievPlamen/rss-feed/internal/config"
)

func main() {
	config := config.Read()

	fmt.Println(config)

	config.SetUser("Plamen")

	fmt.Println(config)

}
