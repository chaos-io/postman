package postman

import (
	_ "embed"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

var (
	//go:embed testdata/collection.json
	col []byte
	//go:embed testdata/quiz.json
	quiz []byte
)

func TestParsePostman(t *testing.T) {
	test := func(file []byte) {
		col := &Collection{}
		err := jsoniter.Unmarshal(file, &col)
		assert.NoError(t, err)

		str, _ := jsoniter.MarshalToString(&col)
		assert.NotEmpty(t, str)
	}

	// test(col)
	test(quiz)
}
