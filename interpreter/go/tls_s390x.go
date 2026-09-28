//go:build s390x

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package golang // import "go.opentelemetry.io/ebpf-profiler/interpreter/go"

import "go.opentelemetry.io/ebpf-profiler/libpf/pfelf"

// extractTLSGOffset returns the TLS offset for the Go 'g' goroutine pointer.
//
// On s390x, the Go runtime stores the 'g' pointer in thread-local storage
// accessible via the access registers (a0:a1), not via a fixed FS-relative
// offset as on amd64. The TLS variable address is stored in the ELF TLS
// section and the offset used by the runtime is 0 relative to the TLS block.
// Return errDecodeSymbol so the caller continues with tlsOffset=0 (debug log).
func extractTLSGOffset(_ *pfelf.File) (int32, error) {
	// s390x uses access registers for g; the conventional TLS offset is 0.
	return 0, errDecodeSymbol
}
