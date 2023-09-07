package main

import (
    "fmt"
    "perturbators"
)
func main() {
    message := perturbators.AddNoise("Martin")
    fmt.Println(message)
}