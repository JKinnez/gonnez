# Environment

Package to manage environment variables.

### Usage

```
package main

import "github.com/gonnez/environment"

func main() {
  environment.Init("YOUR_VALUE")

  if environment.IsProduction() {
    // Do something
  }

  if environment.IsDevelopment() {
    // Do something
  }

  if environment.IsTest() {
    // Do something
  }

  if environment.IsCI() {
    // Do something
  }


}
```
