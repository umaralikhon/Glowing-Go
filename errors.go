package main

import(
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main(){
	
	if len(os.Args) == 1{
		fmt.Println("Please give me one more arguments");
		os.Exit(1);
	}
	
	arguments := os.Args;
	var err error = errors.New("An error");
	k := 1;
	var n float64;
	
	for err != nil {
		if k >= len (arguments){
			fmt.Println("None of these arguments float");
			return;
		}
		
		n, err = strconv.ParseFloat(arguments[k], 64);
		k++;
	}
	
	min, max := n, n;	
}

