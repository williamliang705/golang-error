package main

import (
	"errors"
	"fmt"
)

func main() {
	if err := doSomething(); err != nil {
		fmt.Println("catch err, err:", err)
	}
	fmt.Println("done")
}

func doSomething() error {
	return errors.New("throw error in doSomething")
}
