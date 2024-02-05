package domain

import "context"

type DefenderRepository interface {
	FirstOrCreate(ctx context.Context)
}
