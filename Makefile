
run:
	go run main.go

gen:
	go run github.com/99designs/gqlgen generate

.PHONY: gen run