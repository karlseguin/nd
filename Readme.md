### Non Deterministic Helper
This small library aims to make non deterministic code easier to test.

Functions which rely on unpredicatable values are notoriously hard to test. Functions which rely on the current time are a good example. In static languages, the most common approach is to wrap the value in an interface and inject this everywhere. This is the suggestion you'll find in many books, but it simply isn't pragmatic.

However, in languages which expose functions as first class citizens, such as Go, there's a simpler solution: create your own wrapper functions.

### Installation
Install using the "go get" command:

    go get github.com/karlseguin/nd

### Usage
In your code use `nd.Now()` instead of `time.Now()`. By default, `nd.Now()` wraps `time.Now()`, so you'll get the same `time.Time` structure which represents the current time.

In tests which require a known time value, use `nd.ForceNow(time time.Time)` or `nd.ForceNowTimestamp(timestmap int64)`. After forcing the time, subsequent calls to `nd.Now()` will always return the forced value. The default behavior can be restored by calling `nd.ResetNow()`.

### Supported Generators

* `nd.Now()` - The current time
* `nd.Guidv4()` and `nd.Guidv4String()` - A V4 Guid (either as a []byte or a string)
* `nd.CryptRand(b []byte) (n, error)` - Fills b with a cryptographically secure rand
