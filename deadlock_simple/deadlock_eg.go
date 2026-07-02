package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock1 = sync.Mutex{}
	lock2 = sync.Mutex{}
)

func blueRobot() {
	for {
		fmt.Println("Blue: Acquiring lock1...")
		lock1.Lock()
		fmt.Println("Blue: Acquiring lock2...")
		lock2.Lock()
		fmt.Println("Blue: Both locks acquired.")
		lock1.Unlock()
		lock2.Unlock()
		fmt.Println("Blue: Both locks released.")
	}
}

func redRobot() {
	for {
		fmt.Println("Red: Acquiring lock2...")
		lock1.Lock()
		fmt.Println("Red: Acquiring lock1...")
		lock2.Lock()
		fmt.Println("Red: Both locks acquired.")
		lock1.Unlock()
		lock2.Unlock()
		fmt.Println("Red: Both locks released.")
	}
}

func main() {
	go blueRobot()
	go redRobot()
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Done.")
}
