p1:
	go build -o p1 p1.go common.go

p2:
	go build -o p2 p2.go common.go

.PHONY: run1 run2 clean

test1: p1
	./p1 <../test

test2: p2
	./p2 <../test

run1: p1
	./p1 <../input

run2: p2
	./p2 <../input

clean:
	rm -f p1 p2