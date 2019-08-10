package main

//github test
import (
	"io"
	"log"
	"os"
)

//type txtReader struct{}

func main() {
	//file, err := os.Open("myfile.txt.txt") // For read access.
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// data := make([]byte, 100)
	// count, err := file.Read(data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%q\n", data[:count])
	//fmt.Println(file)

	// // fmt.Println(data)

	// tx := txtReader{}

	// //file.txtReader
	io.Copy(os.Stdout, file)

}

/////this section almost worked but not quite
// func (txtReader) Read(bs []byte) (int, error) {
// 	fmt.Println(string(bs))
// 	fmt.Println("Just wrote this many bytes:", string(bs))
// 	return len(bs), nil
// }
