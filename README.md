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

## Plugin Structure

Aktis plugins follow a standardized command line interface and execution model:

### Command Line Interface

All plugins should implement these standard flags:

```go
-mode string        Environment mode: "dev"/"development" or "prod"/"production" (default "dev")
-config string      Configuration file path (optional)
-quiet              Suppress banner and informational output (enables JSON mode)
-version            Show version information
-help               Show help message
```

### Execution Modes

**Quiet Mode** (`-quiet`): JSON output for aktis-collector integration
**CLI Mode** (default): Human-readable output with banners and formatting

## Quick Start

```go
package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "time"

    plugin "github.com/ternarybob/aktis-plugin-sdk"
)

func main() {
    // Standard command line flags
    mode := flag.String("mode", "dev", "Environment mode: 'dev', 'development', 'prod', or 'production'")
    configFile := flag.String("config", "", "Configuration file path")
    quiet := flag.Bool("quiet", false, "Suppress banner output")
    version := flag.Bool("version", false, "Show version information")
    help := flag.Bool("help", false, "Show help message")
    flag.Parse()

    // Handle version and help
    if *version {
        fmt.Println("my-plugin v1.0.0")
        return
    }
    if *help {
        showHelp()
        return
    }

    // Parse environment from mode
    environment := parseMode(*mode)

    // Load configuration (if needed)
    cfg := loadConfig(*configFile)

    startTime := time.Now()

    // Your collection logic here
    payloads, err := collectData(cfg, environment)
    if err != nil {
        if *quiet {
            // Output error in JSON format for aktis-collector
            output := plugin.CollectorOutput{
                Success:   false,
                Timestamp: time.Now(),
                Error:     err.Error(),
                Collector: plugin.CollectorInfo{
                    Name:        "my-plugin",
                    Type:        plugin.CollectorTypeData,
                    Version:     "1.0.0",
                    Environment: environment,
                },
                Stats: plugin.CollectorStats{
                    Duration:     time.Since(startTime).String(),
                    PayloadCount: 0,
                    ErrorCount:   1,
                },
            }
            json.NewEncoder(os.Stdout).Encode(output)
        } else {
            fmt.Printf("Error: %v\n", err)
        }
        return
    }

    // Build successful output
    output := plugin.CollectorOutput{
        Success:   true,
        Timestamp: time.Now(),
        Payloads:  payloads,
        Collector: plugin.CollectorInfo{
            Name:        "my-plugin",
            Type:        plugin.CollectorTypeData, // or CollectorTypeEvent
            Version:     "1.0.0",
            Environment: environment,
        },
        Stats: plugin.CollectorStats{
            Duration:     time.Since(startTime).String(),
            PayloadCount: len(payloads),
        },
    }

    if *quiet {
        // JSON output for aktis-collector
        json.NewEncoder(os.Stdout).Encode(output)
    } else {
        // Human-readable CLI output
        displayResults(output)
    }
}

func parseMode(mode string) string {
    mode = strings.ToLower(mode)
    switch mode {
    case "prod", "production":
        return "production"
    case "dev", "development":
        return "development"
    default:
        return "development"
    }
}

func collectData(cfg interface{}, environment string) ([]plugin.Payload, error) {
    // Your data collection logic
    payloads := []plugin.Payload{
        {
            Timestamp: time.Now(),
            Type:      "custom_metric",
            Data: map[string]interface{}{
                "value": 42,
                "status": "active",
            },
            Metadata: map[string]string{
                "source": "sensor-1",
                "unit":   "percentage",
            },
        },
    }
    return payloads, nil
}

func loadConfig(configFile string) interface{} {
    // Load your TOML/JSON configuration
    // Return default config if no file specified
    return nil
}

func displayResults(output plugin.CollectorOutput) {
    fmt.Printf("=== %s Results ===\n", output.Collector.Name)
    fmt.Printf("Environment: %s\n", output.Collector.Environment)
    fmt.Printf("Duration: %s\n", output.Stats.Duration)
    fmt.Printf("Payloads: %d\n", output.Stats.PayloadCount)

    for _, payload := range output.Payloads {
        fmt.Printf("\n%s:\n", payload.Type)
        for key, value := range payload.Data {
            fmt.Printf("  %s: %v\n", key, value)
        }
    }
}

func showHelp() {
    fmt.Printf(`My Plugin v1.0.0
Custom data collection plugin

USAGE:
  %s [OPTIONS]

OPTIONS:
  -mode string
        Environment mode: dev/development or prod/production (default "dev")
  -config string
        Configuration file path
  -quiet
        Suppress banner output (JSON output mode)
  -version
        Show version information
  -help
        Show this help message

EXAMPLES:
  # Development mode with CLI output
  %s

  # Production mode
  %s -mode prod

  # JSON output for aktis-collector
  %s -quiet

  # With custom config
  %s -config /path/to/config.toml
`, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])
}
```

## Configuration Structure

Plugins typically use TOML configuration files with the following standard structure:

```toml
# Environment identifier
environment = "development"

# Plugin-specific settings
enable_feature_a = true
enable_feature_b = false
timeout = "30s"

[output]
# Output format: summary, table, json
format = "summary"
timestamp = true
verbose = false

[logging]
# Log level: debug, info, warn, error
level = "info"
output = "console"
file = "plugin.log"

[system]
# System-specific settings
custom_setting = "value"
```

## Example Implementation

See the complete [example plugin](./example/) for a full implementation demonstrating:

- Standard command line interface
- Dual execution modes (CLI/JSON)
- Configuration loading
- Error handling
- Multiple payload types
- Human-readable output formatting

To run the example:

```bash
cd example
go build -o example-plugin main.go
./example-plugin            # CLI mode
./example-plugin -quiet     # JSON mode
./example-plugin -help      # Show help
```

## Documentation

See the [Plugin SDK Documentation](https://github.com/aktis/aktis-collector/blob/main/docs/plugin-sdk.md) for complete guide.

## License

MIT