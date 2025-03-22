package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
)

type WebhookHandler struct{}

func NewWebhookHandler() *WebhookHandler {
	return &WebhookHandler{}
}

func (h *WebhookHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	// Infer type by headers sent
	signature := r.Header.Get("x-hub-signature-256")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	mac := hmac.New(sha256.New, []byte("belekas"))
	mac.Write(body)
	expectedSig := mac.Sum(nil)
	fmt.Println("expected:", hex.EncodeToString(expectedSig))
	fmt.Println("received:", signature)
	// hmac.Equal([]byte(expectedSignature), []byte(signature))
	return nil
}
