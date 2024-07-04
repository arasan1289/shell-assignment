package http

import (
	"github.com/arasan1289/shell-test/internal/core/domain"
	"github.com/arasan1289/shell-test/internal/core/port"
	"github.com/gin-gonic/gin"
)

// VisitorHandler handles visitor-related HTTP requests.
type VisitorHandler struct {
	svc port.IVisitorService
}

// NewVisitorHandler creates a new VisitorHandler instance.
func NewVisitorHandler(svc port.IVisitorService) *VisitorHandler {
	return &VisitorHandler{
		svc: svc,
	}
}

// newVisitor represents a new visitor request.
type newVisitor struct {
	Name string `json:"name" binding:"required"`
	Url  string `json:"url" binding:"required,http_url"`
}

// NewVisitor handles the creation of a new visitor.
func (vh *VisitorHandler) NewVisitor(ctx *gin.Context) {
	var req newVisitor
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}
	visitor := domain.Visitor{
		Name: req.Name,
		URL:  req.Url,
	}
	rsp, err := vh.svc.AddVisitor(ctx, &visitor)
	if err != nil {
		handleError(ctx, err)
		return
	}
	handleSuccess(ctx, rsp)
}

// GetVisitorCount handles the retrieval of visitor counts.
func (vh *VisitorHandler) GetVisitorCount(ctx *gin.Context) {
	rsp, err := vh.svc.GetCount(ctx)
	if err != nil {
		handleError(ctx, err)
		return
	}
	handleSuccess(ctx, rsp)
}
