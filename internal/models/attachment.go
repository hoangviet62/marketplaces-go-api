package internal

type Attachment struct {
	Base
	AttachmentType string `json:"attachment"`
	AttachmentID   uint   `json:"attachment_id"`
	Url            string `json:"url"`
}
