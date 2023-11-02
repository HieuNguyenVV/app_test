package repository

import (
	"app/common"
	"app/modules/authentication/model"
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (t *SQLRepository) GetSession(ctx context.Context, id uuid.UUID) (*model.Session, error) {
	session := model.Session{}
	if err := t.Model(&model.Session{}).Where("id=?", id).First(&session).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrorDB(err)
	}
	return &session, nil
}
func (t *SQLRepository) RenewSession(ctx context.Context, rfToken string, newSession *model.Session) error {
	err := t.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Session{}).Create(&newSession).Error; err != nil {
			return err
		}

		if err := t.Where("refresh_token=?", rfToken).Delete(&model.Session{}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return common.ErrorDB(err)
	}
	return nil
}
