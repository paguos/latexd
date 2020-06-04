PWD=$(shell pwd)
docker/build:
	docker build . -t paguos/latex-docker

shell: docker/build
	docker run -ti -v $(PWD):/docs paguos/latex-docker /bin/zsh 
