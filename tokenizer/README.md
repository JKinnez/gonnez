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
- PublicKeyEnvName: Defaults to `GONNEZ_PUBLIC_KEY`
- PrivateKeyEnvName: Defaults to `GONNEZ_PRIVATE_KEY`
- SymetricKeyEnvName: Defaults to `GONNEZ_SYMMETRIC_KEY`

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

func main() {
  tokenizer = tokenizer.New(tokenizer.Options{})
  subject := "0x3"
  token, err := tokenizer.GenerateSymetricBearerToken(subject, tokenizer.Durations.OneDay)
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

