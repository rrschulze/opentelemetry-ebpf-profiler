//go:build linux && s390x

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package process // import "go.opentelemetry.io/ebpf-profiler/process"

import (
	"debug/elf"
	"encoding/binary"
	"fmt"
)

// GetMachineData returns s390x machine data. No PAC masks on s390x.
func (sp *ptraceProcess) GetMachineData() MachineData {
	return MachineData{Machine: elf.EM_S390}
}

// getThreadInfo reads the general-purpose registers for a given thread on s390x
// using ptrace NT_PRSTATUS. The register block (psw + 16 gprs) starts at
// offset 72 in the elf_prstatus structure.
func (sp *ptraceProcess) getThreadInfo(tid int) (ThreadInfo, error) {
	// s390x elf_prstatus: fixed header (72 bytes) + psw (16 bytes) + gprs (16×8=128 bytes)
	prStatus := make([]byte, 72+16+128)
	if err := ptraceGetRegset(tid, int(elf.NT_PRSTATUS), prStatus); err != nil {
		return ThreadInfo{}, fmt.Errorf("failed to get LWP %d thread info: %v", tid, err)
	}
	return ThreadInfo{
		LWP:    uint32(tid),
		GPRegs: prStatus,
		// s390x TLS base: read from a separate regset; return 0 as non-critical fallback.
		TPBase: binary.BigEndian.Uint64(prStatus[72+8+15*8:]),
	}, nil
}
