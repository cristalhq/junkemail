# junkemail

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]
[![version-img]][version-url]

Fast junk email detector based on https://github.com/disposable-email-domains/disposable-email-domains data set.

## Features

* Simple API.
* Clean and tested code.
* Optimized for speed.
* Concurrent-safe.
* Dependency-free.

## Install

Go version 1.18+

```
go get github.com/cristalhq/junkemail
```

## Example

Build new token:

```go
email := "drop@10-minute-mail.com"

isJunk := junkemail.Check(email)
fmt.Printf("is '%s' junk? - %v", email, isJunk)

// Output: is 'drop@10-minute-mail.com' junk? - true
```

See examples: [example_test.go](https://github.com/cristalhq/junkemail/blob/main/example_test.go).

## Documentation

See [these docs][pkg-url].

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/junkemail/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/junkemail/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/junkemail
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/junkemail
[reportcard-img]: https://goreportcard.com/badge/cristalhq/junkemail
[reportcard-url]: https://goreportcard.com/report/cristalhq/junkemail
[coverage-img]: https://codecov.io/gh/cristalhq/junkemail/branch/main/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristalhq/junkemail
[version-img]: https://img.shields.io/github/v/release/cristalhq/junkemail
[version-url]: https://github.com/cristalhq/junkemail/releases
