package main

import (
	"fmt"
	"sync"

	goroomserver "github.com/pratts/goroomserver/server"
)

type AppExtension struct {
	event goroomserver.Event
}

func (extension AppExtension) InitExtension(event goroomserver.Event) {
	fmt.Println("AppExtension Event:", event)
	// extension.event = event
}

type RoomExtension struct {
	event goroomserver.Event
}

func (extension *RoomExtension) InitExtension(event goroomserver.Event) {
	fmt.Println("RoomExtension Event:", event)
	// extension.event = event
}

func main() {
	mainService := goroomserver.GetInstance()
	mainService.Init()

	appExtension := AppExtension{}
	ext := &appExtension

	mainService.CreateAppService("test", ext)
	var wg sync.WaitGroup
	wg.Add(1)
	go mainService.StartServer(&wg)
	wg.Wait()
	fmt.Println("hello world")
}
