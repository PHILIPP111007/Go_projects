package greeting

import (
	"errors"
	"fmt"
	"math/rand"
)

func randomFormat() string {

	var msgTypes = []string{
		"Hello, Mister %v!",
		"Nice to see you, %v!",
		"Hi, %v.",
	}
	var i int = rand.Intn(len(msgTypes))

	return msgTypes[i]
}

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Name is not provided")
	}

	msg := fmt.Sprintf(randomFormat(), name)
	return msg, nil
}
