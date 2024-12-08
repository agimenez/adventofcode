
default:
	go run main.go -debug < test.txt

test:
	go run main.go < test.txt

a:
	go run main.go < input.txt

help:
	@just -l
	@just --evaluate
