package dbrepo

import (
	"audit/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
)

type SQLServerDBRepo struct {
	DB *sql.DB
}

//const dbTimeout = time.Second * 3

func (m *SQLServerDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *SQLServerDBRepo) GetMachineBySiteName(siteName string) (*models.CloudMachine, error) {
	ctx := context.Background() //context.WithTimeout(context.Background(), dbTimeout)
	//defer cancel()

	where := "Where SiteName = @SiteName"

	query := fmt.Sprintf(`
		select
			ID, coalesce(ServerName, ''), coalesce(SiteName, ''), coalesce(Customercode, ''), coalesce(CustomerName, '')
		from
			CloudServers %s
		order by ID desc
	`, where)

	log.Println(query)
	row := m.DB.QueryRowContext(ctx, query, sql.Named("SiteName", siteName))

	var cloudMachine models.CloudMachine

	err := row.Scan(
		&cloudMachine.ID,
		&cloudMachine.ServerName,
		&cloudMachine.SiteName,
		&cloudMachine.CustomerDLCode,
		&cloudMachine.CustomerTitle,
	)
	if err != nil {
		return nil, err
	}

	return &cloudMachine, nil
}

func (m *SQLServerDBRepo) GetAllMachines() ([]*models.CloudMachine, error) {
	ctx := context.Background() //context.WithTimeout(context.Background(), dbTimeout)
	//defer cancel()

	where := ""

	query := fmt.Sprintf(`
		select
			ID, coalesce(ServerName, ''), coalesce(SiteName, ''), coalesce(Customercode, ''), coalesce(CustomerName, '')
		from
			CloudServers %s
		order by ID desc
	`, where)

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cloudMachines []*models.CloudMachine

	for rows.Next() {
		var cm models.CloudMachine
		//cm.New(&cm.ID,&cm.ServerName, &cm.SiteName,  &cm.CustomerTitle, &cm.CustomerDLCode)
		err := rows.Scan(
			&cm.ID,
			&cm.ServerName,
			&cm.SiteName,
			&cm.CustomerDLCode,
			&cm.CustomerTitle,
		)
		if err != nil {
			return nil, err
		}
		cloudMachines = append(cloudMachines, &cm)
	}

	return cloudMachines, nil
}

func (m *SQLServerDBRepo) ExecuteAuditScripts(cloudMachine *models.CloudMachine, contextualParameters []*models.ContextualParameter, auditScrips []*models.AuditScript) ([]*models.AuditScript, error) {
	ctx := context.Background()
	var finalAuditScripts []*models.AuditScript

	for _, v := range auditScrips {

		for _, cp := range contextualParameters {
			if strings.Contains(v.Query, cp.Key) {
				if cp.DataType == "int" || cp.DataType == "float" {
					v.Query = strings.ReplaceAll(v.Query, cp.Key, cp.Value)
				} else {
					v.Query = strings.ReplaceAll(v.Query, cp.Key, "'"+cp.Value+"'")
				}
			}
		}

		row := m.DB.QueryRowContext(ctx, v.Query)
		var resultValue int
		err := row.Scan(&resultValue)
		if err != nil {
			fmt.Printf("faild to exec script %v with error %v", v.Query, err.Error())
			continue
		}

		switch v.Operator {
		case models.E:
			v.Succeeded = resultValue == (v.AcceptancePercent/100)*v.Value
		case models.Lt:
			v.Succeeded = resultValue < (v.AcceptancePercent/100)*v.Value
		case models.Lte:
			v.Succeeded = resultValue <= (v.AcceptancePercent/100)*v.Value
		case models.Gt:
			v.Succeeded = resultValue > (v.AcceptancePercent/100)*v.Value
		case models.Gte:
			v.Succeeded = resultValue >= (v.AcceptancePercent/100)*v.Value
		}
		v.ResultValue = resultValue
	}

	return finalAuditScripts, errors.New("")
}
