# Aktis Plugin SDK

The official Go SDK for building Aktis Collector plugins.

## Installation

```bash
go get github.com/ternarybob/aktis-plugin-sdk
```

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

    // Collect your metrics
    metrics := []plugin.Metric{
        {
            Timestamp:  time.Now(),
            PluginName: "my-plugin",
            Type:       "custom_metric",
            Data: map[string]interface{}{
                "value": 42,
            },
            Metadata: map[string]string{
                "environment": *environment,
            },
        },
    }

    // Build output
    output := plugin.CollectionOutput{
        Success:   true,
        Timestamp: time.Now(),
        Metrics:   metrics,
        Plugin: plugin.PluginInfo{
            Name:        "my-plugin",
            Version:     "1.0.0",
            Environment: *environment,
        },
        Stats: plugin.CollectionStats{
            Duration:     time.Since(startTime).String(),
            MetricsCount: len(metrics),
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