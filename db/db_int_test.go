// +build integration

package db

import (
	"context"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Initialise(context.Background())
	result := m.Run()
	os.Exit(result)
}
