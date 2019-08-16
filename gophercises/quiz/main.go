package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

// type results struct {
// 	answer string
// }

func main() {
	//open the file, declare type, name and "help". The flag is only a POINTER TO A STRING
	csvFilename := flag.String("csv", "problems.csv", "a csv in the format of 'question,answer'")
	//use the flag.Parse() function
	flag.Parse()
	//just so code compiles but nothing really happens
	//_ = csvFilename
	//use the asterix for the pointer to a string
	file, err := os.Open(*csvFilename)
	if err != nil {
		//the %s means use a string, so we'll use the pointer to the string of filename. \n is newline character
		//fmt.Printf("Failed to open the CSV file: %s\n", *csvFilename)
		//use the exit function to print the resulting string from the Sprintf
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
		//csv.NewReader take in an io.Reader which is what file is.

		//os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to open the CSV file.")
	}

	//problems are the result of the parseLines function on lines
	//a set of structs {} inside a slice [], so the problems is a slice
	problems := parseLine(lines)

	//start a counter
	correct := 0

	for i, p := range problems {
		//%d means a number
		fmt.Printf("Problem #%d: %s =\n", i+1, p.q)
		var answer string
		//SCanf won't work on long strings - it gets rid of spaces
		//&answer is a reference to answer, as a pointer variable so we can access it
		fmt.Scanf("%s\n", &answer)
		//if answer from user euqal answer brought in from the csv 'p.a'
		if answer == p.a {
			//add 1 to the correct counter
			correct++
		}
	}

	//RESEARCH PRINTF VS PRINTLN
	fmt.Printf("You scored correct %d out of %d.\n", correct, len(problems))
	//code compiles temporarily
	//_ = file
}

//USE := TO DEFINE A NEW
//NOTE OUTSIDE OF MAIN FUNC, AND 2D STRING [][]string --- research this
//a function with variable lines wihch is a 2d string, returns a []problem struct
func parseLine(lines [][]string) []problem {
	//make a problem slice going by the number of lines
	//use the append function method when you dont know the sizes of things, her ewe do know the size
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
			//use the trim space function in case there are spaces in the csv which will mean our answers will never be right

		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

//this function create the error message string, so it's reusable elsewhere. THen Exit
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// 	var saveResults results

// 	file, err := os.Open(os.Args[1])
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//use a csv parser then read the whole file
// 	reader, err := csv.NewReader(file).ReadAll()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//for loop through each row

// 	//change below to a function which adds 1 each reponse.
// 	//

// 	input := bufio.NewReader(os.Stdin)

// 	for probs := range reader {
// 		fmt.Printf("Problem #%v: %v =\n", probs+1, reader[probs][0])
// 		var CommandLine = NewFlagSet(os.Args[0], ExitOnError)
// 		if (reader[probs][0] == CommandLine)

// 		// saveResults = append(saveResults, results{
// 		// 	saveResults.answer = input,
// 		// }
// 	}

// 	//fmt.Println(reader[0])
// }

// func Test(saveResults []results) []int {
// 	slice = append(slice, 100)

// 	fmt.Println(slice)

// 	return slice
// }

// a = Test(a)
