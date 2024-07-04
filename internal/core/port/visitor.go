package port

import (
	"context"

	"github.com/arasan1289/shell-test/internal/core/domain"
)

// IVisitorRepository interface defines the methods for interacting with the visitor repository
type IVisitorRepository interface {
	AddVisitor(ctx context.Context, visitor *domain.Visitor) (bool, error)
	GetCount(ctx context.Context) (*[]domain.VisitorCount, error)
}

// IVisitorService interface defines the methods for interacting with the visitor service
type IVisitorService interface {
	AddVisitor(ctx context.Context, visitor *domain.Visitor) (bool, error)
	GetCount(ctx context.Context) (*[]domain.VisitorCount, error)
}
