# Command line usage

To use unexportedglobal from the command line,
first install the standalone program by running:

```bash
go install go.abhg.dev/unexportedglobal/cmd/unexportedglobal@latest
```

Then use it like so:

```bash
go vet -vettool=$(which unexportedglobal) ./...
```

You can also invoke it directly,
although the output can be noisier in this form.

```bash
unexportedglobal ./...
```
