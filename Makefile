all:
	go run make.go -v autoDev

auto:
	go run make.go -v auto

test:
	go test ./...
	./output/velociraptor --config artifacts/testdata/windows/test.config.yaml \
	     golden artifacts/testdata/server/testcases/ --env srcDir=`pwd`

release:
	go run make.go -v release

linux:
	go run make.go -v linux

windows:
	go run make.go -v windowsDev

windowsx86:
	go run make.go -v windowsx86

clean:
	go run make.go -v clean

generate:
	go generate ./vql/windows/win32_windows.go

check:
	staticcheck ./...

build_docker:
	echo Building the initial docker container.
	docker build --tag velo_builder docker

build_release: build_docker
	echo Building release into output directory.
	docker run --rm -v `pwd`:/build/ -u `id -u`:`id -g` -e HOME=/tmp/  velo_builder

debug:
	dlv debug --build-flags="-tags 'server_vql extras'" \
		./bin/ -- frontend -v --debug

debug_client:
	dlv debug --build-flags="-tags 'server_vql extras'" \
		./bin/ -- client -v
