package postgres

import (
	"github.com/sarmayex/phishdefender/domain"
	"time"
)

type UserModel struct {
	Id           int64
	NationalCode string
	CreatedAt    time.Time
	UpdateAt     time.Time
}

func (u *UserModel) TableName() string {
	return "users"
}

func (u *UserModel) convertEntityToModel(entity *domain.UserEntity) {
	u.Id = entity.Id
	u.NationalCode = entity.NationalCode
	u.CreatedAt = entity.CreatedAt
	u.UpdateAt = entity.UpdatedAt
}

func (u *UserModel) convertModelToEntity() *domain.UserEntity {
	return &domain.UserEntity{
		Id:           u.Id,
		NationalCode: u.NationalCode,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdateAt,
	}
}
