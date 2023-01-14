package main

import (
	"fmt"
	"strings"
)

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func main() {
	var data = []string{"gajah", "buaya", "ular"}
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "u")
	})
	var dataLenght5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t\t: ", data)
	// data asli : [gajah buaya ular]
	fmt.Println("filter ada huruf \"u\"\t:", dataContainsO)
	// filter ada huruf "u" : [gajah]
	fmt.Println("filter jumlah huruf \"5\"\t: ", dataLenght5)
	// filter jumlah huruf "5" : [gajah buaya]
}
