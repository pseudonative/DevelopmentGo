data "aws_caller_identity" "current" {}

module "lambda" {
  source  = "terraform-aws-modules/lambda/aws"
  version = "6.5.0"

  source_path = "../../lambda"

  function_name = "real_time_lambda_function"
  handler       = "main"  # Update with your handler path
  runtime       = "go1.x"         # Specify the runtime environment
  s3_bucket     = var.s3_bucket   # Ensure this is correctly set to the S3 bucket name
  
  create_role = true

  # Define the assume role policy for the Lambda function's IAM role
  assume_role_policy_statements = {
    account_root = {
      effect    = "Allow"
      actions   = ["sts:AssumeRole"]
      principals = {
        account_principal = {
          type        = "AWS"
          identifiers = ["arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"]
        }
      }
    }
  }

  # Inline policy directly attached to the Lambda function's IAM role
  attach_policy_json = true
  policy_json = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action   = ["xray:GetSamplingStatisticSummaries"]
        Resource = "*"
      }
    ]
  })

  # Additional policies to attach to the Lambda function's IAM role
  attach_policies    = true
  policies           = ["arn:aws:iam::aws:policy/AWSXrayReadOnlyAccess"]
  number_of_policies = 1

  # Environment variables for the Lambda function
  environment_variables = {
    ENVIRONMENT = terraform.workspace
  }

  # Custom IAM policy statements
  attach_policy_statements = true
  policy_statements = {
    dynamodb = {
      effect    = "Allow"
      actions   = ["dynamodb:BatchWriteItem"]
      resources = ["arn:aws:dynamodb:${var.region}:${data.aws_caller_identity.current.account_id}:table/Test"]
    }
    s3_read = {
      effect    = "Deny"
      actions   = ["s3:HeadObject", "s3:GetObject"]
      resources = ["arn:aws:s3:::my-bucket/*"]
    }
  }

  timeouts = {
    create = "20m"
    update = "20m"
    delete = "20m"
  }

}
