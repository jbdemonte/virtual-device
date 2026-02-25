LOCAL_IMG=virtual_device_test:latest

build:
	docker build . -t ${LOCAL_IMG}

shell:
	docker run --rm --privileged --name virtual_device_test -it ${LOCAL_IMG} bash

test:
	go test ./... -race

test-integration: build
	docker run --rm --privileged ${LOCAL_IMG}
