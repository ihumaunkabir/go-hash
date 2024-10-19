## go-hash
A golang module to generate/compare hash from/with text.

#### Installation
```go
go get github.com/ihumaunkabir/go-hash
```

#### Import
```go
import "github.com/ihumaunkabir/go-hash"
```

#### Usage
Generate hash from plain text
```go
generatedHash, err := hash.GenerateHashFromText("text")
```
Compare hash and plain text
```go
err := hash.CompareHashAndText("hash", "text")
```

#### License
MIT


