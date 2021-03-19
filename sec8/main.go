package main

import (
	"fmt"

)
func main(){
	// formated input
	//read from standard in (stdin)
	//fmt.Scan(), fmt.Scanln(), and fmt.Scanf()
	// Read from string
	//fmt.Sscan(), fmt.Sscanln(), and fmt.Sscanf()
	// read from specified file (io.Reader)
	//fmt.Fscan(), fmt.Fscanln(), and fmt.Fscanf()
	input := "Soy tu deuna 35 001-11-2333 5.6"

	var i int 
	var f float64
	var s0, s1, s2, s3 string
	fmt.Sscan(input, &s0, &s1,&s2, &i, &s3, &f)
	fmt.Printf("got int i: %v\n", i)
	fmt.Printf("got  float f: %v\n", f)
	fmt.Printf("got string s0: %v\n", s0)
	fmt.Printf("got string s1: %v\n", s1)
	fmt.Printf("got string s2: %v\n", s2)
	fmt.Printf("got string s3: %v\n", s3)
	///try to convert float to string
	s := fmt.Sprint(78.12466)
	fmt.Println(s)
}