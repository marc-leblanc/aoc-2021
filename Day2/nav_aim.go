package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var horizontal = 0
	var vertical = 0
	var aim = 0

	for scanner.Scan() {
		nav := strings.Split(scanner.Text(), " ")
		fmt.Println(nav[0])
		i, err := strconv.Atoi(nav[1])
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		if nav[0] == "forward" {
			horizontal += i
			vertical += aim * i
		} else if nav[0] == "up" {
			//vertical -= i
			aim -= i
		} else if nav[0] == "down" {
			//vertical += i
			aim += i
		}
	}

	fmt.Println("Vertical", vertical)
	fmt.Println("Horizontal", horizontal)
	fmt.Println(horizontal * vertical)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
