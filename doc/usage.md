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
