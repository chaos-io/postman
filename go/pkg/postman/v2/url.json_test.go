package postman

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestUrlQueryParamCodec_Decode(t *testing.T) {

	codec := &Url_QueryParam{
		Key:         "foo",
		Value:       "bar",
		Disable:     false,
		Description: nil,
	}
	str, err := jsoniter.ConfigDefault.MarshalToString(codec)
	assert.NoError(t, err)
	assert.NotEmpty(t, str)
}

func TestUrlQueryParamCodec_Encode(t *testing.T) {
	str := `["foo", "bar"]`
	var query []*Url_QueryParam
	err := jsoniter.UnmarshalFromString(str, &query)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(query))

	str = `[{"key":"param","value":"val"},"foo",{"value":"bar"}]`
	query = []*Url_QueryParam{}
	err = jsoniter.UnmarshalFromString(str, &query)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(query))
}
