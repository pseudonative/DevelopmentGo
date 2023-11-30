docker run --name jeremy-postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres:14-alpine

docker exec -it 52c8c3cb15e0 psql -U postgres -d postgres