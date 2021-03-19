package main

import (
	"os"
	"io"

)
var (
	data = []string{
		"helllo, good morning",
		"ola, mi llamo myrachanto",
		"y estoy apprendrerando golang",
		"megusta trabaja con golang",
	}
)
func main(){
	//write to standard out(stdout)
	//fmt.Print(), fmt.Println(), and fmt.Printf()
	//write to strin, return string
	//fmt.Sprint(), fmt.Sprintln() and fmt.Sprintf()
	//write to speciafied file (io.Writer)
	//fmt.Fprint(), fmt.FprintLn(), and fmt.Fprintf()
	out, _ := os.Create("./sec7/data.txt")
	defer out.Close()

	for _, s := range data {
		io.WriteString(out,s)
		io.WriteString(out, "\n")
	}
	out = os.Stdout

	for _, s := range data {
		io.WriteString(out,s)
		io.WriteString(out, "\n")
	}
}
