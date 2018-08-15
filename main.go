package main

import (
	"fmt"

	"github.com/Vinelab/go-base/config"
	http "github.com/Vinelab/go-base/httpclient"
)

func main() {
	fmt.Println("got here")
	fmt.Println(config.MyVariable)
	fmt.Println(http.GET)
	fmt.Println(http.POST)
}
