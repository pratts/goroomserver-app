package main

import (
	"fmt"
	"sync"

	goroomserver "github.com/pratts/goroomserver/server"
)

func main() {
	mainService := goroomserver.GetInstance()
	mainService.Init()
	var wg sync.WaitGroup
	wg.Add(1)
	go mainService.StartServer(&wg)
	wg.Wait()
	fmt.Println("hello world")
}
