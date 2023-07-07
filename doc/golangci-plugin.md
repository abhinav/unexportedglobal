# Use as a golangci-lint plugin

To use unexportedglobal as a golangci-lint plugin,
take the following steps:

- Clone the repository or download a source archive
  from the Releases page.

    ```bash
    git clone https://github.com/abhinav/unexportedglobal.git
    ```

- Build the plugin:

    ```bash
    cd unexportedglobal
    go build -buildmode=plugin ./cmd/unexportedglobal
    ```

- Add the linter under `linter-settings.custom` in your `.golangci.yml`,
  referring to the compiled plugin (usually called 'unexportedglobal.so').

    ```yaml
    linter-settings:
      custom:
        unexportedglobal:
          path: unexportedglobal.so
          description: Verify unexported globals have an underscore prefix.
          original-url: go.abhg.dev/unexportedglobal
    ```

- Run golangci-lint as usual.

> **Warning**:
> For this to work, your plugin must be built for the same environment
> as the golangci-lint binary you're using.
>
> See [How to add a private linter to golangci-lint][1] for details.

[1]: https://golangci-lint.run/contributing/new-linters/#how-to-add-a-private-linter-to-golangci-lint
