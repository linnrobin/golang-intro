package main

import "fmt"

func main() {

	names := map[string]string{
		"1001": "Andi",
		"1002": "Budi",
		"1003": "Carli",
	}
	fmt.Println(names)

	mapNames := map[string]map[string]string{
		"1001": map[string]string{
			"name":    "Andi",
			"address": "Indonesia",
			"gender":  "male",
		},
		"1002": map[string]string{
			"name":    "Budi",
			"address": "Indonesia",
			"gender":  "male",
		},
		"1003": map[string]string{
			"name":    "Carli",
			"address": "Indonesia",
			"gender":  "male",
		},
	}
	fmt.Println(mapNames)
	fmt.Println(mapNames["1001"]["gender"])

	delete(mapNames, "1003")
	fmt.Println(mapNames)
}
