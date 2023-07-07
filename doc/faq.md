# FAQ

**What about unexported sentinel errors?**

This linter does not require you to rename sentinel error values.
That is, the following will be left alone:

```go
var errNotFound = errors.New("not found")
```

This aligns with the exception carved for this case in the style guide.
