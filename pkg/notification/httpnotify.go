package notification

import (
	"bytes"
	"net/http"

	"github.com/pkg/errors"
)

type simpleHttpNotify struct {
	Endpoint string
}

func (h *simpleHttpNotify) Notify(payload []byte) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, h.Endpoint, bytes.NewBuffer(payload))
	req.Header.Add("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		return errors.Wrap(err, "notify "+h.Endpoint+" failed")
	}

	return nil
}

func NewSimpleHttpNotify(endpoint string) *simpleHttpNotify {
	return &simpleHttpNotify{Endpoint: endpoint}
}
