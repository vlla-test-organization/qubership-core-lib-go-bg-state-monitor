package consul

import "encoding/base64"

type KVInfo struct {
	Key         string      `json:"Key"`
	Value       Base64Value `json:"Value"`
	ModifyIndex int         `json:"ModifyIndex"`
}

type Base64Value string

func MarshalBase64Value(value string) Base64Value {
	return Base64Value(base64.StdEncoding.EncodeToString([]byte(value)))
}

func (b64v *Base64Value) Unmarshal() (string, error) {
	v, err := base64.StdEncoding.DecodeString(string(*b64v))
	return string(v), err
}
