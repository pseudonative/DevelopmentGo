data "aws_caller_identity" "current" {}

terraform {
  required_version = ">= 1.5.0"
  backend "s3" {
    bucket         = "shoshone-tfstate"
    key            = "shoshonekey/prodRealTimeDataAnalytics"
    region         = "us-east-1"
    kms_key_id     = "arn:aws:kms:us-east-1:520291287938:key/4fc9e509-04c4-4881-89e7-46fb49790093"
    dynamodb_table = "shoshone-state-lock"
  }
}

module "s3_static_bucket" {
  source = "../modules/aws/s3_static_bucket"
  providers = {
    aws = aws.produseast1
  }
  bucket_acl = "private"
  bucket_name = "prod-data"
}

module "lambda_function" {
  source = "../modules/aws/lambda_function"
  providers = {
    aws = aws.produseast1
  }
  s3_bucket = module.s3_static_bucket.s3_bucket_id
  aws_account_id = data.aws_caller_identity.current.account_id
  region     = "us-east-1"
}