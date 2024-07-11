package ton

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TON struct {
	base string
	key  string
}

func New(base, key string) *TON {
	return &TON{
		base: base,
		key:  key,
	}
}

func (ton *TON) CheckTransaction(txID string) error {
	endpoint := fmt.Sprintf("%s/%s/%s", ton.base, "v2/blockchain/transactions", txID)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("TRON-PRO-API-KEY", ton.key)

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
