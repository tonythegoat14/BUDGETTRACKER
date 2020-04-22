all:  build docker

build: pack
	go get
	go build ./...

pack:
	go-bindata static/ init.sql

docker:
	cd deployment && ./build.sh && cd ..

run:
	docker run --rm -p 8080:8080 -e BUDGET_PASSWORD=foo tony/budget-tracker

clean:
	rm -f budget-tracker

