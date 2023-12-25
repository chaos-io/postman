package postman

import (
	_ "embed"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

//go:embed testdata/collection.json
var col1 []byte

func TestParsePostman(t *testing.T) {
	col := &Collection{}
	err := jsoniter.Unmarshal(col1, &col)
	assert.NoError(t, err)
}
