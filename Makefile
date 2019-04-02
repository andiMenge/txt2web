binaryName := txt2web
golangContainerTag := 1.12-alpine
gitCommitId := $(shell git rev-parse head)
dockerImageName := txt2web

binary:
	@docker run --rm -v ${PWD}:/usr/src/myapp -w /usr/src/myapp -e CGO_ENABLED=0 -e GOOS=linux golang:${golangContainerTag} go build -v -a -installsuffix cgo -o ${binaryName}

docker:
	@docker build -t ${dockerImageName}:${gitCommitId} .

push:
	@docker push ${dockerImageName}:${gitCommitId}

clean:
	rm -v ./txt2web