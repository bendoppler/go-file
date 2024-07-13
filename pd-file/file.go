package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Employee struct {
	Name       string
	Occupation string
	Year       int
}

func (e *Employee) age() int {
	now := time.Now().Year()
	return now - e.Year + 1
}

func (e *Employee) isFitWithOccupation() bool {
	len := len(e.Occupation)
	return e.Year%len == 0
}

func main() {
	f, err := os.Open("employee.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ", ")
		year, err := strconv.Atoi(strings.TrimSpace(fields[2]))
		if err != nil {
			panic(err)
		}
		newEmployee := Employee{Name: strings.TrimSpace(fields[0]), Occupation: strings.TrimSpace(fields[1]), Year: year}
		fmt.Printf("Employee name: %v, age: %v, is fit with occupation: %v %v\n", newEmployee.Name, newEmployee.age(), newEmployee.Occupation, newEmployee.isFitWithOccupation())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
