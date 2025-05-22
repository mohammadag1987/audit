package main

import (
	"audit/internal/models"
	"audit/internal/repository/dbrepo"
	"audit/internal/services"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/go-chi/chi/v5"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello World from %s!", app.Domain)

	type payload struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}
	pl := payload{Status: "active", Message: "Go Mohammad Ag Project up and running", Version: "1.0.0"}

	_ = app.writeJSON(w, http.StatusOK, pl)

}
func (app *application) TestSQLServer(w http.ResponseWriter, r *http.Request) {

	machines, err := app.SQLServerDBConnection.GetAllMachines()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	_ = app.writeJSON(w, http.StatusOK, machines)
}

func (app *application) GetContexParams(w http.ResponseWriter, r *http.Request) {

	machines, err := app.SQLiteDBConnection.GetContexParams()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	_ = app.writeJSON(w, http.StatusOK, machines)
}

func (app *application) CheckSiteURL(w http.ResponseWriter, r *http.Request) {
	sitename := chi.URLParam(r, "sitename")
	if sitename == "" {
		app.errorJSON(w, errors.New("مقدار آدرس سایت مشترک اجباری است"))
		return
	}

	cm, err := app.SQLServerDBConnection.GetMachineBySiteName(sitename)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, cm)
}

func (app *application) GetAuditScripts(w http.ResponseWriter, r *http.Request) {

	cm, err := services.GetAllAuditScripts()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, cm)
}

func (app *application) ExecuteAudit(w http.ResponseWriter, r *http.Request) {
	sitename := chi.URLParam(r, "sitename")
	if sitename == "" {
		app.errorJSON(w, errors.New("مقدار آدرس سایت مشترک اجباری است"))
		return
	}
	cm, err := app.SQLServerDBConnection.GetMachineBySiteName(sitename)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var arp *models.AuditReceivePayload
	err = app.readJSON(w, r, arp)
	if err != nil {
		app.errorJSON(w, err)
	}

	var errors []string

	for _, cp := range arp.ContextualParameter {
		if slices.Contains(arp.Modules, cp.Schema) {
			//یعنی ماژول انتخاب شده و می بایست مقدار پارامتر ضمنی چک شود
			if cp.Value == "" || cp.Value == "-" {
				errors = append(errors, "مقدار پارامتر ضمنی"+cp.Key+" "+cp.Description+" وارد نشده است ")
			}
		}
	}

	auditScripts, err := services.GetAllAuditScripts()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	nCon, err := app.ConnectToMachineSQLServer()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	app.SQLServeMachineDSN = fmt.Sprintf("sqlserver://sa:1@%v:14330?database=SGDB&encrypt=disable", cm.ServerName)
	app.SQLServerMachineDBConnection = &dbrepo.SQLServerDBRepo{DB: nCon}
	defer app.SQLServerDBConnection.Connection().Close()
	_, err = app.SQLServerMachineDBConnection.ExecuteAuditScripts(cm, arp.ContextualParameter, auditScripts)

	if err != nil {
		resp := JSONResponse{
			Error:   false,
			Message: "movie Inserted",
		}
		_ = app.writeJSON(w, http.StatusBadRequest, resp)
	} else if len(errors) > 0 {
		resp := JSONResponse{
			Error:   false,
			Message: "movie Inserted",
			Data:    errors,
		}
		_ = app.writeJSON(w, http.StatusBadRequest, resp)
	}

	_ = app.writeJSON(w, http.StatusOK, auditScripts)
}
