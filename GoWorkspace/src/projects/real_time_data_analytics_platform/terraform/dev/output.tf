output "bucket_id_s3" {
  description = "The name of the bucket."
  value       = module.s3_static_bucket.s3_bucket_id
}

output "bucket_arn_s3" {
  description = "The ARN of the bucket."
  value       = module.s3_static_bucket.s3_bucket_arn
}