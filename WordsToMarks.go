package main

import "fmt"

func WordsToMarks(s string) int {
	var sum int
	for _,v := range(s) {
		sum += (int(v) - 96);
	}
	return sum
}

func main() {	
	var s string
	fmt.Scanln(&s)
	fmt.Println(WordsToMarks(s))	
}
