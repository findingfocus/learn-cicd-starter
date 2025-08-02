package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		wantErr     bool
		expectedErr error
	}{
		{
			name: "valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey abc123"},
			},
			expectedKey: "abc123",
			wantErr:     false,
		},
		{
			name:        "missing authorization header",
			headers:     http.Header{},
			expectedKey: "",
			wantErr:     true,
			expectedErr: ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed header - wrong prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer abc123"},
			},
			expectedKey: "",
			wantErr:     true,
		},
		{
			name: "malformed header - missing key",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedKey: "",
			wantErr:     true,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.expectedErr != nil && err != tt.expectedErr {
				t.Errorf("GetAPIKey() error = %v, expectedErr %v", err, tt.expectedErr)
				return
			}

			if key != tt.expectedKey {
				t.Errorf("GetAPIKey() = %v, want %v", key, tt.expectedKey)
			}
		})
	}
}
