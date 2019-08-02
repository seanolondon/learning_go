package main

//github test
import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Print("Errors", err)
		os.Exit(1)
	}

	//wont resize the slice, so use a largeone
	// //99999 elements/spaces
	// //bs := make([]byte, 99999)
	// //resp is from above
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	//pass to writer interface
	//take some info into, then out of code
	//fins something that implements writer interface
	//writer type is an interface
	//first argument implements writer
	//second argument implements the reader
	//resp.Body - is a reader
	//os.Stdout is a File - write interface
	// *file has a function called write - which is why we can use it
	//therefor os.Stdout implements write interface
	lw := logWriter{}

	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
