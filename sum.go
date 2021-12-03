package main

func main() {
	
}

func Sum(arr [5]int) int {
	sum := 0
	for i :=0; i < len(arr); i++ {
		sum += arr[i]
	}

	return sum
}
