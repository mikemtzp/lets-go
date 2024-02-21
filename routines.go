package main

import (
	"fmt"
	"time"
)

func routines()  {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	samurais := []string{"Joe", "Kyle", "Lee", "Choo"}

	for _, samurai := range samurais {
		go attack(samurai)
	}

	time.Sleep(time.Second*2)
}

func attack(target string) {
	fmt.Println("Attacking on", target)
	time.Sleep(time.Second)
}