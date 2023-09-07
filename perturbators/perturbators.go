package perturbators

import (
	"fmt"
	"structs"
)

func ApplyPert(request structs.PertRequest) string {
	message := fmt.Sprintf("Applying %v to dataset", request.Perturb)
	return message
}

// AddNoise returns a greeting for the named person.
func AddNoise(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome! I am the noise adder :)", name)
	return message
}

func Jitter(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome! I am the jitterer :)", name)
    return message
}

func Downsample(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome! I am the downsampler :)", name)
    return message
}
