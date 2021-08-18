package main

import (
	"fmt"
	"os"
	"strconv"
)

func chop(elem int, v []int) (index int) {
	for len(v) > 0 {
		middle := (len(v) - 1) / 2

		switch {
		case v[middle] < elem:
			v = v[middle+1:]
			index += middle + 1
		case v[middle] > elem:
			v = v[:middle]
		default:
			return middle + index
		}
	}
	return -1
}

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
	result := chop(elem, v)
	fmt.Printf("The result is %d", result)
	return
}
