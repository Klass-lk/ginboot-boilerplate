.PHONY: build clean build-MyGinbootProjectFunction

build: clean
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap main.go
	mkdir -p bin/
	zip bin/myginbootproject.zip bootstrap
	rm bootstrap

build-MyGinbootProjectFunction:
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap main.go
	mkdir -p $(ARTIFACTS_DIR)
	cp bootstrap $(ARTIFACTS_DIR)/

clean:
	rm -rf bin/
	rm -f bootstrap
	rm -f myginbootproject.zip