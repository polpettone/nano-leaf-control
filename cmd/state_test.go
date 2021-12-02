package cmd

import (
	"fmt"
	"testing"
)

func Test_onStateBody(t *testing.T) {
	jsonValue := stateBody(true)
	fmt.Printf("r: %s", string(jsonValue))
	jsonValue = stateBody(false)
	fmt.Printf("r: %s", string(jsonValue))
}
