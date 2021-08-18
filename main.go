package main

import (
	"Karate_chop/chop"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var v []int

	elem, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Error in the argument prompted: %v\n", err)
		return
	}

	for _, i := range os.Args[2:] {
		temp, err := strconv.Atoi(i)
		if err != nil {
			fmt.Printf("Error in the argument prompted: %v\n", err)
			return
		}
		v = append(v, temp)
	}
	result := chop.Chop(elem, v)
	fmt.Printf("The result is %d", result)
	return
}
