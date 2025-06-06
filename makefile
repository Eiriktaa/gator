DBDIR=./sql/schema/
DBURL="postgres://postgres:postgres@localhost:5432/gator"
make: 
	go build .
generate: 
	sqlc generate
up: 
	goose -dir $(DBDIR) postgres  $(DBURL) up
clean:
	-rm ./gator
	goose -dir $(DBDIR) postgres $(DBURL) down
	goose -dir $(DBDIR) postgres  $(DBURL) up
