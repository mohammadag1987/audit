package models

type ContextualParameter struct {
	ID          int    `json:"id"`
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"datatype"`
	Schema      string `json:"schema"`
	DataType    string `json:"data_type"`
}

func (cp *ContextualParameter) New(id int, key string, value string, description string, schema string, dataType string) *ContextualParameter {
	return &ContextualParameter{
		ID:          id,
		Key:         key,
		Value:       value,
		Description: description,
		Schema:      schema,
		DataType:    dataType,
	}

}
