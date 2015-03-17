package affinity

import "unsafe"
import "syscall"

func SetAffinity(pid int, mask *int64) {
	syscall.Syscall(syscall.SYS_SCHED_SETAFFINITY,
		uintptr(pid), 8, uintptr(unsafe.Pointer(mask)))
}

func GetAffinity(pid int, mask *int64) {
	syscall.Syscall(syscall.SYS_SCHED_GETAFFINITY,
		uintptr(pid), 8, uintptr(unsafe.Pointer(mask)))
}
