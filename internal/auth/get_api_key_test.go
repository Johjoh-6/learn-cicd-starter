package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name      string
		input     http.Header
		want      string
		wantError bool
	}{
		{
			name:      "valid bearer token",
			input:     http.Header{"Authorization": []string{"Bearer string"}},
			want:      "",
			wantError: true,
		},
		{
			name:      "malformed bearer token",
			input:     http.Header{"Authorization": []string{"Bearer"}},
			want:      "",
			wantError: true,
		},
		{
			name:      "no authorization header",
			input:     http.Header{},
			want:      "",
			wantError: true,
		},
		{
			name:      "invalid authorization header",
			input:     http.Header{"Authorization": []string{"Basic string"}},
			want:      "",
			wantError: true,
		},
		{
			name:      "Valid authorization ApiKey",
			input:     http.Header{"Authorization": []string{"ApiKey string"}},
			want:      "string",
			wantError: false,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			got, err := GetAPIKey(c.input)
			if c.wantError {
				if err == nil {
					t.Errorf("GetAPIKey(%v) should have returned an error", c.input)
				}
			} else {
				if err != nil {
					t.Errorf("GetAPIKey(%v) returned an error: %v", c.input, err)
				}
			}
			if got != c.want {
				t.Errorf("GetAPIKey(%v) == %q, want %q", c.input, got, c.want)
			}
		})
	}
}
