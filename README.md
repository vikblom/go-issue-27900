# go-issue-27900

This module is an example of where
```
> go list -m all
go-issue-27900
github.com/google/go-cmp v0.6.0
github.com/yuin/goldmark v1.4.13
golang.org/x/mod v0.19.0
golang.org/x/net v0.27.0
golang.org/x/sync v0.7.0
golang.org/x/sys v0.22.0
golang.org/x/telemetry v0.0.0-20240521205824-bda55230c457
golang.org/x/tools v0.23.0
```

contains a module (`github.com/yuin/goldmark`) which is not needed for the output of
```
> go list all | grep github.com/yuin/goldmark
[empty]
```
and is not explained by
```
> go mod why -m github.com/yuin/goldmark
# github.com/yuin/goldmark
(main module does not need module github.com/yuin/goldmark)
```

Output from `go version go1.22.5 linux/amd64`.

## References

Upstream [issue](https://github.com/golang/go/issues/27900).

Go list [docs](https://pkg.go.dev/cmd/go#hdr-Package_lists_and_patterns).
