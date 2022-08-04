package main

import "fmt"

func WordsToMarks(s string) int {
	var sum int
	for i :=0; i< len(s); i++ {
		sum += (int(s[i]) - 96);
	}
	return sum
}

func main() {	
	var s string
	fmt.Scanln(&s)
	fmt.Println(WordsToMarks(s))	
}
