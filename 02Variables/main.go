package main

import "fmt"

//capital letter starting denotes public variable
const jwtToken = "qwert4557"

func main() {
	//stringht forward
	var userName string = "Mahesh123"
	fmt.Println(userName)
	fmt.Printf("variable of is type %T \n", userName)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable of is type %T \n", isLoggedIn)

	//implicit type
	var appName = "Golearning123"
	fmt.Println(appName)
	fmt.Printf("variable of is type %T \n", appName)

	//short declaration operator : only allowed inside any method
	userAllowedCount := 30000
	fmt.Println(userAllowedCount)
	fmt.Printf("variable of is type %T \n", userAllowedCount)

	fmt.Println(jwtToken)
	fmt.Printf("variable of is type %T \n", jwtToken)

}
