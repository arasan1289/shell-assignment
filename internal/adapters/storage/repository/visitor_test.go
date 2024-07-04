package repository

import (
	"context"
	"testing"

	"github.com/arasan1289/shell-test/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewVisitorRepository(t *testing.T) {
	visitors := &domain.Visitors{}
	repo := NewVisitorRepository(visitors)
	assert.NotNil(t, repo)
}

func TestAddVisitor(t *testing.T) {
	visitors := &domain.Visitors{}
	repo := NewVisitorRepository(visitors)

	visitor := &domain.Visitor{Name: "John Doe", URL: "http://asdf.com?q=a", Domain: "http://asdf.com"}
	success, err := repo.AddVisitor(context.Background(), visitor)
	assert.True(t, success)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(visitors.Visitors))

	success, err = repo.AddVisitor(context.Background(), nil)
	assert.False(t, success)
	assert.EqualError(t, err, "Nil object passed.")
}

func TestAddMultipleVisitors(t *testing.T) {
	visitors := &domain.Visitors{}
	repo := NewVisitorRepository(visitors)

	visitor1 := &domain.Visitor{Name: "John Doe", URL: "http://asdf.com?q=a", Domain: "http://asdf.com"}
	visitor2 := &domain.Visitor{Name: "Jane Doe", URL: "http://asdf.com", Domain: "http://asdf.com"}
	visitor3 := &domain.Visitor{Name: "Alice", URL: "http://example.com", Domain: "http://example.com"}

	repo.AddVisitor(context.Background(), visitor1)
	repo.AddVisitor(context.Background(), visitor2)
	repo.AddVisitor(context.Background(), visitor3)

	assert.Equal(t, 3, len(visitors.Visitors))
}

func TestAddVisitorsSameNameDifferentDomain(t *testing.T) {
	visitors := &domain.Visitors{}
	repo := NewVisitorRepository(visitors)

	visitor1 := &domain.Visitor{Name: "John Doe", URL: "http://asdf.com?q=a", Domain: "http://asdf.com"}
	visitor2 := &domain.Visitor{Name: "John Doe", URL: "http://example.com", Domain: "http://example.com"}

	repo.AddVisitor(context.Background(), visitor1)
	repo.AddVisitor(context.Background(), visitor2)

	assert.Equal(t, 2, len(visitors.Visitors))
}

func TestAddVisitorsSameDomainDifferentName(t *testing.T) {
	visitors := &domain.Visitors{}
	repo := NewVisitorRepository(visitors)

	visitor1 := &domain.Visitor{Name: "John Doe", URL: "http://asdf.com?q=a", Domain: "http://asdf.com"}
	visitor2 := &domain.Visitor{Name: "Jane Doe", URL: "http://asdf.com", Domain: "http://asdf.com"}

	repo.AddVisitor(context.Background(), visitor1)
	repo.AddVisitor(context.Background(), visitor2)

	assert.Equal(t, 2, len(visitors.Visitors))
}

func TestGetCount(t *testing.T) {
	visitors := &domain.Visitors{}
	repo := NewVisitorRepository(visitors)

	visitor1 := &domain.Visitor{Name: "John Doe", URL: "http://asdf.com?q=a", Domain: "http://asdf.com"}
	visitor2 := &domain.Visitor{Name: "Jane Doe", URL: "http://asdf.com", Domain: "http://asdf.com"}
	visitor3 := &domain.Visitor{Name: "Alice", URL: "http://example.com", Domain: "http://example.com"}
	visitor4 := &domain.Visitor{Name: "John Doe", URL: "http://asdf.com?q=ab", Domain: "http://asdf.com"}

	repo.AddVisitor(context.Background(), visitor1)
	repo.AddVisitor(context.Background(), visitor2)
	repo.AddVisitor(context.Background(), visitor3)
	repo.AddVisitor(context.Background(), visitor4)

	counts, err := repo.GetCount(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, counts)

	expectedCounts := []domain.VisitorCount{
		{Domain: "http://asdf.com", Count: 2},
		{Domain: "http://example.com", Count: 1},
	}

	assert.ElementsMatch(t, expectedCounts, *counts)
}

func TestGetCount_EmptyRepository(t *testing.T) {
	visitors := &domain.Visitors{}
	repo := NewVisitorRepository(visitors)

	counts, err := repo.GetCount(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, counts)
	assert.Equal(t, 0, len(*counts))
}
