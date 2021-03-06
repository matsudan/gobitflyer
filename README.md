# gobitflyer
[![Test Status](https://github.com/matsudan/gobitflyer/workflows/tests/badge.svg)](https://github.com/matsudan/gobitflyer/actions?query=workflow%3Atests)
[![codecov](https://codecov.io/gh/matsudan/gobitflyer/branch/main/graph/badge.svg)](https://codecov.io/gh/matsudan/gobitflyer)
[![Go Report Card](https://goreportcard.com/badge/matsudan/gobitflyer)](https://goreportcard.com/report/github.com/matsudan/gobitflyer)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/matsudan/gobitflyer/blob/master/LICENSE)

:construction: This library is under construction.

`gobitflyer` is bitFlyer Lightning API library for Go.

## Installation
```sh
go get github.com/matsudan/gobitflyer
```

## Usage
```go
import "github.com/matsudan/gobitflyer/bitflyer"
```

Construct a new bitFlyer client, then access the bitFlyer APIs. For example:
```go
client := bitflyer.NewClient(nil)

ticker, err := client.GetTicker(contexnt.Background(), "BTC_JPY")
```

### Authentication
Set the `BITFLYER_API_KEY` and `BITFLYER_API_SECRET` environment variables.
To set these variables on Linux, macOS, or Unix, use `export` :

```shell
export BITFLYER_API_KEY=your_api_key
export BITFLYER_API_SECRET=your_api_secret
```

For more detailed information, see [here](https://lightning.bitflyer.com/docs?lang=en#authentication)

## License
This library is distributed under the MIT License.

## Author
@matsudan
