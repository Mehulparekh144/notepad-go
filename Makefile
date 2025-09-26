build:
	go build -o bin/main ./cmd

run:
	./bin/main $(ARGS)

br:
	go run ./cmd $(ARGS)

debug:
	dlv debug ./bin/main $(ARGS)

clean:
	rm -f bin/main
