# Example Aktis Plugin

This is a complete example demonstrating how to build an Aktis plugin using the SDK.

## Features

- Standard command line interface with all required flags
- Dual execution modes (CLI and JSON)
- Proper error handling for both modes
- Example data collection with multiple payload types
- Human-readable output formatting
- Help and version information

## Building

```bash
cd example
go build -o example-plugin main.go
```

## Usage

### Development Mode (CLI Output)

```bash
./example-plugin
```

### Production Mode

```bash
./example-plugin -mode prod
```

### JSON Output (for aktis-collector)

```bash
./example-plugin -quiet
```

### Help and Version

```bash
./example-plugin -help
./example-plugin -version
```

## Output Examples

### CLI Mode Output

```
🔧 example-plugin v1.0.0
📍 Environment: development
⏰ Started: 2025-01-15 10:30:45

✅ Collection completed successfully!

📊 Summary:
   Duration: 1.234ms
   Payloads: 3
   Environment: development

📦 Collected Data:

1. system_status [10:30:45]
   • uptime_seconds: 86400
   • status: healthy
   • load_average: 0.75
   📋 Metadata: hostname=localhost platform=example

2. application_metrics [10:30:45]
   • requests_per_second: 150.5
   • response_time_ms: 45.2
   • error_rate: 0.01
   📋 Metadata: service=web-server version=2.1.0

3. business_metrics [10:30:45]
   • active_users: 1250
   • revenue_today: 15420.5
   • conversion_rate: 0.035
   • customer_satisfaction: 4.7
   📋 Metadata: region=us-east datacenter=primary

🎉 Plugin execution completed!
```

### JSON Mode Output

```json
{
  "success": true,
  "timestamp": "2025-01-15T10:30:45.123Z",
  "payloads": [
    {
      "timestamp": "2025-01-15T10:30:45.123Z",
      "type": "system_status",
      "data": {
        "uptime_seconds": 86400,
        "status": "healthy",
        "load_average": 0.75
      },
      "metadata": {
        "hostname": "localhost",
        "platform": "example"
      }
    }
  ],
  "collector": {
    "name": "example-plugin",
    "type": "data",
    "version": "1.0.0",
    "environment": "development"
  },
  "stats": {
    "duration": "1.234ms",
    "payload_count": 3
  }
}
```

## Code Structure

The example demonstrates:

1. **Standard CLI Interface**: All required flags (`-mode`, `-config`, `-quiet`, `-version`, `-help`)
2. **Mode Parsing**: Proper handling of dev/development and prod/production modes
3. **Dual Output**: JSON for aktis-collector integration, CLI for human use
4. **Error Handling**: Appropriate error responses for both modes
5. **Data Collection**: Multiple payload types with structured data and metadata
6. **Output Formatting**: Clean, professional presentation of results

## Integration with Aktis Collector

When using `-quiet` flag, the plugin outputs JSON in the exact format expected by aktis-collector, including:

- Success/failure status
- Timestamps
- Payload arrays with structured data
- Collector metadata
- Statistics (duration, counts, errors)

This makes the plugin ready for production use with the Aktis monitoring system.