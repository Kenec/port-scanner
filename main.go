package main

import (
	"fmt"
	"github.com/Kenec/port-scanner/port"
)

func main(){
	fmt.Println("Port Scanner in Go")

	open := port.ScanPort("tcp", "localhost", 22)
	fmt.Printf("Port Open: %t\n", open)

	port.InitialScan("localhost")
	//fmt.Println(results)

}