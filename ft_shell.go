package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		line, _, err := reader.ReadLine()
		if err != nil {
			log.Println("ReadLine Error: messsage = %v", err)
		}

		fmt.Println(string(line))
	}
}
