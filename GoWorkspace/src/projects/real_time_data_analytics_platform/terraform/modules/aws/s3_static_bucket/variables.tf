variable "bucket_name" {
  description = "The name of the bucket. Should be unique."
  type        = string
}

variable "bucket_acl" {
  description = "The ACL policy of the bucket."
  type        = string
  default     = "private"
}

