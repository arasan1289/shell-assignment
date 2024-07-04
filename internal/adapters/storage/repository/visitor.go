package repository

import (
	"context"
	"errors"

	"github.com/arasan1289/shell-test/internal/core/domain"
	"github.com/arasan1289/shell-test/internal/core/port"
)

// VisitorRepository implements the IVisitorRepository interface.
type VisitorRepository struct {
	repo *domain.Visitors
}

// NewVisitorRepository creates a new VisitorRepository instance.
func NewVisitorRepository(visitors *domain.Visitors) port.IVisitorRepository {
	return &VisitorRepository{
		repo: visitors,
	}
}

// AddVisitor adds a new visitor to the repository.
func (vr *VisitorRepository) AddVisitor(ctx context.Context, visitor *domain.Visitor) (bool, error) {
	if visitor != nil {
		vr.repo.MU.Lock()
		defer vr.repo.MU.Unlock()
		vr.repo.Visitors = append(vr.repo.Visitors, *visitor)
		return true, nil
	}
	return false, errors.New("Nil object passed.")
}

// GetCount retrieves the count of visitors from the repository.
func (vr *VisitorRepository) GetCount(ctx context.Context) (*[]domain.VisitorCount, error) {
	vr.repo.MU.Lock()
	defer vr.repo.MU.Unlock()

	domainVisitorMap := make(map[string]map[string]bool)
	for _, visitor := range vr.repo.Visitors {
		if domainVisitorMap[visitor.Domain] == nil {
			domainVisitorMap[visitor.Domain] = make(map[string]bool)
		}
		domainVisitorMap[visitor.Domain][visitor.Name] = true
	}

	visitorounts := make([]domain.VisitorCount, 0, len(domainVisitorMap))
	for host, visitors := range domainVisitorMap {
		visitorounts = append(visitorounts, domain.VisitorCount{Domain: host, Count: len(visitors)})
	}

	return &visitorounts, nil

}
