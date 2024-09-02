package controller

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gitlab.com/DeveloperDurp/DurpAPI/pkg/dadjoke"
	"testing"
)

func Test(t *testing.T) {

	ctx := context.Background()

	request := testcontainers.ContainerRequest{
		Image: "postgres:16",
		Env: map[string]string{
			"POSTGRES_USER":     "user",
			"POSTGRES_PASSWORD": "secret",
			"POSTGRES_DB":       "testdb",
		},
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor: wait.ForLog("database system is ready to accept connections").
			WithOccurrence(2),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: request,
		Started:          true,
	})
	if err != nil {
		t.Fatal("failed to start container:", err)
	}

	endpoint, err := container.Endpoint(ctx, "")
	if err != nil {
		t.Fatal("failed to get endpoint:", err)
	}

	connURI := fmt.Sprintf("postgres://user:secret@%s/testdb?sslmode=disable", endpoint)

	db, err := connectDB(connURI)
	assert.NoError(t, err)

	err = db.AutoMigrate(&dadjoke.DadJoke{})
	assert.NoError(t, err)
}
