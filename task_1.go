/*
	Напишите Go-программу, которая вычисляет сумму всех аргументов командной
	строки, которые являются действительными числами
*/


package main

import(
	"fmt"
	"os"
	"bufio"
)

func main(){
	
	var f *os.File;
	f = os.Stdin;
	defer f.Close();
}

