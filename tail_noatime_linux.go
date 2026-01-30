// Copyright (c) 2019 FOSS contributors of https://github.com/nxadm/tail

//go:build linux

package tail

import (
	"os"

	"golang.org/x/sys/unix"
)

// setNoAtime sets O_NOATIME on the file descriptor to suppress access-time
// updates. This reduces inode writeback overhead when tailing many files.
// Requires file ownership or CAP_FOWNER; silently ignored on failure.
func setNoAtime(f *os.File) {
	if flags, err := unix.FcntlInt(f.Fd(), unix.F_GETFL, 0); err == nil {
		unix.FcntlInt(f.Fd(), unix.F_SETFL, flags|unix.O_NOATIME)
	}
}
