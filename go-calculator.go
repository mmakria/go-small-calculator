package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	//fmt.Print("Enter Revenue: ")
	//fmt.Scan(&revenue)
	//fmt.Print("Enter Expenses: ")
	//fmt.Scan(&expenses)
	//fmt.Print("Enter Tax Rate: ")
	//fmt.Scan(&taxRate)

	//ebt := revenue - expenses
	//profit := ebt * (1 - taxRate/100)
	//ratio := ebt / profit
	revenue, expenses, taxRate = inputOutputFlow(revenue, expenses, taxRate)
	ebt, profit, ratio := calculation(revenue, expenses, taxRate)
	printValues(ebt, profit, ratio)
	//fmt.Println(ebt)
	//fmt.Print(profit)
	//fmt.Println(ratio)

}

func inputOutputFlow(revenue, expenses, taxRate float64) (float64, float64, float64) {
	fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter Tax Rate: ")
	fmt.Scan(&taxRate)
	return revenue, expenses, taxRate
}

func calculation(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func printValues(ebt, profit, ratio float64) {
	fmt.Println(ebt)
	fmt.Print(profit)
	fmt.Println(ratio)
}
