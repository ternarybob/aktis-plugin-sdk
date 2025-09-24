package plugin

import "time"

// CollectionOutput represents the JSON output from a plugin execution
type CollectionOutput struct {
	Success   bool            `json:"success"`         // Overall success status
	Timestamp time.Time       `json:"timestamp"`       // Collection timestamp
	Metrics   []Metric        `json:"metrics"`         // Collected metrics array
	Error     string          `json:"error,omitempty"` // Error message if failed
	Plugin    PluginInfo      `json:"plugin"`          // Plugin metadata
	Stats     CollectionStats `json:"stats,omitempty"` // Collection statistics
}

// PluginInfo contains plugin metadata
type PluginInfo struct {
	Name        string `json:"name"`                  // Plugin name
	Version     string `json:"version,omitempty"`     // Plugin version
	Environment string `json:"environment,omitempty"` // Runtime environment
}

// CollectionStats contains collection statistics
type CollectionStats struct {
	Duration     string `json:"duration,omitempty"`    // Collection duration
	MetricsCount int    `json:"metrics_count"`          // Number of metrics
	ErrorCount   int    `json:"error_count,omitempty"` // Number of errors
}

// Metric represents a single metric collected by a plugin
type Metric struct {
	Timestamp  time.Time              `json:"timestamp"`   // Metric timestamp
	PluginName string                 `json:"plugin_name"` // Plugin identifier
	Type       string                 `json:"type"`        // Metric type
	Data       map[string]interface{} `json:"data"`        // Metric payload
	Metadata   map[string]string      `json:"metadata"`    // Additional context
}

// PluginStatus represents the current status of a plugin
type PluginStatus struct {
	Name           string    `json:"name"`            // Plugin name
	Running        bool      `json:"running"`         // Is plugin running
	LastCollection time.Time `json:"last_collection"` // Last collection time
	MetricsCount   int       `json:"metrics_count"`   // Total metrics collected
	ErrorCount     int       `json:"error_count"`     // Total errors encountered
	Message        string    `json:"message"`         // Status message
}