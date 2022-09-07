package main

// Sum calculates the total from an array of numbers.
func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}

/*func main() {
	sumArr := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(sumArr))
}*/
