package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name        string
		envs        map[string]string
		expect      *Config
		shouldPanic bool
	}{
		{
			name: "All variables set",
			envs: map[string]string{
				"API_PORT":    "8888",
				"PROJECT_ID":  "cloudrun-go-handson-396810",
				"INSTANCE_ID": "test-instance",
				"DB_ID":       "test-db",
				"CF":          "/api/cloudrun-go-handson-396810-d6478fd2d46f.json",
				"STAGE":       "local",
			},
			expect: &Config{
				Port:       "8888",
				ProjectID:  "cloudrun-go-handson-396810",
				InstanceID: "test-instance",
				DatabaseID: "test-db",
				CF:         "/api/cloudrun-go-handson-396810-d6478fd2d46f.json",
				Stage:      "local",
			},
			shouldPanic: false,
		},
		{
			name: "No variables set",
			envs: map[string]string{},
			expect: &Config{
				Port:       "",
				ProjectID:  "",
				InstanceID: "",
				DatabaseID: "",
				CF:         "",
				Stage:      "",
			},
			shouldPanic: false,
		},
		// Add more test cases as needed.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variables
			for key, value := range tt.envs {
				os.Setenv(key, value)
			}

			// Capture panic if any
			defer func() {
				r := recover()
				if (r != nil) != tt.shouldPanic {
					t.Errorf("NewConfig() panic = %v, shouldPanic = %v", r, tt.shouldPanic)
				}
			}()

			// Execute
			config := NewConfig()

			// Validate
			assert.Equal(t, tt.expect, config)
		})
	}
}
