package main

//github test
import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, errp := http.Get("http://google.com")
	if err != nil {
		fmt.Print("Errors", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}
