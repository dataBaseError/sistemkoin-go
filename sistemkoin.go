package sistemkoin

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"net/url"
)

type SistemkoinClient struct {
	apiKey    string
	apiSecret string
}

func NewSistemkoinClient(
	apiKey string,
	apiSecret string,
) *SistemkoinClient {
	return &SistemkoinClient{
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}

func (s *SistemkoinClient) MakeRequest(
	httpClient *http.Client,
	method string, // get / post
	address string,
	params url.Values,
) (*http.Response, error) {
	headers := http.Header{}
	headers.Set("X-STK-ApiKey", s.apiKey)

	paramRaw := params.Encode()

	h := sha256.New()
	h.Write([]byte(s.apiSecret))
	h.Write([]byte(paramRaw))

	sig := h.Sum(nil)

	params.Set("signature", hex.EncodeToString(sig))

	req, err := http.NewRequest(
		method,
		address+"?"+params.Encode(),
		bytes.NewReader([]byte(params.Encode())),
	)
	if err != nil {
		return nil, err
	}

	req.Header = headers

	return httpClient.Do(req)
}
