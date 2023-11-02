package repository

import (
	"app/common"
	"app/modules/authentication/model"
	"context"
)

func (t *SQLRepository) CreateAccount(ctx context.Context, data *model.Account) error {
	if err := t.Model(&model.Account{}).
		Create(data).Error; err != nil {
		return common.ErrorDB(err)
	}
	return nil
}
func (t *SQLRepository) CheckExistByEmail(ctx context.Context, email string) (bool, error) {
	exist := false
	if err := t.Model(&model.Account{}).
		Select("count(*)>0").
		Where("email=?", email).
		Find(&exist).Error; err != nil {
		return exist, common.ErrorDB(err)
	}
	return exist, nil
}
