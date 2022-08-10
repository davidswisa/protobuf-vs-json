run:
	go run main.go

p:
	protoc --go_out=. *.proto

stat:
	stat --format='%s'  person.json 
