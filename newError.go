package main

import (
	"errors"
	"fmt"
	"os"
)

func returnError(a, b int) error{
	
	if a == b {
		err := errors.New("Error in returnError function");
		panic(err);
		os.Exit(10);
		return err;
	}else{
		return nil;
	}					
}

func main(){
	err := returnError(1, 1);
	
	if err == nil {
		fmt.Println("returnError ended normally");
	}else{
		fmt.Println(err);
	}
}