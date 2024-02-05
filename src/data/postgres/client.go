package postgres

import (
	"context"
	"github.com/hamidteimouri/gommon/htenvier"
	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB
}

func NewPostgres(db *gorm.DB) *Postgres {
	ps := &Postgres{db: db}
	if htenvier.IsDebugMode() {
		ps.db = db.Debug()
	}
	return ps
}

func (p Postgres) FirstOrCreate(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}
