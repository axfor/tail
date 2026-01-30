// Copyright (c) 2019 FOSS contributors of https://github.com/nxadm/tail

//go:build !linux && !freebsd && !netbsd && !darwin

package tail

import "os"

func initPageCacheControl(_ *os.File) {}

// dropPageCacheRange is a no-op on platforms that do not support
// posix_fadvise (e.g. Windows).
func dropPageCacheRange(_ *os.File, _, _ int64) {}
