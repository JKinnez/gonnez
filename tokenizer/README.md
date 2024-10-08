## Tokenizer

Package to manage tokenization using [Paseto](https://github.com/o1egl/paseto) under the hood. 

### Config

```
  tokenizer.New(tokenizer.Config{})
```
- Audience
- Footer
- Issuer
- Location: set a location for the token expiracy time. Defaults to `UTC` example: `Europe/Madrid`
- PublicKey
- PrivateKey
- SymetricKey

### Result 

```
type ReaderResult struct {
	paseto.JSONToken
	Footer string
}
```

### Usage

#### Generate keys

Symmetric key
```
  key, err := tokenizer.GenerateSymmetricKey()
```

Public and private key
```
  publicKey, privateKey, err := tokenizer.GenerateKeyPair()
```

#### Basic Usage

```
package main

import (
  "fmt"
  "github.com/gonnez/tokenizer"
)

func main() {
  tokenizer = tokenizer.New(tokenizer.Options{
    SymmetricKey: "your-key"
  })
  subject := "0x3"
  durationInHours := tokenizer.Durations.OneDay
  bearerToken, err := tokenizer.GenerateSymetricBearerToken(subject, durationInHours)
  if err != nil {
    // handle error
  }
  result, err := tokenizer.ReadSymetricBearerToken(token)
  if err != nil {
    // handle error
  }
  fmt.Println(result.Subject)
  // Output: 0x3
}  
```

#### Functions
  
  Generate and read  a private bearer token
  ```
  GeneratePrivateBearerToken(subject string, durationInHours Duration) (bearer string, err error)

  ReadPrivateBearerToken(token string) (jsonToken ReaderResult, err error)
  ```

  Generate and read a symetric bearer token
  ```
  GenerateSymetricBearerToken(subject string, durationInHours Duration) (bearer string, err error)

  ReadSymetricBearerToken(token string) (jsonToken ReaderResult, err error)
  ```

  Generate and read a private token
  ```
  GeneratePrivateToken(subject string, durationInHours Duration) (token string, err error)

  ReadPrivateToken(token string) (jsonToken ReaderResult, err error)
  ```

  Generate and read a symetric token
  ```
  GenerateSymetricToken(subject string, durationInHours Duration) (token string, err error)

  ReadSymetricToken(token string) (jsonToken ReaderResult, err error)
  ```

#### Duration in hours

```
type Duration int64
```

Already set durations:
- `OneDay`
- `TwoDays`
- `ThreeDays`
- `OneWeek`
- `TwoWeeks`
- `ThreeWeeks`
- `OneMonth`
- `TwoMonths`
- `ThreeMonths`

