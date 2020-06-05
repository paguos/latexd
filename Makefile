PWD=$(shell pwd)

.PHONY: cli
cli: 
	cd cli && go build -o ../latexd

build:
	docker build . -t latexd

shell: build
	docker run -ti -v $(PWD):/docs latexd /bin/zsh

