//go:build s390x

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package maccess // import "go.opentelemetry.io/ebpf-profiler/maccess"

// CopyFromUserNoFaultIsPatched checks whether copy_from_user_nofault is patched.
//
// On s390x, nmi_uaccess_okay() always returns true (like arm64), so the
// compiler optimises away the access check and the function is always patched.
func CopyFromUserNoFaultIsPatched(_ []byte, _, _ uint64) (bool, error) {
	return true, nil
}
