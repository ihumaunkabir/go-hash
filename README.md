# go-hash
A golang module to generate/compare hash from/with text.

### Import
```go
import "github.com/ihumaunkabir/go-hash"
```

### Usage

```go
generatedHash, err := hash.GenerateHashFromText("text")

err := hash.CompareHashAndText("hash", "text")

```

