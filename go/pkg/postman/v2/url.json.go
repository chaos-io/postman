package postman

import (
	"fmt"
	"unsafe"

	"github.com/chaos-io/chaos/core"
	jsoniter "github.com/json-iterator/go"
)

func init() {
	core.RegisterJSONTypeDecoder("postman.Url_QueryParam", &UrlQueryParamCodec{})
	core.RegisterJSONTypeEncoder("postman.Url_QueryParam", &UrlQueryParamCodec{})
}

type UrlQueryParamCodec struct{}

func (codec *UrlQueryParamCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	val := iter.ReadAny()
	param := (*Url_QueryParam)(ptr)
	switch val.ValueType() {
	case jsoniter.StringValue:
		param.Value = val.ToString()
	case jsoniter.ObjectValue:
		param.Key = val.Get("key").ToString()
		param.Value = val.Get("value").ToString()
		param.Disable = val.Get("disable").ToBool()
	default:
		iter.ReportError("Url_QueryParam", fmt.Sprintf("invalid Url_QueryParam key value, original type: %d", val.ValueType()))
	}
}

func (codec *UrlQueryParamCodec) IsEmpty(ptr unsafe.Pointer) bool {
	e := (*Url_QueryParam)(ptr)
	return e == nil
}

func (codec *UrlQueryParamCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	param := (*Url_QueryParam)(ptr)
	if len(param.Key) > 0 {
		stream.WriteObjectStart()
		stream.WriteObjectField("key")
		stream.WriteString(param.Key)
		stream.WriteObjectField("value")
		stream.WriteString(param.Value)
		stream.WriteObjectField("disable")
		stream.WriteBool(param.Disable)
		stream.WriteObjectEnd()
	} else {
		stream.WriteString(param.Value)
	}
}
