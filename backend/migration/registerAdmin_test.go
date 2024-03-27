package migration

import (
	"os"
	"testing"
)

func TestCreateAdminUser(t *testing.T) {
	os.Setenv("ADMIN_USERNAME", "admin")
	os.Setenv("ADMIN_PASSWORD", "admin123")

	err := CreateAdminUser()

	if err != nil {
		t.Errorf("Failed to create admin user: %v", err)
	}
}
