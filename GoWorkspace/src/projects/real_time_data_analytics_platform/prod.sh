#!/bin/bash

# Stop executing after any error
set -e

# Define variables
LAMBDA_DIR="./lambda"  # Your Lambda function directory
BUCKET_NAME="prod-data-prod"  # Your S3 bucket name in the production account
FUNCTION_NAME="real_time_lambda_function"  # Your Lambda function's name
BUILD_OUTPUT="main"  # The output binary name
ZIP_FILE="function.zip"  # The zip file name
AWS_PROFILE="test2"  # Your named profile for the production account

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

# Upload the zip file to the specified S3 bucket using the given profile
aws s3 cp "$ZIP_FILE" "s3://$BUCKET_NAME/$FUNCTION_NAME/$ZIP_FILE" --profile "$AWS_PROFILE"

# Clean up the build files
# rm -f "$BUILD_OUTPUT" "$ZIP_FILE"
rm -f "$ZIP_FILE"

# Navigate back to the project root
cd -

echo "Lambda function for production built and uploaded to S3 successfully."


