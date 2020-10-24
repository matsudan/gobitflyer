# gobitflyer
[![Test Status](https://github.com/matsudan/gobitflyer/workflows/tests/badge.svg)](https://github.com/matsudan/gobitflyer/actions?query=workflow%3Atests)
[![codecov](https://codecov.io/gh/matsudan/gobitflyer/branch/main/graph/badge.svg)](https://codecov.io/gh/matsudan/gobitflyer)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/matsudan/gobitflyer/blob/master/LICENSE)

`gobitflyer` is bitFlyer Lightning API library for Go.

## Requirements

## Usage
```go
import "github.com/matsudan/gobitflyer/bitflyer"
```

Construct a new bitFlyer client, then access the bitFlyer APIs. For example:
```go
cfg := bitflyer.LoadConfig()
client := bitflyer.NewClient(cfg)

ticker, err := client.GetTicker(contexnt.Background(), "BTC_JPY")
```

### Authentication
Set the `BITFLYER_API_KEY` and `BITFLYER_API_SECRET` environment variables.
To set these variables on Linux, macOS, or Unix, use `export` :
```
export BITFLYER_API_KEY=your_api_key
export BITFLYER_API_SECRET=your_api_secret
```

For more detailed information, see [here](https://lightning.bitflyer.com/docs?lang=en#authentication)

## License
This library is distributed under the MIT License.

## Author
@matsudan
