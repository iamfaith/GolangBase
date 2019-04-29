package model

type UploadFile struct {
	Status string `json:"status" bson:"status"`
	FilePath   string `json:"-"`
}
