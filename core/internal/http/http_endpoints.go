package http

import (
	"core/pkg/access"
	"core/pkg/ping"
	"core/pkg/project"
	"core/pkg/workspace"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies route from all packages to root handler
func ApplyRoutes(r *gin.Engine) {
	ApplyMetrics(r)
	root := r.Group("/")
	access.ApplyRoutes(root)
	project.ApplyRoutes(root)
	ping.ApplyRoutes(root)
	workspace.ApplyRoutes(root)
}
