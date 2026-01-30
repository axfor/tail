[![Go Reference](https://pkg.go.dev/badge/github.com/nxadm/tail.svg)](https://pkg.go.dev/github.com/nxadm/tail#section-documentation)
![ci](https://github.com/nxadm/tail/workflows/ci/badge.svg)
[![FreeBSD](https://api.cirrus-ci.com/github/nxadm/tail.svg)](https://cirrus-ci.com/github/nxadm/tail)
# tail functionality in Go

nxadm/tail provides a Go library that emulates the features of the BSD `tail`
program. The library comes with full support for truncation/move detection as
it is designed to work with log rotation tools. The library works on all
operating systems supported by Go, including POSIX systems like Linux, *BSD,
MacOS, and MS Windows. Go 1.12 is the oldest compiler release supported.

A simple example:

```Go
// Create a tail
t, err := tail.TailFile(
	"/var/log/nginx.log", tail.Config{Follow: true, ReOpen: true})
if err != nil {
    panic(err)
}

// Print the text of each received line
for line := range t.Lines {
    fmt.Println(line.Text)
}
```

See [API documentation](https://pkg.go.dev/github.com/nxadm/tail#section-documentation).

## Installing

    go get github.com/nxadm/tail/...

## History

This project is an active, drop-in replacement for the
[abandoned](https://en.wikipedia.org/wiki/HPE_Helion) Go tail library at
[hpcloud](https://github.com/hpcloud/tail). Next to
[addressing open issues/PRs of the original project](https://github.com/nxadm/tail/issues/6),
nxadm/tail continues the development by keeping up to date with the Go toolchain
(e.g. go modules) and dependencies, completing the documentation, adding features
and fixing bugs.

## Examples
Examples, e.g. used to debug an issue, are kept in the [examples directory](/examples).

## Page Cache Control

When tailing many files in memory-constrained environments (e.g. Kubernetes pods),
the OS page cache can grow unbounded and cause OOM kills. The `DropPageCache` option
advises the kernel to release cached pages after reading, keeping memory usage stable.

```Go
t, err := tail.TailFile("/var/log/app.log", tail.Config{
    Follow:        true,
    ReOpen:        true,
    DropPageCache: true,
})
```

When enabled, the following optimizations are applied:

| Platform | Mechanism | Effect |
|----------|-----------|--------|
| Linux | `fadvise(FADV_SEQUENTIAL)` | Hints sequential access for better readahead and eviction |
| Linux | `fcntl(O_NOATIME)` | Suppresses access-time updates to reduce inode writeback |
| Linux, FreeBSD, NetBSD | `fadvise(FADV_DONTNEED)` | Evicts read pages from the page cache (every 64KB, at EOF, and on close) |
| macOS | `fcntl(F_NOCACHE)` | Bypasses the unified buffer cache entirely |
