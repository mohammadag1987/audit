package models

type AuditScript struct {
	Code              string `json:"code"`
	Schema            string `json:"schema"`
	Category          string `json:"category"`
	Question          string `json:"question"`
	Query             string `json:"query"`
	Value             int    `json:"value"`
	AcceptancePercent int    `json:"acceptance_percent"`
	Coefficient       int    `json:"coefficient"`
	Type              int    `json:"type"`
	Veto              bool   `json:"veto"`
	Operator          Op     `json:"operator"`
	Succeeded         bool   `json:"result"`
	ResultValue       int    `json:"result_value"`
}

func (cm *AuditScript) New(code string, schema string, category string, question string, query string, value int, acceptancePercent int, coefficient int, typeOfScript int, veto bool, operator Op, succeeded bool, resultValue int) *AuditScript {
	return &AuditScript{
		Code:              code,
		Schema:            schema,
		Category:          category,
		Question:          question,
		Query:             query,
		Value:             value,
		AcceptancePercent: acceptancePercent,
		Coefficient:       coefficient,
		Type:              typeOfScript,
		Veto:              veto,
		Operator:          operator,
		Succeeded:         succeeded,
		ResultValue:       resultValue,
	}
}
