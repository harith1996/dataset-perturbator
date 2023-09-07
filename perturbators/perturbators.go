package perturbators

import "fmt"

// AddNoise returns a greeting for the named person.
func AddNoise(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome! I am the noise adder :)" , name)
    return message
}