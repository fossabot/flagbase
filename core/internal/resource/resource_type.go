package resource

// ResourceType specifies the type of resource object
type ResourceType string

func (r ResourceType) String() string {
	return string(r)
}

const (
	// Workspace represents a workspace resource
	Workspace = "workspace"
	// Project represents a project resource
	Project = "project"
	// Environment represents a environment resource
	Environment = "environment"
	// Access represents a environment resource
	Access = "access"
	// Flag represents a flag resource
	Flag = "flag"
	// Segment represents a segment resource
	Segment = "segment"
	// Identity represents a identity resource
	Identity = "identity"
	// Targeting represents a targeting resource
	Targeting = "targeting"
)
