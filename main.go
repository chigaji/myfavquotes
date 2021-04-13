package main

import (
	"fmt"

	"github.com/chigaji/myfavquotes/financialquotes"
	"github.com/chigaji/myfavquotes/lifequotes"
)

func main() {
	var choice int
	// financeQuote := financialquotes.Quotes()
	fmt.Println("*********************************************")
	fmt.Println("*      Q U O T E   O F   T H E   D A Y.     *")
	fmt.Println("*********************************************")
	fmt.Println("What quote would you like to recieve?")
	fmt.Printf("Select 1. Financial quote OR 2. Life quote: ")

	// choice, _ = fmt.Scan(&choice)
	fmt.Scan(&choice)

	if choice == 1 {
		financeQuote := financialquotes.Quotes()
		fmt.Println("")
		fmt.Println(financeQuote)
		fmt.Println("")
	} else if choice == 2 {
		lifeQuote := lifequotes.Quote()
		fmt.Println("")
		fmt.Println(lifeQuote)
		fmt.Println("")
	} else {
		fmt.Println("")
		fmt.Println("Invalid choice")
		fmt.Println("Enter an integer 1 or 2")
		fmt.Println("")
	}

}
