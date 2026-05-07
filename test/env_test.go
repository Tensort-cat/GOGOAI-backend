package test

import (
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	key := os.Getenv("OPENAI_API_KEY")
	modelName := os.Getenv("OPENAI_MODEL_NAME")
	baseURL := os.Getenv("OPENAI_BASE_URL")

	t.Logf("OPENAI_API_KEY: %s", key)
	t.Logf("OPENAI_MODEL_NAME: %s", modelName)
	t.Logf("OPENAI_BASE_URL: %s", baseURL)
}
