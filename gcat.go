package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("--------- Help ---------")
		fmt.Println("This is Help Sample")
		fmt.Println("------------------------")
		return
	}
	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal("Not Found !!")
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
