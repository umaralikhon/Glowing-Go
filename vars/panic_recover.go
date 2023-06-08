package main

import (
	"fmt"
)

func a(){
	fmt.Println("Inside A method");
	defer func(){
		
		if c := recover(); c != nil{
			fmt.Println("Recover inside A method");
		}
	}
}