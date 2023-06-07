package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	
	var f *os.File;
	f= os.Stdin;
	defer f.Close();
	
	fmt.Print("Write here: ");
	scanner := bufio.NewScanner(f);
	
	for scanner.Scan(){
		fmt.Println("> ", scanner.Text());
	}
}

