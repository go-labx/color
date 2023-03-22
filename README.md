# color

The color library is used for outputting colored text in the console using ANSI escape sequences.

## Installation

```bash
go get github.com/go-labx/color
```

## Usage

To use this color lib, import it in your Go code:

```go
import "github.com/go-labx/color"
```

Here is an example of how to use color:

```go
package main

import (
	"fmt"
	"github.com/go-labx/color"
)

func main() {
	fmt.Println(color.BlackString("username = %s age = %d city = %s", "Jack", 20, "Hangzhou"))
	fmt.Println(color.RedString("username = %s age = %d city = %s", "Jack", 20, "Hangzhou"))
	fmt.Println(color.GreenString("username = %s age = %d city = %s", "Jack", 20, "Hangzhou"))
	fmt.Println(color.YellowString("username = %s age = %d city = %s", "Jack", 20, "Hangzhou"))
	fmt.Println(color.BlueString("username = %s age = %d city = %s", "Jack", 20, "Hangzhou"))
	fmt.Println(color.MagentaString("username = %s age = %d city = %s", "Jack", 20, "Hangzhou"))
	fmt.Println(color.CyanString("username = %s age = %d city = %s", "Jack", 20, "Hangzhou"))
	fmt.Println(color.WhiteString("username = %s age = %d city = %s", "Jack", 20, "Hangzhou"))
}
```

## API Documentation

For detailed API documentation and usage examples, please refer to the [documentation](https://pkg.go.dev/github.com/go-labx/color).

## Contributing

Contributions are welcome! Please see [CONTRIBUTING](./CONTRIBUTING.md) for more information.

## License

My SDK is licensed under the MIT License. See [LICENSE](./LICENSE) for more information.
