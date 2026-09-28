.PHONY: all

CC ?= cc

# On big-endian architectures (s390x), pfelf only supports little-endian ELF.
# Cross-compile binaries as amd64 so unit tests can parse them.
# Use llvm-objcopy/llvm-strip for cross-architecture binary manipulation.
GOARCH_FOR_TEST ?= $(shell uname -m | sed 's/s390x/amd64/;s/x86_64/amd64/;s/aarch64/arm64/')
OBJCOPY ?= $(shell uname -m | grep -q s390x && echo llvm-objcopy-18 || echo objcopy)
# GNU strip uses --strip-section-headers; llvm-strip uses --strip-sections
STRIP_SECHDR_FLAG ?= $(shell uname -m | grep -q s390x && echo --strip-sections || echo --strip-section-headers)
STRIP ?= $(shell uname -m | grep -q s390x && echo llvm-strip-18 || echo strip)

BINARIES=fixed-address \
	go-binary \
	go-binary-no-buildinfo-section \
	go-binary-no-sections \
	without-debug-syms

all: $(BINARIES)

clean:
	rm -f $(BINARIES)

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
	$(STRIP) $(STRIP_SECHDR_FLAG) -o $@ $<
