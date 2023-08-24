package main

import (
	"belajar-go/doctor"
	"bufio"
	"fmt"
	"os"
)

func main() {
	dataToCapture := bufio.NewReader(os.Stdin)
	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)
	dataToPrint, _ := dataToCapture.ReadString('\n')
	fmt.Println("hi " + dataToPrint)
}
