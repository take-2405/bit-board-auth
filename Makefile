GOCMD=go
DOCKERCMD=docker
DOCKERCOMPOSECMD=docker-compose
GO_RUN=$(GOCMD) run
GO_BUILD=$(GOCMD) build
DOCKER_BUILD=$(DOCKERCMD) build
DOCKER_RUN=$(DOCKERCMD) run

all:
clean:
	rm server
build:
	$(GO_BUILD) -o server cmd/main.go
run:
	$(GO_RUN) cmd/main.go
compose-up:
	$(DOCKERCOMPOSECMD) -f ./build/docker-compose.yml up -d
compose-down:
	$(DOCKERCOMPOSECMD) -f ./build/docker-compose.yml down
docker-build:
	$(DOCKER_BUILD) ./ -t take2405/bit-board-auth:0.1.0
docker-run:
	$(DOCKER_RUN) -d -p 8080:8080 take2405/bit-board-auth:0.1.0