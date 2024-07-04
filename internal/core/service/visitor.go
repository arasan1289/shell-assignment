package service

import (
	"context"
	"net/url"

	"github.com/arasan1289/shell-test/internal/core/domain"
	"github.com/arasan1289/shell-test/internal/core/port"
)

// VisitorService struct represents the visitor service with its dependencies
type VisitorService struct {
	repo port.IVisitorRepository // visitor repository interface
}

// NewVisitorService constructor function
func NewVisitorService(repo port.IVisitorRepository) port.IVisitorService {
	return &VisitorService{
		repo: repo,
	}
}

// AddVisitor adds a new visitor to the repository
func (vs *VisitorService) AddVisitor(ctx context.Context, visitor *domain.Visitor) (bool, error) {
	// Parse the URL to extract the domain
	u, err := url.Parse(visitor.URL)
	if err != nil {
		return false, err
	}
	visitor.Domain = u.Scheme + "://" + u.Host
	res, err := vs.repo.AddVisitor(ctx, visitor)
	if err != nil {
		return false, err
	}
	return res, nil
}

// GetCount retrieves the visitor count from the repository
func (vs *VisitorService) GetCount(ctx context.Context) (*[]domain.VisitorCount, error) {
	res, err := vs.repo.GetCount(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
