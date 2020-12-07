# gcr.io/septapig/
docker-build:
	docker build --no-cache -t gcr.io/septapig/go-delete-gmail:test -f Dockerfile .

push:
	docker push gcr.io/septapig/go-delete-gmail:test

build:
	go build -v .

run:
	docker run --rm -it -p 3000:3000  gcr.io/septapig/go-delete-gmail:test
