package postman

import (
	"fmt"
	"unsafe"

	"github.com/chaos-io/chaos/core"
	jsoniter "github.com/json-iterator/go"
)

func init() {
	core.RegisterJSONTypeDecoder("postman.Description", &DescriptionCodec{})
}

type DescriptionCodec struct{}

func (codec *DescriptionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	val := iter.ReadAny()

	e := (*Description)(ptr)
	if val.ValueType() == jsoniter.StringValue {
		e.Content = val.ToString()
	} else if val.ValueType() == jsoniter.ObjectValue {
		e.Type = val.Get("type").ToString()
		e.Content = val.Get("content").ToString()
		e.Version = val.Get("version").ToString()
	} else {
		iter.ReportError("Description", fmt.Sprintf("invalid Description type value, original type: %d", val.ValueType()))
	}
}

func (codec *DescriptionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	e := (*Description)(ptr)
	return e == nil
}
