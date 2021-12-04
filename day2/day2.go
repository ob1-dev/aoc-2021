package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(filename string) []string {
	dat, err := os.Open(filename)
	check(err)

	var input []string
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func part2() {
	instructions_list := readFile("input1.txt")
	horizontal := 0
	depth := 0
	aim := 0
	for i := range instructions_list {
		instruction := instructions_list[i]
		instructions := strings.Split(instruction, " ")
		direction := instructions[0]
		val, err := strconv.Atoi(instructions[1])
		check(err)
		fmt.Println(direction, val)
		switch direction {
		case "forward":
			fmt.Println("Going Forward")
			horizontal += val
			depth += aim * val
		case "up":
			fmt.Println("Going Up")
			// vertical -= val
			aim -= val
		case "down":
			fmt.Println("Going Down")
			// vertical += val
			aim += val
		}
	}
	fmt.Println(horizontal, depth)
	fmt.Println(horizontal * depth)
}

func part1() {
	instructions_list := readFile("input1.txt")
	horizontal := 0
	vertical := 0
	for i := range instructions_list {
		instruction := instructions_list[i]
		instructions := strings.Split(instruction, " ")
		direction := instructions[0]
		val, err := strconv.Atoi(instructions[1])
		check(err)
		fmt.Println(direction, val)
		switch direction {
		case "forward":
			fmt.Println("Going Forward")
			horizontal += val
		case "up":
			fmt.Println("Going Up")
			vertical -= val
		case "down":
			fmt.Println("Going Down")
			vertical += val
		}
	}
	fmt.Println(horizontal, vertical)
	fmt.Println(horizontal * vertical)
}

func main() {

	part2()
}
