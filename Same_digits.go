
package main

import (
    "fmt"
    "math"
)

func number_of_digits (x int) int {
	n := 1
	for n < 5 {     
        if ( int(float64(x) / math.Pow(10, float64(n))) == 0){           
            break
        } else {
        n++       
        }
	}
    return n
}

func main() {
    var x, y int
    fmt.Scan(&x,&y) 
    n := number_of_digits(x)
  
	for n > 0 {
		number_n := int(float64(x) / math.Pow(10, float64(n-1))) % 10		
		n--	
		m := number_of_digits(y)
		for m > 0 {
			number_m := int(float64(y) / math.Pow(10, float64(m-1))) % 10				
			m--
			if ( number_n == number_m) {
				fmt.Print(number_m, " ")				
			} else {
				continue
			}			
		}		
	}
	fmt.Println()
}