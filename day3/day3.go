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

func part1(inputFilename string) {

	m := make(map[int][]int)
	input := readFile(inputFilename)
	for i := range input {
		bin := input[i]
		bin_list := strings.Split(bin, "")
		fmt.Println(bin_list)
		for j := range bin_list {
			val, _ := strconv.Atoi(bin_list[j])
			// fmt.Println()
			// m_map := m[j]
			// m_map[val] = m_map[val] + 1
			// m[j] = m_map

			m_list := m[j]
			m_list = append(m_list, val)
			m[j] = m_list
		}
	}
	fmt.Println(m)
}

func main() {

	part1("test_input.txt")
}
