package alpha

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const baseUrl = "https://www.alphavantage.co/query"
const token = "MY_TOKEN"

var tests = []struct {
	name     string
	baseUrl  string
	function string
	token    string
	opts     *Options
	want     string
	isError  bool
}{
	{"TestCpiEndpointCreateMonthly", baseUrl, "CPI", token, &Options{Interval: Monthly}, strings.TrimSpace(baseUrl + fmt.Sprintf("?apikey=%s&function=%s&interval=%s", token, "CPI", "monthly")), false},
	{"TestCpiEndpointCreateSemiAnnual", baseUrl, "CPI", token, &Options{Interval: SemiAnnual}, strings.TrimSpace(baseUrl + fmt.Sprintf("?apikey=%s&function=%s&interval=%s", token, "CPI", "semiannual")), false},
	{"TestCpiEndpointCreateIntervalUnknown", baseUrl, "CPI", token, &Options{Interval: 99}, strings.TrimSpace(baseUrl + fmt.Sprintf("?apikey=%s&function=%s&interval=%s", token, "CPI", "unknown")), false},
	{"TestCpiEndpointCreateNoOptions", baseUrl, "CPI", token, nil, strings.TrimSpace(baseUrl + fmt.Sprintf("?apikey=%s&function=%s", token, "CPI")), false},
	{"TestCpiEndpointCreateNoFunction", baseUrl, "", token, nil, "", true},
	{"TestCpiEndpointCreateNoToken", baseUrl, "", token, nil, "", true},
	{"TestCpiEndpointCreateNoBaseUrl", "", "CPI", token, nil, "", true},
}

func TestCpiEndpointCreate(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createEndpoint(tt.baseUrl, tt.token, tt.function, tt.opts)
			if tt.isError && err == nil {
				t.Errorf("expected error but did not get one")
			} else {
				assert.Equal(t, tt.want, got, "endpoints mismatch")
				if got != tt.want {
					t.Errorf("createCpiEndpoint(%s) got %v, want %v", tt.name, got, tt.want)
				}
			}
		})
	}
}
