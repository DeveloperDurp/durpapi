package service

import (
	"fmt"
	"time"
)

func RetryOperation(maxRetries int, delay time.Duration, operation func() error) error {
	var err error
	for i := 0; i <= maxRetries; i++ {
		err = operation()
		if err == nil {
			return nil
		}
		fmt.Printf("Error encountered: %v\n", err)
		if i < maxRetries {
			fmt.Printf("Retrying after %v...\n", delay)
			time.Sleep(delay)
		}
	}
	return err
}
