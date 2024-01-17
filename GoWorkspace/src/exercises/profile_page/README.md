http://localhost:8080/login?username=yourname

http://localhost:8080/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InlvdXJuYW1lIiwiZXhwIjoxNzA0NDMzNTgyfQ.9TTGofkg0F1ALs53EkEvA3rpBMDWBqAunk64l0RDesk

aws eks update-kubeconfig --region us-east-1 --name dev-ops-profile

aws ecr create-repository --repository-name profile

aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 520291287938.dkr.ecr.us-east-1.amazonaws.com

docker tag profile:v1.0.1 520291287938.dkr.ecr.us-east-1.amazonaws.com/profile:v1.0.1

http://a0202ad0ac38b474f99f050f86e0a2b9-613866598.us-east-1.elb.amazonaws.com:8080/login?username=jeremy

http://rezboy.org:8080/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImtpbSIsImV4cCI6MTcwNDQ5MDQ3MX0.PgLHUJbPr6fKL6ha6ZTuy8mcX3WCP03y-sfKcLFJ7Vo


https://rezboy.org:8080/login?username=jeremy

http://ac6174ecc903a46e6b0a8aef430924df-1140790072.us-east-1.elb.amazonaws.com:8080/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImplcmVteSIsImV4cCI6MTcwNDQzMzc3Mn0.ft1VNQqDsDE8LI3JE1OQr_W-KVoB1d44iAfzlSzL8mo


