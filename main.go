package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) > 1 {
		fileName := args[1]

		fileStat, err := os.Stat(fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if fileStat.IsDir() {
			fmt.Println("It is directory!")
			os.Exit(1)
		}

		file, err := os.Open(fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		count := 0

		for scanner.Scan() {
			count++
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("%d %s", count, fileName)

	} else {
		fmt.Println("No args!")
	}
}
