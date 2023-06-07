package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	programName := filepath.Base(os.Args[0]);
	sysLog, err := syslog.New(syslog.LOG_INFO | syslog.LOG_LOCAL7, programName);
	
	if err != nil{
		log.Fatal(err);
	}else{
		log.setOutput(sysLog);
	}
	
	log.Println("Logging in GO");
	fmt.Printlb("Will you see this?");
}