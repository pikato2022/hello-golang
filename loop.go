package main

import "fmt"

func main() {
	i:=1
	for i <=3 {
		i++
		fmt.Println(i)
	}
	for j:=0; j <= 6;j++ {
		fmt.Println(j)
	}
	for { //while true
		fmt.Println("loop")
		break
	}
}