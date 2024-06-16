LOCAL_IMG=virtual-input-test:latest

build:
	docker build . -t ${LOCAL_IMG}

shell:
	docker run --rm --privileged --entrypoint=/usr/bin/bash --name virtual-input-test -it ${LOCAL_IMG}
