package repository

import (
	"audit/internal/models"
	"database/sql"
)

type SQLServerDatabaseRepo interface {
	Connection() *sql.DB
	GetAllMachines() ([]*models.CloudMachine, error)
	GetMachineBySiteName(siteName string) (*models.CloudMachine, error)
	ExecuteAuditScripts(cloudMachine *models.CloudMachine, contextualParameters []*models.ContextualParameter, auditScrips []*models.AuditScript) ([]*models.AuditScript, error)
}
