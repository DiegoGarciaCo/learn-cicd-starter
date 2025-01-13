package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		headers        http.Header
		expectedAPIKey string
		expectError    error
	}{
		{
			headers:        http.Header{"Authorization": []string{"ApiKey dsjgvgeor32049349hjbse0j390"}},
			expectedAPIKey: "dsjgvgeor32049349hjbse0j390",
			expectError:    nil,
		},
		{
			headers:        http.Header{"Authorization": []string{"ApiKey djv9j024p9343jf23"}},
			expectedAPIKey: "djv9j024p9343jf23",
			expectError:    nil,
		},
		{
			headers:        http.Header{"Authorization": []string{"Bearer dsjgvgeor32049349hjbse0j390"}},
			expectedAPIKey: "",
			expectError:    errors.New("malformed authorization header"),
		},
		{
			headers:        http.Header{"Content-Type": []string{"application/json"}},
			expectedAPIKey: "",
			expectError:    ErrNoAuthHeaderIncluded,
		},
	}

	for _, tt := range tests {
		t.Run(tt.headers.Get("Authorization"), func(t *testing.T) {
			apiKey, err := GetAPIKey(tt.headers)

			// Check for expected error
			if err != nil && err.Error() != tt.expectError.Error() {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}

			// Check if the API key matches the expected value
			if apiKey != tt.expectedAPIKey {
				t.Errorf("expected API key: %s, got: %s", tt.expectedAPIKey, apiKey)
			}
		})
	}
}
