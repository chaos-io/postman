package postman

import (
	"fmt"
	"unsafe"

	"github.com/chaos-io/chaos/core"
	jsoniter "github.com/json-iterator/go"
)

func init() {
	core.RegisterJSONTypeDecoder("postman.Url_QueryParam", &UrlQueryParamCodec{})
}

type UrlQueryParamCodec struct{}

func (codec *UrlQueryParamCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	val := iter.ReadAny()
	e := (*Url_QueryParam)(ptr)

	switch val.ValueType() {
	case jsoniter.StringValue:
		e.Value = val.ToString()
	case jsoniter.ObjectValue:
		e.Key = val.Get("key").ToString()
		e.Value = val.Get("value").ToString()
		e.Disable = val.Get("disable").ToBool()
		// e.Description = val.Get("description").ToVal()
	default:
		iter.ReportError("Url_QueryParam", fmt.Sprintf("invalid Url_QueryParam key value, original type: %d", val.ValueType()))
	}
}

func (codec *UrlQueryParamCodec) IsEmpty(ptr unsafe.Pointer) bool {
	e := (*Description)(ptr)
	return e == nil
}

func (codec *UrlQueryParamCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	e := (*Url_QueryParam)(ptr)
	if len(e.Key) > 0 {
		stream.WriteObjectStart()
		stream.WriteObjectField("key")
		stream.WriteString(e.Key)
		stream.WriteObjectField("value")
		stream.WriteString(e.Value)
		stream.WriteObjectField("disable")
		stream.WriteBool(e.Disable)
		stream.WriteObjectEnd()
	} else {
		stream.WriteString(e.Value)
	}
}
