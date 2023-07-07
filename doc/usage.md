# Usage

To use unexportedglobal from the command line,
first install the standalone program by running:

```bash
go install go.abhg.dev/unexportedglobal/cmd/unexportedglobal@latest
```

Then use it like so:

```bash
unexportedglobal ./...
```

Use it with `go vet` for cleaner output.

```bash
go vet -vettool=$(which unexportedglobal) ./...
```

It is also usable as a `golangci-lint` plugin:
 ```bash
 $ go build -buildmode=plugin go.abhg.dev/unexportedglobal/cmd/unexportedglobal
 $ cat .golangci.yml
 linter-settings:
   custom:
     unexportedglobal:
       path: unexportedglobal.so
       description: Verify unexported globals have an underscore prefix.
       original-url: go.abhg.dev/unexportedglobal
 ```
