# Aktis Plugin SDK

[![Build and Release](https://github.com/ternarybob/aktis-plugin-sdk/actions/workflows/ci-cd.yml/badge.svg)](https://github.com/ternarybob/aktis-plugin-sdk/actions/workflows/ci-cd.yml)
[![Go Version](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

The official Go SDK for building Aktis Collector plugins.

## Installation

```bash
go get github.com/ternarybob/aktis-plugin-sdk
```

## CI/CD

This SDK includes automated build and release workflows:

- **Continuous Integration**: Runs tests, go vet, and formatting checks on all PRs and pushes
- **Automated Releases**: Creates GitHub releases with source code archives when code is pushed to main
- **Go Modules**: Automatically validates and tidies dependencies

### Workflow Features

- Tests run on Go 1.24+
- Code formatting validation with `gofmt`
- Static analysis with `go vet`
- Dependency caching for faster builds
- Automatic versioning and tagging
- Source code archives in releases

## Quick Start

```go
package main

import (
    "encoding/json"
    "flag"
    "os"
    "time"

    plugin "github.com/ternarybob/aktis-plugin-sdk"
)

func main() {
    environment := flag.String("env", "development", "Environment")
    flag.Parse()

    startTime := time.Now()

    // Collect your payloads
    payloads := []plugin.Payload{
        {
            Timestamp: time.Now(),
            Type:      "custom_data",
            Data: map[string]interface{}{
                "value": 42,
            },
            Metadata: map[string]string{
                "source": "sensor-1",
            },
        },
    }

    // Build output
    output := plugin.CollectorOutput{
        Success:   true,
        Timestamp: time.Now(),
        Payloads:  payloads,
        Collector: plugin.CollectorInfo{
            Name:        "my-plugin",
            Type:        plugin.CollectorTypeData,
            Version:     "1.0.0",
            Environment: *environment,
        },
        Stats: plugin.CollectorStats{
            Duration:     time.Since(startTime).String(),
            PayloadCount: len(payloads),
        },
    }

    // Output JSON to stdout
    json.NewEncoder(os.Stdout).Encode(output)
}
```

## Documentation

See the [Plugin SDK Documentation](https://github.com/aktis/aktis-collector/blob/main/docs/plugin-sdk.md) for complete guide.

## License

MIT