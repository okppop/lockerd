package lockerd

import (
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestDefaultValueGenerator(t *testing.T) {
	g := defaultValueGenerator
	values := make([]string, 0, 100)
	prefix := getHostName() + "_"

	for range 100 {
		values = append(values, g())
	}

	for _, value := range values {
		value = strings.TrimPrefix(value, prefix)

		_, err := uuid.Parse(value)
		if err != nil {
			t.Error(err)
		}
	}
}
