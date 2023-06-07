package main

import (
	"errors"
	"fmt"
)

func returnError(a, b int) error{
	
	if a == b {
		err := errors.New("Error in returnError function");
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