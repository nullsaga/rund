package api

import (
	"net/http"
)

type WebhookHandler struct {
}

func NewWebhookHandler() *WebhookHandler {
	return &WebhookHandler{}
}

func (wh *WebhookHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	return nil
}
