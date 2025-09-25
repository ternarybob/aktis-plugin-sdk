package plugin

import "time"

// CollectorType represents the type of collector
type CollectorType string

const (
	CollectorTypeEvent CollectorType = "event" // Event-based collector
	CollectorTypeData  CollectorType = "data"  // Data-based collector
)

// CollectorOutput represents the JSON output from a plugin execution
type CollectorOutput struct {
	Success   bool           `json:"success"`         // Overall success status
	Timestamp time.Time      `json:"timestamp"`       // Collector timestamp
	Payloads  []Payload      `json:"payloads"`        // Collected payloads array
	Error     string         `json:"error,omitempty"` // Error message if failed
	Collector CollectorInfo  `json:"collector"`       // Collector metadata
	Stats     CollectorStats `json:"stats,omitempty"` // Collector statistics
}

// CollectorInfo contains collector metadata
type CollectorInfo struct {
	Name        string        `json:"name"`                  // Collector name
	Type        CollectorType `json:"type"`                  // Collector type (event|data)
	Version     string        `json:"version,omitempty"`     // Collector version
	Environment string        `json:"environment,omitempty"` // Runtime environment
}

// CollectorStats contains collector statistics
type CollectorStats struct {
	Duration     string `json:"duration,omitempty"`    // Collector duration
	PayloadCount int    `json:"payload_count"`         // Number of payloads
	ErrorCount   int    `json:"error_count,omitempty"` // Number of errors
}

// Payload represents a single data payload collected by a plugin
type Payload struct {
	Timestamp time.Time              `json:"timestamp"` // Payload timestamp
	Type      string                 `json:"type"`      // Payload type
	Data      map[string]interface{} `json:"data"`      // Payload data
	Metadata  map[string]string      `json:"metadata"`  // Additional context
}

// PluginStatus represents the current status of a plugin
type PluginStatus struct {
	Name         string    `json:"name"`          // Plugin name
	Running      bool      `json:"running"`       // Is plugin running
	LastRun      time.Time `json:"last_run"`      // Last execution time
	PayloadCount int       `json:"payload_count"` // Total payloads collected
	ErrorCount   int       `json:"error_count"`   // Total errors encountered
	Message      string    `json:"message"`       // Status message
}
