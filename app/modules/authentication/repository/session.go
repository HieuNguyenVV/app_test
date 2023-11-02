package repository

import (
	"app/common"
	"app/modules/authentication/model"
	"context"
)

func (t *SQLRepository) GetAccountByEmail(ctx context.Context, email string) (*model.Account, error) {
	account := model.Account{}
	if err := t.Model(&model.Account{}).
		Where("email=?", email).
		First(&account).Error; err != nil {
		return nil, common.ErrorDB(err)
	}
	return &account, nil
}

func (t *SQLRepository) CreateSession(ctx context.Context, data *model.Session) error {
	if err := t.Model(&model.Session{}).
		Create(data).Error; err != nil {
		return common.ErrorDB(err)
	}
	return nil
}
