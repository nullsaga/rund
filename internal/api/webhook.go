package api

import (
	"net/http"
)

type WebhookHandler struct{}

func NewWebhookHandler() *WebhookHandler {
	return &WebhookHandler{}
}

func (h *WebhookHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	return nil
}
