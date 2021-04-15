RUN=api ls

.PHONY: up
up:
	docker-compose build
	docker-compose up

.PHONY: build
build:
	docker-compose build

.PHONY: proxy
proxy:
	docker-compose build --build-arg http_proxy=http://proxy.nagaokaut.ac.jp:8080 --build-arg https_proxy=http://proxy.nagaokaut.ac.jp:8080

.PHONY: run
run:
	docker-compose run --rm -e http_proxy=http://proxy.nagaokaut.ac.jp:8080 -e https_proxy=http://proxy.nagaokaut.ac.jp:8080 ${RUN}
