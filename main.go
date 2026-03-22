package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"encoding/json"
)

func main() {
	// ReadFromFile()

	h := Message{Hello: "Hello World", ignored: "Yes Please ignore it for me"}
	fmt.Printf("%s\n", h.Hello)
	PrintStruct(h, "json")

	data := Employee {
		FirstName: "Abdullateef",
		LastName: "Suleiman",
		Email: "sabdullateef87@gmail.com",
		Age: 43,
		MonthlySalary: []Salary {
			{Basic: 15000.00}, {Basic: 16000.00},
		},
	}

	WriteStructToFile(data)
}

func ReadFromStandardOutput () {
	fmt.Print ("What is your name ? > ")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')

	if err != nil { // err must not contain any value else there is an exception
		log.Fatal(err)
	}

	fmt.Printf("Hello %s", line)
}


func ReadFromFile () {
	file, _ := os.OpenFile("hello.txt", os.O_RDONLY, 0666)
	defer file.Close() // this will ensure that the file is closed after the function is done executing, even if there is an error
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("> %s", line)
		if err != nil {
				return
		}
	}
}

func WriteStructToFile[T any] (data T) {
	file, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile("salary.json", file, 0644)
}

type Message struct {
	Hello string
	ignored string // this field is basically private and does not get exported when you marshall it.
}

type Salary struct {
	Basic float64
}

type Employee struct {
	FirstName, LastName, Email 	string
	Age 												int
	MonthlySalary 							[]Salary
}