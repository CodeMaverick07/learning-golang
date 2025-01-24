package main

import "fmt"

func main() {
	var revenue, expenss, taxrate float64
	fmt.Print("enter the revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("enter the expenss: ")
	fmt.Scan(&expenss)
	fmt.Print("enter the taxrate: ")
	fmt.Scan(&taxrate)

	var ebt = revenue - expenss
	var aft = ebt * (100-taxrate)/100
	var ratio = ebt/aft

	fmt.Println("earings before tax: ",ebt)
	fmt.Println("earings after tax: ",aft)
	fmt.Println("ratio: ",ratio)


}