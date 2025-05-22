package services

import (
	"audit/internal"
	"audit/internal/models"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	zipEncrypted "github.com/alexmullins/zip"
)

const downloadAuditScriptPath string = "https://techlogger.systemgroup.net/SGTroubleshooterRepository/Home/SDFDSFDSFSDFSDFdfdsfDSFWSDFSDf"
const zipPass = "sghektor"

func GetAllAuditScripts() ([]*models.AuditScript, error) {

	resp, err := http.Get(downloadAuditScriptPath)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	tmpNameForZip := fmt.Sprintf("temp%s.zip", internal.GenerateRandom4DigitString())
	err = ioutil.WriteFile(tmpNameForZip, body, 0644)
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpNameForZip)

	reader, err := zipEncrypted.OpenReader(tmpNameForZip)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	var auditScripts []*models.AuditScript

	for _, f := range reader.File {

		if f.IsEncrypted() {
			f.SetPassword(zipPass)
		}

		rc, err := f.Open()
		if err != nil {
			fmt.Println("Error opening file inside zip:", err)
			continue
		}
		defer rc.Close()

		//fmt.Println("Reading file:", f.Name)
		content, err := io.ReadAll(rc)
		if err != nil {
			fmt.Println("Error reading content:", err)
			continue
		}

		var as models.AuditScript
		err = xml.Unmarshal(content, &as)
		if err != nil {
			fmt.Println("Error reading content:", err)
			continue
		}

		auditScripts = append(auditScripts, &as)
	}

	return auditScripts, nil
}
