package models

type AuditReceivePayload struct {
	Modules             []string               `json:"modules"`
	ContextualParameter []*ContextualParameter `json:"contextual_parameter"`
}
