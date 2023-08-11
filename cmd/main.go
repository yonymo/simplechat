package main

import (
	"fmt"
	"github.com/yonymo/simplechat/api"
)

func main() {
	fmt.Println("hello world")

	api.NewAPP("api").Run()

	return
}
