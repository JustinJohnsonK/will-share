# will-share
will-share the expenses among us! Basically cloning splitwise :)

docker run --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -d postgres:latest

docker exec -it postgres createdb --username=postgres --owner=postgres willshare-db

go run cmd/migrate/main.go