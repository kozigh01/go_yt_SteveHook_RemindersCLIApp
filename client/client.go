package client

import "fmt"

func wrapError(customerMessage string, originalError error) error {
	return fmt.Errorf("%s: %v", customerMessage, originalError)
}