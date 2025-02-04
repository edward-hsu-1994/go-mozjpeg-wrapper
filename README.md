# go-mozjpeg-wrapper

`go-mozjpeg-wrapper` is a Go wrapper for the MozJPEG command-line JPEG encoder. It allows you to easily encode JPEG images using the powerful and efficient MozJPEG library through Go interfaces.

## Features

- Encode JPEG images synchronously and asynchronously
- Builder pattern to configure JPEG encoding options
- Customize quality, DC scan optimization, DCT method, and many more settings
- Supports well-known MozJPEG options like progressive, baseline, smooth, etc.

## Installation

To use `go-mozjpeg-wrapper`, ensure you have [MozJPEG](https://github.com/mozilla/mozjpeg) installed on your system. Then, include the package in your Go project by importing it:

```go
import "github.com/edward-hsu-1994/go-mozjpeg-wrapper"
```

## Usage

Below is a basic example of how to use this package to encode a JPEG image:

```go
package main

import (
    "fmt"
    "go-mozjpeg-wrapper"  // Import the package
    "log"
)

func main() {
    builder := go_mozjpeg_wrapper.NewMozJpegEncoderBuilder()
	builder = builder.SetCliExecutablePath("/path/to/mozjpeg/cjpeg")
    encoder, err := builder.SetQuality(85).Build()
    if err != nil {
        log.Fatalf("Failed to build encoder: %v", err)
    }

    err = encoder.Encode("input.jpg", "output.jpg")
    if err != nil {
        log.Fatalf("Failed to encode image: %v", err)
    }

    fmt.Println("Image successfully encoded.")
}
```

### Using the Builder

The `MozJpegEncoderBuilder` provides a fluent API to configure the encoder. Here are some of the available methods:

- `SetQuality(quality int)`: Set the image quality (0-100).
- `SetDCScanOpt(mode int)`: Set DC scan optimization (0-2).
- `SetDCTMethod(method DCTMethod)`: Set the DCT method (int, fast, float).
- `SetProgressive(enabled bool)`: Enable/disable progressive encoding.

Each of these methods returns a new builder instance, allowing for chained calls.

### Asynchronous Encoding

The package also supports asynchronous image encoding:

```go
errChan := encoder.EncodeAsync("input.jpg", "output.jpg")
for err := range errChan {
    if err != nil {
        log.Printf("Error encoding image: %v", err)
    } else {
        fmt.Println("Image successfully encoded asynchronously.")
    }
}
```

## Contributing

Contributions are welcome. Feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Acknowledgments

- [MozJPEG](https://github.com/mozilla/mozjpeg) for the powerful JPEG encoder.
