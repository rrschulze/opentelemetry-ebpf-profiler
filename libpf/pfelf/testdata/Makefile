.PHONY: all

CC ?= cc
OBJCOPY ?= objcopy
STRIP ?= strip

BINARIES=fixed-address \
	go-binary \
	go-binary-no-buildinfo-section \
	go-binary-no-sections \
	without-debug-syms

all: $(BINARIES)

clean:
	rm -f $(BINARIES)

# On big-endian architectures (s390x), pfelf only supports little-endian ELF.
# Cross-compile binaries as amd64 so unit tests can parse them.
GOARCH_FOR_TEST ?= $(shell uname -m | sed 's/s390x/amd64/;s/x86_64/amd64/;s/aarch64/arm64/')

without-debug-syms: test.c
	$(CC) $< -s -o $@

fixed-address: fixed-address.c fixed-address.ld
	# The following command will likely print a warning (about a missing -T option), which should be ignored.
	# Removing the warning would require passing a fully-fledged linker script to bypass gcc's default.
	$(CC) $^ -o $@

go-binary: gotest.go
	GOARCH=$(GOARCH_FOR_TEST) CGO_ENABLED=0 go build -o go-binary -ldflags "-w -s" gotest.go

go-binary-no-buildinfo-section: go-binary
	$(OBJCOPY) --rename-section .go.buildinfo=.go.buildinfo.hidden $< $@

go-binary-no-sections: go-binary
	$(STRIP) --strip-section-headers -o $@ $<
