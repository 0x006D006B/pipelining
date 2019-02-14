package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	pipe1()
	fmt.Println(buildVersion)
	fmt.Println(buildTime)
	fmt.Println("Done!")
}

// Pipeline, 3 Stages
func pipe1() string {
	fmt.Println("Pipe1 - Start!")

	c1 := make(chan string, 2)
	c2 := make(chan string, 2)

	var result string

	go stage1(c1)
	go stage2(c1, c2)
	result = stage3(c2)

	fmt.Println("Pipe1 - Done!")
	//fmt.Println(result)
	return result
}

func stage1(cout chan string) {
	// Stage 1: Werte erzeugen
	for i := 0; i < 10; i++ {
		var s = strconv.Itoa(i) + "-A"
		time.Sleep(100 * time.Millisecond)
		fmt.Println("s1: " + s)
		cout <- s
	}
	close(cout)
}

func stage2(cin chan string, cout chan string) {
	for s := range cin {
		var x string
		x = s + "B"
		fmt.Println("s2: " + x)
		cout <- x
	}
	close(cout)
}

func stage3(cin chan string) string {
	var b = strings.Builder{}
	for s := range cin {
		var x string
		x = s + "C"
		fmt.Println("s3: " + x)
		b.WriteString(x)
		b.WriteString("\n")
	}
	return b.String()
}
