# Why

The linter is inspired by the guidance,
[Prefix Unexported Globals with `_`](https://github.com/uber-go/guide/blob/master/style.md#prefix-unexported-globals-with-_)
in the [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md).

Prefixing unexported global variables with an underscore
effectively places them in a namespace separate from normal variables,
eliminating risk of accidentally shadowing or overwriting
such a global variable.
