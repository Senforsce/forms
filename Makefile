build:
	@go build .

gen:
	@t1 generate

genbuild: gen build

up: genbuild
	go run .