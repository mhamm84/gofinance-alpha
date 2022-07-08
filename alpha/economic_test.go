package alpha

import (
	"testing"
)

const baseUrl = "https://www.alphavantage.co/query"
const token = "MY_TOKEN"

var tests = []struct {
	name    string
	opts    *CpiOptions
	want    string
	isError bool
}{
	{"TestCpiEndpointCreateMonthly", &CpiOptions{Interval: Monthly}, baseUrl + "?function=CPI&interval=monthly", false},
	{"TestCpiEndpointCreateSemiAnnual", &CpiOptions{Interval: SemiAnnual}, baseUrl + "?function=CPI&interval=semiannual", false},
	{"TestCpiEndpointCreateUndefined", &CpiOptions{Interval: 999}, baseUrl + "?function=CPI&interval=monthly", false},
	{"TestCpiEndpointCreateNoOptions", nil, baseUrl + "?function=CPI&interval=monthly", false},
}

func TestCpiEndpointCreate(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createCpiEndpoint(baseUrl, token, tt.opts)
			if tt.isError && err != nil {

			} else {
				if got != tt.want {
					t.Errorf("createCpiEndpoint(%s) got %v, want %v", tt.name, got, tt.want)
				}
			}
		})
	}
}
