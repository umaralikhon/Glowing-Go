package main

import (
	"fmt"
)

func main(){
	
	sliceLiteral := []int{1,2,3,4,5};
	intSlice := make([]int, 20);
	
	for i := 0; i < len(sliceLiteral); i++ {
		fmt.Println(i, " ", sliceLiteral[i]);
	}
	
	for i := 0; i < len(intSlice); i++ {
		fmt.Print(intSlice[i]);
	}
	
	fmt.Println();	
	
	//////////////////////////////  //////////////////////////////
	
	sliceLiteral = nil;
	
	for i := 0; i < len(sliceLiteral); i++ {
		fmt.Print(sliceLiteral[i]);
	}
	
	fmt.Println();
	
	//////////////////////////////  //////////////////////////////
	
	intSlice = append(intSlice, 10, 20, 30);
	
	for i := 0; i < len(intSlice); i++ {
		fmt.Print(intSlice[i]);
	}
	
	fmt.Println();
}
