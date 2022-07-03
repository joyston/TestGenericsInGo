package main

import "fmt"

func TryGenericValue[T comparable](s []T, x T) int {
	for _, v := range s {
		if v == x {
			return 1
		}
	}

	return -1
}

func main() {
	si := []int{1, 2, 3}
	fmt.Println(TryGenericValue(si, 3))

	ss := []string{"foo", "bar", "golang"}
	fmt.Println(TryGenericValue(ss, "not equal"))
}
