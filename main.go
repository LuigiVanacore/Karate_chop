package Karate_chop

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

}
