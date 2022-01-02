package hfuzz

// #cgo CFLAGS: -I../honggfuzz
// #cgo LDFLAGS: -L../libs -lhfuzz -lhfcommon -ldl -lrt
//
// #include <stdint.h>
//
// #include <honggfuzz.h>
//
// extern void HonggfuzzFetchData(const uint8_t** buf_ptr, size_t* len_ptr);
// extern void __cyg_profile_func_enter(uintptr_t func, uintptr_t caller);
// extern void __cyg_profile_func_exit(uintptr_t func, uintptr_t caller);
// extern void hfuzz_trace_cmp1(uintptr_t pc, uint8_t Arg1, uint8_t Arg2);
// extern void hfuzz_trace_cmp2(uintptr_t pc, uint16_t Arg1, uint16_t Arg2);
// extern void hfuzz_trace_cmp4(uintptr_t pc, uint32_t Arg1, uint32_t Arg2);
// extern void hfuzz_trace_cmp8(uintptr_t pc, uint64_t Arg1, uint64_t Arg2);
// extern void __sanitizer_cov_trace_pc_guard(uint32_t* guard);
// extern void hfuzz_trace_pc(uintptr_t pc);
// extern void hfuzzInstrumentInit();
// extern void instrumentClearNewCov();
//
// void win(uint8_t code) {
//   uint64_t ptr = code;
//   *(uint8_t*)ptr = 0;
// }
//
// uint8_t *dataPtr = NULL;
// size_t dataLen = 0;
import "C"

import (
	"fmt"
	"os"
	"unsafe"
)

type Honggfuzz struct {
}

type Callback func([]byte)

var instantiated *Honggfuzz = nil

func New() *Honggfuzz {
	if instantiated == nil {
		instantiated = new(Honggfuzz)
		C.hfuzzInstrumentInit()
	}
	return instantiated
}

func (hf *Honggfuzz) Persistent(callback Callback) {
	C.instrumentClearNewCov()
	for {
		C.HonggfuzzFetchData(&C.dataPtr, &C.dataLen)
		data := C.GoBytes(unsafe.Pointer(C.dataPtr), C.int(C.dataLen))
		callback(data)
	}
}

func (hf *Honggfuzz) TraceCmp1(pc uint, arg1 uint8, arg2 uint8) {
	C.hfuzz_trace_cmp1(C.uintptr_t(pc), C.uint8_t(arg1), C.uint8_t(arg2))
}

func (hf *Honggfuzz) TraceCmp2(pc uint, arg1 uint16, arg2 uint16) {
	C.hfuzz_trace_cmp2(C.uintptr_t(pc), C.uint16_t(arg1), C.uint16_t(arg2))
}

func (hf *Honggfuzz) TraceCmp4(pc uint, arg1 uint32, arg2 uint32) {
	C.hfuzz_trace_cmp4(C.uintptr_t(pc), C.uint32_t(arg1), C.uint32_t(arg2))
}

func (hf *Honggfuzz) TraceCmp8(pc uint, arg1 uint64, arg2 uint64) {
	C.hfuzz_trace_cmp8(C.uintptr_t(pc), C.uint64_t(arg1), C.uint64_t(arg2))
}

func (hf *Honggfuzz) TracePc(pc uint) {
	C.hfuzz_trace_pc(C.uintptr_t(pc))
}

func (hf *Honggfuzz) TraceEdge(pc uint) {
	guard := C.uint32_t(pc)
	guard %= C._HF_PC_GUARD_MAX
	C.__sanitizer_cov_trace_pc_guard(&guard)
}

func (hf *Honggfuzz) TraceFuncEnter(funk uint64, caller uint64) {
	C.__cyg_profile_func_enter(C.uintptr_t(funk), C.uintptr_t(caller))
}

func (hf *Honggfuzz) TraceFuncExit(funk uint64, caller uint64) {
	// NOTE: this is no-op as of 2022-01-01
	C.__cyg_profile_func_exit(C.uintptr_t(funk), C.uintptr_t(caller))
}

func Win(msg string) {
	fmt.Fprintln(os.Stderr, "Win: ", msg)
	C.win(0)
}

func WinWithCode(code uint8, msg string) {
	fmt.Fprintln(os.Stderr, "Win: ", msg)
	C.win(C.uint8_t(code))
}
