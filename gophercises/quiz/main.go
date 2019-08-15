package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	//open the file
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	//use a csv parser then read the whole file
	reader, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Check: %s\n", string(csvread))
	//print the first line
	fmt.Println(reader[0])

	//io.Copy(os.Stdout, reader)

	//	fmt.Println("Quiz Game")
}
