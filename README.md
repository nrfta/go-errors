# go-errors ![](https://github.com/neighborly/go-errors/workflows/CI/badge.svg)

Error handling for Go. It uses [pkg/errors](https://github.com/pkg/errors)
under the hood.

## Installation

```sh
go get github.com/neighborly/go-errors
```

## Usage

### New error

Create a new error. It uses [pkg/errors](https://github.com/pkg/errors) for the
actual error.

```go
err := errors.New("an error message")
```

### Wrapping errors

The `errors.Wrap` function returns a new error that adds context to the original error. For example

```go
_, err := ioutil.ReadAll(r)
if err != nil {
        return errors.Wrap(err, "read failed")
}
```

### Custom Error Code

We implement a custom error type/code for pre-defined errors. These are useful
for application code to return a correct HTTP status or GraphQL error.

| Type               | String            |
|--------------------|-------------------|
| `InternalError`    | Internal Error    |
| `NotFound`         | Not Found         |
| `InvalidArgument`  | Invalid Argument  |
| `Unauthenticated`  | Unauthenticated   |
| `PermissionDenied` | Permission Denied |
| `Unknown`          | Unknown           |

You can create an error using the following:

```go
err := errors.PermissionDenied.New("user does not have access")
```

or you can wrap an existing as follows:

```go
err = errors.PermissionDenied.Wrap(err)
```

### WithDisplayMessage

Use this function to a display message to an error. Usually, error messages
are meant to be used internally only, instead of displaying them to users.

You can use `errors.WithDisplayMessage` to assign a display message to a given
error.

```go
err := errors.New("internal message")
err = errors.WithDisplayMessage(err, "Display message goes here!")
```

### Retrieving display messages

You can retrieve a display message by using `errors.DisplayMessage`.

If no message were assigned, it would use the error code string.

```go
err := errors.New("internal message")
err = errors.WithDisplayMessage(err, "Display message goes here!")
errors.DisplayMessage(err) // -> Display message goes here!
```

```go
err := errors.New("an error")
errors.DisplayMessage(err) // -> Internal Error
```

```go
err := errors.NotFound.New("an error")
errors.DisplayMessage(err) // -> Not Found
```

### Retrieving the cause of an error

Behaves the same as [pkg/errors](https://github.com/pkg/errors).

Note that if you create an error using this package, it returns the
underlying `pkg/error` error.

`errors.Cause`  recursively retrieves the topmost error which does not
implement `causer` which is assumed to be the original cause. For example:

```go
switch err := errors.Cause(err).(type) {
case *MyError
  // handle specifically
default:
  // unknown error
}
```

## License

This project is licensed under the [MIT License](LICENSE.md).
