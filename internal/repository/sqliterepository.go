package repository

import (
	"audit/internal/models"
)

type SQLiteDatabaseRepo interface {
	GetContexParams() ([]*models.ContextualParameter, error)
}
