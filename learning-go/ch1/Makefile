# Each operation is called target.
# .DEFAULT_GOAL defines which target is run for default.
.DEFAULT_GOAL := build 

# keeps make out of confusion if exists a directory with same name as one of targets
.PHONY:fmt vet build
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

# the 'vet' means: Build target needs vet target first.
build: vet
	go build

clean:
	go clean
