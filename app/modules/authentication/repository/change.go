package repository

import (
	"app/common"
	"app/modules/authentication/model"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

func (t *SQLRepository) ChangePassword(ctx context.Context, hashedPass, email string, accountId uint32) error {
	err := t.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(model.Account{}.TableName()).
			Where("email=?", email).Updates(&model.Account{
			SQLModel: common.SQLModel{
				UpdateBy: accountId,
				UpdateAt: time.Now()},
			Password: hashedPass,
		}).Error; err != nil {
			return common.ErrorDB(err)
		}
		return nil
	})
	if err != nil {
		return common.ErrorDB(err)
	}
	return nil
}
func (t *SQLRepository) GetAccountByID(ctx context.Context, id uint32) (*model.Account, error) {
	account := model.Account{}
	if err := t.Model(&model.Account{}).Where("id=?", id).First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrorDB(err)
	}
	return &account, nil
}
