ifeq ($(mode),debug)
	LDFLAGS="-X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`' -X main.GIT_HASH=`git rev-parse HEAD`"
else
	LDFLAGS="-s -w -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`' -X main.GIT_HASH=`git rev-parse HEAD`"
endif

.PHONY: build
build:
	go build -ldflags ${LDFLAGS} -o build/nftkit main.go

clean:
	rm -rf ./build