package main

import "fmt"

func main() {
    var num1, num2 float64
    var operator string

    fmt.Print("Enter the first number: ")
    fmt.Scan(&num1)
    fmt.Print("Enter the second number: ")
    fmt.Scan(&num2)
    fmt.Print("Enter the operator (+, -, *, /): ")
    fmt.Scan(&operator)

    switch operator {
    case "+":
        result := num1 + num2
        fmt.Println("The result is:", result)
    case "-":
        result := num1 - num2
        fmt.Println("The result is:", result)
    case "*":
        result := num1 * num2
        fmt.Println("The result is:", result)
    case "/":
        result := num1 / num2
        fmt.Println("The result is:", result)
    default:
        fmt.Println("Invalid operator. Please try again.")
    }
}