package alpha

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetAndDecode creates a new HTTP request with a Context using the passed in endpoint url
// The body is then JSON decoded to the dst pointer passed in
func GetAndDecode(ctx context.Context, httpClient *http.Client, endpoint string, dst interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return err
	}

	httpRes, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		if tempErr := httpRes.Body.Close(); tempErr != nil {
			err = tempErr
		}
	}()

	decoder := json.NewDecoder(httpRes.Body)
	err = decoder.Decode(dst)
	if err != nil {
		return err
	}
	return nil
}
