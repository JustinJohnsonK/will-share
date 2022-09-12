# will-share
will-share the expenses among us! Basically cloning splitwise :)

## Local setup and create migrations

### Install all the packages

    go get ./...

#### Start postgres container (optional)

    docker run --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -d postgres:latest

#### Create the database
    docker exec -it postgres createdb --username=postgres --owner=postgres willshare-db

### Run migrations

    go run cmd/migrate/main.go
    
### Run app

    go run cmd/server/main.go


##### Migrations Usage

    -num int
        number of migrations to be run (default 0)
        runs all pending migrations if num is 0
    -type string
        migrate 'up' or 'down' (default "up")
