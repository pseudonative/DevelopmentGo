variable "s3_bucket" {
  description = "The S3 bucket for the Lambda function"
  type        = string
}

variable "aws_account_id" {
  description = "The AWS account ID"
  type        = string
}

variable "region" {
  description = "The AWS region"
  type        = string
}