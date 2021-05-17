package main

import "fmt"

type AppError struct {
	err error
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%v", e)
}
