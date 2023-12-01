#!/bin/bash

# Stop executing after any error
set -e

# Define variables
BUCKET_NAME="dev-data-dev"  # Replace with your actual S3 bucket name
FUNCTION_NAME="real_time_lambda_function"  # Replace with your actual Lambda function name
ZIP_FILE="function.zip"  # Name for the zip file
BUILD_OUTPUT="main"  # The output binary name

# Define the Lambda function directory relative to this script
LAMBDA_DIR="./lambda"  # This is the directory where your Go Lambda function code is located

# Check if the LAMBDA_DIR exists
if [ ! -d "$LAMBDA_DIR" ]; then
    echo "Lambda directory not found: $LAMBDA_DIR"
    exit 1
fi

# Navigate to the Lambda function directory
cd "$LAMBDA_DIR"


go get github.com/aws/aws-lambda-go/lambda

# Fetch dependencies
go mod tidy

# Build the Go Lambda function for Linux
GOOS=linux go build -o "$BUILD_OUTPUT"

# Zip the compiled binary
zip "$ZIP_FILE" "$BUILD_OUTPUT"

# Upload the zip file to the specified S3 bucket
aws s3 cp "$ZIP_FILE" "s3://$BUCKET_NAME/$FUNCTION_NAME/$ZIP_FILE"

# Clean up the build files
rm -f  "$ZIP_FILE"

# Navigate back to the project root
cd -

echo "Lambda function built and uploaded to S3 successfully."
