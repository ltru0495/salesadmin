package models

import (
	"encoding/json"
	"log"
)

type ApiResponse struct {
	Status  int         `json:"status"`
	Content interface{} `json:"content"`
}

func (r *ApiResponse) Prepare() []byte {
	j, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
	}
	return j
}
