package main

import (
	"fmt"
	"os"
	"bufio"
	//"strings"
)

func main() {
	fmt.Println("")
	fmt.Println("################### WELCOME TO SYSTEM METAOS #########")
	fmt.Println("|----------------------------------------------------|")
	fmt.Println("|    Enter command or exit to finish the execution:  |")
	fmt.Println("|----------------------------------------------------|")
	fmt.Println("")
	for {
		fmt.Print(">> ")
		com := bufio.NewScanner(os.Stdin)
		if com.Scan() {
			if com.Text() == "exit" {
				break
			} else {
				//strace(strings.Fields(com.Text()))
			}
		}
	}
}
