http://localhost:8080/login?username=yourname

http://localhost:8080/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImplciIsImV4cCI6MTcwMzQ1NzU4NH0.99_RcpS0Wt531Dxlz2FkAnJh-vRtpFcufxt4n20394E


aws ecr create-repository --repository-name profile

aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 520291287938.dkr.ecr.us-east-1.amazonaws.com

docker tag profile 520291287938.dkr.ecr.us-east-1.amazonaws.com/profile

