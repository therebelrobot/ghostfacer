PWD = $(shell pwd)
# parameters:
# in: the input file, sans extension (e.g. blinky or *)
build:
	tinygo build -serial=usb -target=xiao -scheduler=tasks -o ./$(in).build.bin ./$(in).go

flash:
	tinygo flash -serial=usb -target=xiao -scheduler=tasks ./$(in).go

buildtrace:
	tinygo build -target=arduino -scheduler=coroutines -o ./$(in).build.bin ./$(in).go

flashtrace:
	tinygo flash -target=arduino -scheduler=coroutines ./$(in).go

build-docker:
	docker run --rm -v $(PWD):/src tinygo/tinygo:0.21.0 tinygo build -target=arduino -scheduler=none -o /src/$(in).build.bin /src/$(in).go

setup-docker:
	docker pull tinygo/tinygo:0.21.0

buildgo:
	go build -o ./$(in).build.bin ./$(in).go
