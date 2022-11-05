package main

import "fmt"

func main() {
	for i := 1; i <= 50; i++ {
		res := ""
		if i % 3 == 0 {
			res += "fizz"
		}
		if i % 5 == 0 {
			res += "buzz"
		}
		if res == "" {
			res = fmt.Sprint(i)
		}
		fmt.Println(res)
	}	
}

