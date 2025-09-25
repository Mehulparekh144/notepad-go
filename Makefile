build:
	go build -o bin/main ./cmd

run:
	./bin/main $(ARGS)

br:
	go run ./cmd $(ARGS)

clean:
	rm -f bin/main
