
default:
	go run main.go -debug < test.txt

t:
	go run main.go < test.txt

a:
	go run main.go < input.txt

ad:
	go run main.go -debug < input.txt

test:
	go test -v

help:
	@just -l
	@just --evaluate
