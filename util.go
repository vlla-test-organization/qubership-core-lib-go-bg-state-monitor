package blue_green_state_monitor_go

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func GetModifyIndex(response *http.Response) (string, error) {
	indexHeaderVal := response.Header["X-Consul-Index"]
	if indexHeaderVal == nil || len(indexHeaderVal) == 0 {
		return "", fmt.Errorf("X-Consul-Index header not provided")
	} else {
		return indexHeaderVal[0], nil
	}
}

func ToConsulTTL(ttl time.Duration) string {
	return strings.ReplaceAll(ttl.String(), "PT", "")
}

func ToConsulTTLAsSeconds(seconds int64) string {
	return fmt.Sprintf("%ds", seconds)
}
