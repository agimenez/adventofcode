d name="test.txt":
	go run main.go -debug < {{name}}

t:
	go run main.go < test.txt

a:
	go run main.go < input.txt

ad:
	go run main.go -debug < input.txt

test:
	go test -v

submit part answer:
	go run ../../submit.go {{part}} {{answer}}

s part answer: (submit part answer)
s1 answer: (submit "1" answer)
s2 answer: (submit "2" answer)

help:
	@just -l
	@just --evaluate
