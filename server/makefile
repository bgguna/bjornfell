# Project
current_dir	=	${PWD}
bin_dir		=	${current_dir}/bin
out_dir		=	${current_dir}/out
binary		=	BjornFell Server
project		=	bjornfell
main		=	main.go
coverage	=	coverage.out

all: clean format build test

# Go Commands
gobuild		=	go build
gorun		=	go run
goget		=	go goget
gofmt		=	go fmt
golint		=	golint
gotest		=	go test
gocover		=	go tool cover
ldflags		=	"-s -w"
modules		=	./...

clean:
	@echo "Cleaning project repository..."
	rm -rf ${bin_dir}
	rm -rf ${out_dir}
	rm -rf ${coverage}
	@echo "Done cleaning."

format:
	@ echo "Formatting code..."
	@ $(gofmt) $(modules)
	@ echo "Done formatting."

lint:
	@ echo "Linting the code to Go standards..."
	@ $(golint) $(modules)
	@ echo "Done linting."

run-dev:
	@ echo "Running in dev mode..."
	$(gorun) ${main}

build:
	@echo "Building binaries..."
	$(gobuild) $(modules)
	$(gobuild) -o ${bin_dir}/${binary} ./
	@echo "Done building binaries."

run:
	@echo "Running binary..."
	${bin_dir}/${binary}

test:
	@ echo "Running tests with package level coverage..."
	$(gotest) $(modules) -v -cover
	@echo "Done running tests."

coverage:
	@ echo "Running tests and printing coverage report..."
	@ -$(gotest) $(modules) -race -v -covermode=atomic -coverprofile $(coverage)
	@ $(gocover) -func $(coverage)
	@ echo "Done writing test coverage report."

help:
	@echo "\
	Manual for BjornFell Server\n\
	\n\
	Available make commands:\n\
	------------------------------\n\
	* all:		formats, cleans, builds and tests the binaries\n\
	* build:	builds binaries\n\
	* run:		runs the server\n\
	* format:	formats the code to the Golang standards\n\
	* lint:		highlights smells in the code\n\
	* test:		runs the tests with coverage for all packages in the project\n\
	* coverage	runs the tests and writes the coverage to file\n\
	* clean:	deletes the generated binaries and output files (log and health report)\n\
	* help:		prints this menu again\n\
	\n\
	Done."
