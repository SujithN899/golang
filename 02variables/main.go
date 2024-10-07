package main

import "fmt" 

const LoginToken string = "bfdbghdj" // Public variable 

func main() {
	var username string = "Sujith"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n",username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n",isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n",smallVal)

	var smallFloat float32 = 255.8763485857
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n",smallFloat)

	var smallFloat1 float64 = 255.8763485857
	fmt.Println(smallFloat1)
	fmt.Printf("Variable is of type: %T \n",smallFloat1)

	// default values and some alias 

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n",anotherVariable)

	// implicit type or way of declaring the variables, Here lexer will read the value of the variable 

	var website = "gmail.com"
	fmt.Println(website)

	// no var style

	numberOfUsers := 300000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n",LoginToken)
}


/*
Types:-
uint --> uint8, uint16, uint32, uint64
int --> int8, int16, int32, int64
float --> float32, float64
complex --> complex64 ,complex128
alias --> byte = uint8, rune = int32
*/