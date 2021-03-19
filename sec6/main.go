package main

import (
	"io"
	"log"
	"os"
)
func main(){
	var fn = "data1.txt"
	if len(os.Args)> 1{
		fn = os.Args[1]
	}
	//opening and closing of files
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	io.Copy(os.Stdout, f)
}