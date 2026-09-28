//go:build s390x

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pfelf // import "go.opentelemetry.io/ebpf-profiler/libpf/pfelf"

import "debug/elf"

const CurrentMachine = elf.EM_S390
