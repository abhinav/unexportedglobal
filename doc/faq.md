# FAQ

**What about unexported sentinel errors?**

This linter does not require you to rename sentinel error values.
That is, the following will be left alone:

```go
var errNotFound = errors.New("not found")
```

This aligns with the exception carved for this case in the style guide.

This exception exists because adding the `_` prefix to sentinel errors
would not serve any value in idiomatic Go code.
It's rare to capture a returned error in a variable not named `err`,
so there's little risk of conflict between the captured error variables
and the global sentinel error variables.
