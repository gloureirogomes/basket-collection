package repository

import (
	"context"

	"github.com/GabrielLoureiroGomes/basket-collection/internal/core/domain"
)

// PlayerFirestoreRepository implements database repository interface
type PlayerFirestoreRepository struct {
	firestoreClient FirestoreClient
}

// NewPlayerFirestoreRepository returns a new instance of PlayerFirestoreRepository
func NewPlayerFirestoreRepository() PlayerFirestoreRepository {
	return PlayerFirestoreRepository{
		firestoreClient: NewFirestoreClient(),
	}
}

// InsertPlayer used to save on database player data
func (p PlayerFirestoreRepository) InsertPlayer(ctx context.Context, player *domain.Player) error {
	return nil
}
