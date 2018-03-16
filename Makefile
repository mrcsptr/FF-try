fight:
	go build -i -v -o bin/fight ./cmd/fight/

clean:
	ls bin/* | grep -v json | xargs rm
