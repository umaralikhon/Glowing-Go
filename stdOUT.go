package main

import(
	"io"
	"os"
	_"fmt"
)

func main(){
	
	myString := "";
	arguments := os.Args;
	
	if len(arguments) == 1 {
		myString = "Please give me one argument";
	}else{
		myString = arguments[1];
	}
	
	io.WriteString(os.StdOut, myString);
	io.WriteString(os.StdOut, "\n");
}