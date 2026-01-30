// Copyright (c) 2019 FOSS contributors of https://github.com/nxadm/tail

//go:build !linux

package tail

import "os"

// setNoAtime is a no-op on non-Linux platforms where O_NOATIME is
// not available.
func setNoAtime(_ *os.File) {}
