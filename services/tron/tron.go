package tron

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Tron struct {
	base string
	key  string
}

func New(base, key string) *Tron {
	return &Tron{
		base: base,
		key:  key,
	}
}

func (tron *Tron) CheckTransaction(txID string) error {
	endpoint := fmt.Sprintf("%s/%s?hash=%s", tron.base, "api/transaction-info", txID)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("TRON-PRO-API-KEY", tron.key)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(req.Body).Decode(nil); err != nil {
		return err
	}
	// TODO: not completed yet.
	return nil
}
