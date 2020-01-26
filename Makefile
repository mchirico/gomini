

docker-build:
	docker build --no-cache -t gcr.io/mchirico/gomini:test -f Dockerfile .
	docker build --no-cache -t gcr.io/mchirico/gomini:pv -f Dockerfile-pv .

push:
	docker push gcr.io/mchirico/gomini:test
	docker push gcr.io/mchirico/gomini:pv

build:
	go build -v .

run:
	docker run --rm -it -p 3000:3000  gcr.io/mchirico/gomini:test
