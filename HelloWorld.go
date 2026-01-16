package main
import "fmt"

func main() {
	fmt.Println("Hello, World!")

	name, age := "Mika", 18

	fmt.Println("Name:", "I am " + name)
	fmt.Println("Age: I am ", age, " years old.")
	fmt.Println("And I am learning Go!")


	var isHappy bool = true

	if isHappy {
		fmt.Println("I am happy!")
	} else {
		fmt.Println("I am not happy.")
	}


	fmt.Println("\nDrücke 'Enter', um das Programm zu beenden...")
    fmt.Scanln()
}