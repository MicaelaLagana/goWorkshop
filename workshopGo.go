package main

import "fmt"

func main() {

	fmt.Println("Micaela")

	solved := reverse("Micaela")

	fmt.Println(solved)

}

func reverse(toConvert string) string {
	converted := []rune(toConvert)
	for i, j := 0, len(converted)-1; i < j; i, j = i+1, j-1 {
		converted[j] = converted[i]

	}
	return string(converted)
}
