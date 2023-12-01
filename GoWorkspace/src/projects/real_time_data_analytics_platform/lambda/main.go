package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context) (string, error) {
	env := getEnv()
	return fmt.Sprintf("Hello, World! Environment: %s", env), nil
}

func getEnv() string {
	// Read the environment name from an environment variable
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		return "unknown" // Default value if not set
	}
	return env
}

func main() {
	lambda.Start(handleRequest)
}
