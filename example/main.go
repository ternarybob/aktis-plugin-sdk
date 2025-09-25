package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	plugin "github.com/ternarybob/aktis-plugin-sdk"
)

const (
	pluginName    = "example-plugin"
	pluginVersion = "1.0.0"
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
		fmt.Printf("%s v%s\n", pluginName, pluginVersion)
		return
	}
	if *help {
		showHelp()
		return
	}

	// Parse environment from mode
	environment := parseMode(*mode)

	// Load configuration (simplified for example)
	cfg := loadConfig(*configFile)

	// Show banner unless quiet
	if !*quiet {
		showBanner(environment)
	}

	startTime := time.Now()

	// Collect data
	payloads, err := collectExampleData(cfg, environment)
	if err != nil {
		handleError(err, *quiet, environment, startTime)
		return
	}

	// Build successful output
	output := plugin.CollectorOutput{
		Success:   true,
		Timestamp: time.Now(),
		Payloads:  payloads,
		Collector: plugin.CollectorInfo{
			Name:        pluginName,
			Type:        plugin.CollectorTypeData,
			Version:     pluginVersion,
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

func loadConfig(configFile string) map[string]interface{} {
	// Simplified config loading for example
	// In a real plugin, you would load from TOML/JSON file
	return map[string]interface{}{
		"enabled":        true,
		"sample_rate":    1000,
		"include_system": true,
	}
}

func showBanner(environment string) {
	fmt.Printf("🔧 %s v%s\n", pluginName, pluginVersion)
	fmt.Printf("📍 Environment: %s\n", environment)
	fmt.Printf("⏰ Started: %s\n\n", time.Now().Format("2006-01-02 15:04:05"))
}

func collectExampleData(cfg map[string]interface{}, environment string) ([]plugin.Payload, error) {
	var payloads []plugin.Payload

	// Example: System status payload
	systemPayload := plugin.Payload{
		Timestamp: time.Now(),
		Type:      "system_status",
		Data: map[string]interface{}{
			"uptime_seconds": 86400,
			"status":         "healthy",
			"load_average":   0.75,
		},
		Metadata: map[string]string{
			"hostname": "localhost",
			"platform": "example",
		},
	}
	payloads = append(payloads, systemPayload)

	// Example: Application metrics payload
	appPayload := plugin.Payload{
		Timestamp: time.Now(),
		Type:      "application_metrics",
		Data: map[string]interface{}{
			"requests_per_second": 150.5,
			"response_time_ms":    45.2,
			"error_rate":          0.01,
		},
		Metadata: map[string]string{
			"service": "web-server",
			"version": "2.1.0",
		},
	}
	payloads = append(payloads, appPayload)

	// Example: Custom business metric
	businessPayload := plugin.Payload{
		Timestamp: time.Now(),
		Type:      "business_metrics",
		Data: map[string]interface{}{
			"active_users":          1250,
			"revenue_today":         15420.50,
			"conversion_rate":       0.035,
			"customer_satisfaction": 4.7,
		},
		Metadata: map[string]string{
			"region":     "us-east",
			"datacenter": "primary",
		},
	}
	payloads = append(payloads, businessPayload)

	return payloads, nil
}

func handleError(err error, quiet bool, environment string, startTime time.Time) {
	if quiet {
		// Output error in JSON format for aktis-collector
		output := plugin.CollectorOutput{
			Success:   false,
			Timestamp: time.Now(),
			Error:     err.Error(),
			Collector: plugin.CollectorInfo{
				Name:        pluginName,
				Type:        plugin.CollectorTypeData,
				Version:     pluginVersion,
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
		fmt.Printf("❌ Error: %v\n", err)
		os.Exit(1)
	}
}

func displayResults(output plugin.CollectorOutput) {
	fmt.Printf("✅ Collection completed successfully!\n\n")
	fmt.Printf("📊 Summary:\n")
	fmt.Printf("   Duration: %s\n", output.Stats.Duration)
	fmt.Printf("   Payloads: %d\n", output.Stats.PayloadCount)
	fmt.Printf("   Environment: %s\n\n", output.Collector.Environment)

	fmt.Printf("📦 Collected Data:\n")
	for i, payload := range output.Payloads {
		fmt.Printf("\n%d. %s [%s]\n", i+1, payload.Type, payload.Timestamp.Format("15:04:05"))

		// Display key metrics
		for key, value := range payload.Data {
			fmt.Printf("   • %s: %v\n", key, value)
		}

		// Display metadata if present
		if len(payload.Metadata) > 0 {
			fmt.Printf("   📋 Metadata:")
			for key, value := range payload.Metadata {
				fmt.Printf(" %s=%s", key, value)
			}
			fmt.Println()
		}
	}

	fmt.Printf("\n🎉 Plugin execution completed!\n")
}

func showHelp() {
	fmt.Printf(`%s v%s
Example Aktis plugin demonstrating SDK usage

USAGE:
  %s [OPTIONS]

MODES:
  dev/development:  Development environment (default)
  prod/production:  Production environment

OPTIONS:
  -mode string
        Environment mode (default "dev")
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

CONFIGURATION:
  This example plugin supports basic configuration through TOML files.
  All configuration is optional and has sensible defaults.

OUTPUT MODES:
  • CLI Mode (default): Human-readable output with banners and formatting
  • JSON Mode (-quiet): Structured JSON output for aktis-collector integration

`, pluginName, pluginVersion, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])
}
