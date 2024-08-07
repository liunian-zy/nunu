package main

import (
	"fmt"
	"github.com/liunian-zy/nunu/cmd/nunu"
)

func main() {
	err := nunu.Execute()
	if err != nil {
		fmt.Println("execute error: ", err.Error())
	}
}
