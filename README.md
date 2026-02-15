# Gonika

Gonika is a Go library for decoding AVL (Automatic Vehicle Location) data packets and GPRS messages (Teltonika-style). It supports multiple device codecs and provides typed models and parsers for TCP and UDP variants.

## Features

- Decode Device Data Sending packets: Codec 8, 8e, 16
- Decode GPRS messages: Codec 12 (13/14 placeholders)
- Separate parsing for TCP and UDP channels
- Typed output models and flat representations for easy consumption
- Modular parser utilities to handle big-endian binary fields

## Installation

Require Go 1.20+ (or the Go version declared in `go.mod`).

Install with:

```bash
go get github.com/rojack96/gonika
```

Or add the module to your project's `go.mod` and run `go build`.

## Quick Start

Example: decode a device packet using the factory:

```go
package main

import (
    "fmt"
    cdc "github.com/rojack96/gonika/codec"
)

func main() {
    raw := /* your packet bytes or hex string input */
    decoder, err := cdc.DeviceDataSendingDecoderFactory(raw)
    if err != nil {
        panic(err)
    }

    // For TCP packets
    avl := decoder.DecodeTCP()

    // For UDP packets
    avlUdp := decoder.DecodeUDP()

    // Flat representation
    flat := decoder.DecodeFlat()

    fmt.Println(avl, avlUdp, flat)
}
```

GPRS messages use a separate factory:

```go
gprsDecoder, err := cdc.GprsMessageDecoderFactory(rawGprsPacket)
resp := gprsDecoder.Decode()
```

## Project Layout

- `codec/` — main factories, decoders and parsers
  - `device_data_sending/` — codec_8, codec_8e, codec_16 implementations
  - `gprs_message/` — codec_12 and related builders
  - `parsers/` — base parser utilities
  - `models/` — typed structs for decoded data
- `test/` — small example runner

## Testing & Development

Run tests and build locally:

```bash
go test ./...
go build ./...
go run test/test.go
```

## Contributing

Issues and pull requests are welcome. Please keep changes focused and include tests for new behavior.

## License

See the `LICENSE` file in the repository root.
# Gonika

Gonika is a Go library for decoding data from Teltonika devices using 
Teltonika codecs. 
This library supports various codecs, including Codec 8, Codec 8 Extended, and Codec 16,
enabling developers to easily interpret GPS and IoT data from Teltonika devices.
