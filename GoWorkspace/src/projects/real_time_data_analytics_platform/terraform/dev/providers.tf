provider "aws" {
  alias  = "produseast1"
  region = "us-east-1"

  assume_role {
    role_arn     = "arn:aws:iam::778339273960:role/Test2FullAccess"
    session_name = "Terraform-DevUpdate"
  }
}

provider "aws" {
  alias  = "devuseast1"
  region = "us-east-1"

  assume_role {
    role_arn = "arn:aws:iam::520291287938:role/ProdFullAccess"
    session_name = "Terraform-ProdUpdate"
  }
}