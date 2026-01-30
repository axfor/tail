// Copyright (c) 2019 FOSS contributors of https://github.com/nxadm/tail

//go:build linux || freebsd || netbsd

package tail

import (
	"os"

	"golang.org/x/sys/unix"
)

// initPageCacheControl hints sequential access so the kernel can
// optimise readahead and page eviction for the file.
func initPageCacheControl(f *os.File) {
	unix.Fadvise(int(f.Fd()), 0, 0, unix.FADV_SEQUENTIAL)
}

// dropPageCacheRange uses posix_fadvise to advise the kernel that the
// specified region of the file is no longer needed and can be evicted
// from the page cache.
func dropPageCacheRange(f *os.File, offset, length int64) {
	unix.Fadvise(int(f.Fd()), offset, length, unix.FADV_DONTNEED)
}
