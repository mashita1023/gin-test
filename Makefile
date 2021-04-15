RUN = api ls
GET = none

.PHONY: up
up:
	docker-compose build
	docker-compose up

.PHONY: build
build:
	docker-compose build

.PHONY: run
run:
	docker-compose run --rm ${RUN}

.PHONY: tidy
tidy:
	docker-compose run --rm api go mod tidy

.PHONY: get
get:
	docker-composerun --rm api go get ${GET}
