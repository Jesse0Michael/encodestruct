package encodestruct

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
)

// Encode will convert the passed in value to a []byte using the gob encoder and base64 encode the bytes
func Encode(v interface{}) string {
	var buffer bytes.Buffer
	_ = gob.NewEncoder(&buffer).Encode(v)
	return base64.StdEncoding.EncodeToString(buffer.Bytes())
}

// Decode will convert a base64 encoded string into bytes to decode into the value passed in the parameters using the gob decoder
func Decode(s string, v interface{}) {
	b, _ := base64.StdEncoding.DecodeString(s)
	_ = gob.NewDecoder(bytes.NewReader(b)).Decode(v)
}
