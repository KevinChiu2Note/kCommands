BUILD_DIR=bin/

init:
	mkdir -p "${BUILD_DIR}"

kcat:
	cd kcat ; \
	go build ; \
	mv kcat ../bin

.PHONY: init kcat