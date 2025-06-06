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

testagg: clean make
	go run . register test
	go run . addfeed "Hacker News RSS" "https://hnrss.org/newest"
	go run . addfeed "Techcrunch" "https://techcrunch.com/feed/"
	go run . addfeed "Boot.dev" "https://blog.boot.dev/index.xml"
	go run . agg 5s

