package main

import (
    "fmt"
    "os"
)

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func main() {

	var a,n int
	var m []int	
	fmt.Scanln(&n)
	for i := 0; i<n; i++ {
		fmt.Fscan(os.Stdin, &a)
		m = append(m, a)	
	}
	
	var res PosPeaks

	if( n <= 2) {
		fmt.Println(res.Pos, res.Peaks)
	}

	for i := 0; i< len(m)-2; i++ {
		s := m[i:i+3]
		if(s[0] < s[1] && s[1] > s[2]) {
			res.Pos = append(res.Pos, i+1)
			res.Peaks = append(res.Peaks, s[1])
		}
	}

	fmt.Println(res.Pos, res.Peaks) 
	fmt.Println(m)
}