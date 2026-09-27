run:
	go build -o simulator
	./simulator

test:
	go test -v ./...

clean:
	rm -f simulator
