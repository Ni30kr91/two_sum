package main

import "fmt"

func main() {

	num := []int{2, 7, 11, 15}

	target := 18
	for i := 0; i < len(num); i++ {
		for j := 0; j < len(num); j++ {
			if num[i]+num[j] == target {
				fmt.Println(i, j)
				break

			}
		}
	}

}
