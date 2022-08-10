run:
	go run main.go

gen:
	protoc --go_out=. pbjson/bench.proto 

test:
	go test -bench=.
