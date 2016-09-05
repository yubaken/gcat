package main

import (
	"fmt"
	"github.com/yubaken/gcat/cmd"
	"os"
)

func main() {
	if err := cmd.Root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	os.Exit(0)
}
