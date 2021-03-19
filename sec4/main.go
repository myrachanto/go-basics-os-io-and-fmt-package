package main

import (
	"io"
	"log"
	"os"
)
func main(){
	// var file *os.File
	// fmt.Printf("file: value = %v, type =%T\n", file, file)
	f, err := os.Create("data1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString("mira quiero tener un trabajo - trabajando con golang")
	io.WriteString(f, "yeah another string ok!\n")
}