package model

import (
	"fmt"
)

type GcpError struct {
	Err error
}

func (e *GcpError) Error() string {
	return fmt.Sprintf("GCC Storage Error: %v", e.Err)
}

type AwsError struct {
	Err error
}

func (e *AwsError) Error() string {
	return fmt.Sprintf("AWS S3 Error: %v", e.Err)
}
