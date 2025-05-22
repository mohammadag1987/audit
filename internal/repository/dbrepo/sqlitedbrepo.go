package dbrepo

import (
	"audit/internal/models"
	"context"
	"database/sql"
	"fmt"
)

type SQLiteDBRepo struct {
	DB *sql.DB
}

//const dbTimeout = time.Second * 3

func (m *SQLiteDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *SQLiteDBRepo) GetContexParams() ([]*models.ContextualParameter, error) {
	ctx := context.Background() //context.WithTimeout(context.Background(), dbTimeout)
	//defer cancel()

	where := ""

	query := fmt.Sprintf(`
		select id, key, description, value, schema 
		from
			auditconfig %s
		order by id desc
	`, where)

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cloudMachines []*models.ContextualParameter

	for rows.Next() {
		var cm models.ContextualParameter
		//cm.New(&cm.ID,&cm.ServerName, &cm.SiteName,  &cm.CustomerTitle, &cm.CustomerDLCode)
		err := rows.Scan(
			&cm.ID,
			&cm.Key,
			&cm.Description,
			&cm.Value,
			&cm.Schema,
		)
		if err != nil {
			return nil, err
		}
		cloudMachines = append(cloudMachines, &cm)
	}

	return cloudMachines, nil
}
