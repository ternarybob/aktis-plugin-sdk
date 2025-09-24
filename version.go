package plugin

// Build-time variables set via ldflags
var (
	version   = "dev"
	build     = "unknown"
	gitCommit = "unknown"
)

// GetVersion returns the plugin SDK version
func GetVersion() string {
	return version
}

// GetBuild returns the build timestamp
func GetBuild() string {
	return build
}

// GetGitCommit returns the git commit hash
func GetGitCommit() string {
	return gitCommit
}

// SetVersion allows plugins to set their own version info
type VersionInfo struct {
	Version   string
	Build     string
	GitCommit string
}

// NewVersionInfo creates a new VersionInfo with defaults
func NewVersionInfo() *VersionInfo {
	return &VersionInfo{
		Version:   "dev",
		Build:     "unknown",
		GitCommit: "unknown",
	}
}

// WithVersion sets the version
func (v *VersionInfo) WithVersion(version string) *VersionInfo {
	v.Version = version
	return v
}

// WithBuild sets the build
func (v *VersionInfo) WithBuild(build string) *VersionInfo {
	v.Build = build
	return v
}

// WithGitCommit sets the git commit
func (v *VersionInfo) WithGitCommit(commit string) *VersionInfo {
	v.GitCommit = commit
	return v
}