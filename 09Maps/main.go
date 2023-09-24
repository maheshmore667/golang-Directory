package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps playgroud")
	languagesMap := make(map[string]string)

	languagesMap["Js"] = "Javascript"
	languagesMap["Rb"] = "Ruby"
	languagesMap["Py"] = "Python"

	fmt.Println("Languages map is as follows : ", languagesMap)
	for key, value := range languagesMap {
		fmt.Println("Key : ", key, "value : ", value)
	}

	answer := languagesMap["Pb"]
	fmt.Println(answer)
}
