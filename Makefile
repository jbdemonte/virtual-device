LOCAL_IMG=virtual_device_test:latest

build:
	docker build . -t ${LOCAL_IMG}

shell:
	docker run --rm --privileged --entrypoint=/usr/bin/bash --name virtual_device_test -it ${LOCAL_IMG}
