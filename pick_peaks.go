package main

import (
    "fmt"
    "os"
)

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(m []int) PosPeaks {

	var res PosPeaks

	if( len(m) <= 2) {		
		return res
	}

	for i := 0; i < len(m)-2; i++ {
		s := m[i:i+3]
		if(s[0] < s[1] && s[1] > s[2]) {
			res.Pos = append(res.Pos, i+1)
			res.Peaks = append(res.Peaks, s[1])
		}
		if(s[0] < s[1] && s[1] == s[2]) {			
			for j := i+2; j < len(m) - 1; j++ {
				if (m[j+1] == s[2]) {
				continue
				} else {
					if (m[j+1] < s[2]) {
					res.Pos = append(res.Pos,i+1)
					res.Peaks = append(res.Peaks, s[1])
					}
				}
			}
		}
	}  
	return res
}

func main() {

	var a,n int
	var m []int	
	fmt.Scanln(&n)
	for i := 0; i<n; i++ {
		fmt.Fscan(os.Stdin, &a)
		m = append(m, a)	
	}	

	fmt.Println(PickPeaks(m))
}