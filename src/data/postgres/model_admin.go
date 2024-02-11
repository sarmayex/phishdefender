package postgres

import (
	"github.com/sarmayex/phishdefender/domain"
	"time"
)

type AdminModel struct {
	Id           int64
	NationalCode string
	CreatedAt    time.Time
	UpdateAt     time.Time
}

func (a *AdminModel) TableName() string {
	return "admins"
}

func (a *AdminModel) convertEntityToModel(entity *domain.UserEntity) {
	a.Id = entity.Id
	a.NationalCode = entity.NationalCode
	a.CreatedAt = entity.CreatedAt
	a.UpdateAt = entity.UpdatedAt
}

func (a *AdminModel) convertModelToEntity() *domain.UserEntity {
	return &domain.UserEntity{
		Id:           a.Id,
		NationalCode: a.NationalCode,
		CreatedAt:    a.CreatedAt,
		UpdatedAt:    a.UpdateAt,
	}
}
