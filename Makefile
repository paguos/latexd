PWD=$(shell pwd)

clean:
	find . -type f -name '*.log' -delete
	find . -type f -name '*.pdf' -delete

.PHONY: cli
cli: 
	cd cli && go build -o ../latexd

docker/build:
	docker build . -t paguos/latex-docker

shell: docker/build
	docker run -ti -v $(PWD):/docs paguos/latex-docker /bin/zsh

