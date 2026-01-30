// Copyright (c) 2019 FOSS contributors of https://github.com/nxadm/tail

//go:build darwin

package tail

import (
	"os"
	"syscall"
)

// initPageCacheControl disables the unified buffer cache for this file
// descriptor via F_NOCACHE. All subsequent reads bypass the page cache
// entirely, which is the macOS equivalent of Linux FADV_DONTNEED.
func initPageCacheControl(f *os.File) {
	syscall.Syscall(syscall.SYS_FCNTL, f.Fd(), syscall.F_NOCACHE, 1)
}

// dropPageCacheRange is a no-op on macOS because F_NOCACHE already
// prevents data from entering the page cache.
func dropPageCacheRange(_ *os.File, _, _ int64) {}
