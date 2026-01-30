# Version v1.4.12
* Add `DropPageCache` config option to control OS page cache memory usage.
  When enabled, advises the kernel to release cached pages after reading,
  preventing OOM kills in memory-constrained environments (e.g. Kubernetes pods).
  - Linux/FreeBSD/NetBSD: uses `posix_fadvise(FADV_DONTNEED)` to evict read
    pages (every 64KB, at EOF, and on file close), and `FADV_SEQUENTIAL` to
    hint sequential access at file open.
  - Linux: uses `fcntl(O_NOATIME)` to suppress access-time updates, reducing
    inode writeback overhead when tailing many files.
  - macOS: uses `fcntl(F_NOCACHE)` to bypass the unified buffer cache entirely.
  - Other platforms: no-op.
* Bump minimum Go version to 1.23.
* Migrate build tags from `// +build` to `//go:build` syntax.
* Add `golang.org/x/sys` as a direct dependency for `posix_fadvise` and
  `fcntl` syscall wrappers.

# Version v1.4.11
* Bump fsnotify to v1.6.0. Should fix some issues.

# Version v1.4.9
* Bump fsnotify to v1.5.1 fixes issue #28, hpcloud/tail#90.
* PR #27: "Add timeout to tests"by @kokes++. Also timeout on FreeBSD.
* PR #29: "Use temp directory for tests, instead of relative" by @ches++.

# Version v1.4.7-v1.4.8
* Documentation updates.
* Small linter cleanups.
* Added example in test.

# Version v1.4.6

* Document the usage of Cleanup when re-reading a file (thanks to @lesovsky) for issue #18.
* Add example directories with example and tests for issues.

# Version v1.4.4-v1.4.5

* Fix of checksum problem because of forced tag. No changes to the code.

# Version v1.4.1

* Incorporated PR 162 by by Mohammed902: "Simplify non-Windows build tag".

# Version v1.4.0

* Incorporated PR 9 by mschneider82: "Added seekinfo to Tail".

# Version v1.3.1

* Incorporated PR 7: "Fix deadlock when stopping on non-empty file/buffer",
fixes upstream issue 93.


# Version v1.3.0

* Incorporated changes of unmerged upstream PR 149 by mezzi: "added line num
to Line struct".

# Version v1.2.1

* Incorporated changes of unmerged upstream PR 128 by jadekler: "Compile-able
code in readme".
* Incorporated changes of unmerged upstream PR 130 by fgeller: "small change
to comment wording".
* Incorporated changes of unmerged upstream PR 133 by sm3142: "removed
spurious newlines from log messages".

# Version v1.2.0

* Incorporated changes of unmerged upstream PR 126 by Code-Hex: "Solved the
 problem for never return the last line if it's not followed by a newline".
* Incorporated changes of unmerged upstream PR 131 by StoicPerlman: "Remove
deprecated os.SEEK consts". The changes bumped the minimal supported Go
release to 1.9.

# Version v1.1.0

* migration to go modules.
* release of master branch of the dormant upstream, because it contains
fixes and improvement no present in the tagged release.

