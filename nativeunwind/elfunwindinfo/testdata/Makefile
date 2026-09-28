.PHONY: all

BINARIES=helloworld \
	helloworld.pie \
	helloworld.stripped.pie \
	helloworld.arm64 \
	helloworld.linkexternal \
	helloworld.linkexternal.stripped

# Use the default go executable if it is not specified otherwise.
GO_BINARY ?= go

all: $(BINARIES)

clean:
	rm -f $(BINARIES)

# On big-endian architectures (s390x), pfelf only supports little-endian ELF.
# Cross-compile helloworld* as amd64 so unit tests can parse them.
GOARCH_FOR_TEST ?= $(shell uname -m | sed 's/s390x/amd64/;s/x86_64/amd64/;s/aarch64/arm64/')

helloworld:
	GOARCH=$(GOARCH_FOR_TEST) CGO_ENABLED=0 $(GO_BINARY) build -o $@ helloworld.go

helloworld.pie:
	GOARCH=$(GOARCH_FOR_TEST) CGO_ENABLED=1 $(GO_BINARY) build -buildmode=pie -o $@ helloworld.go

helloworld.stripped.pie:
	GOARCH=$(GOARCH_FOR_TEST) CGO_ENABLED=1 $(GO_BINARY) build -buildmode=pie -ldflags="-s -w" -o $@ helloworld.go

helloworld.arm64:
	GOARCH=arm64 $(GO_BINARY) build -o $@ helloworld.go

helloworld.linkexternal:
	GOARCH=$(GOARCH_FOR_TEST) CGO_ENABLED=1 GOTOOLCHAIN=go1.26.7 $(GO_BINARY) build -ldflags="-linkmode=external" -o $@ helloworld.go

helloworld.linkexternal.stripped: helloworld.linkexternal
	objcopy -S $< $@
