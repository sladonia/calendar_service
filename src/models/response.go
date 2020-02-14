package models

type ResponseCreated struct {
	Message   string `json:"message"`
	CreatedId string `json:"created_id"`
}

type ResponseDeleted struct {
	Message   string `json:"message"`
	DeletedId string `json:"deleted_id"`
}
