package hooks

import (
	"net/http"
)

type HookHandler interface {
	Handle()
}

func InferNewHookFromHeaders(r *http.Request) (HookHandler, error) {
	return nil, nil
}
