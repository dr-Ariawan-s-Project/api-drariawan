package helpers

import "fmt"

func GenerateCodeResponse(httpCode int, featureCode string, layerCode string) string {
	codeResponse := fmt.Sprintf("%d", httpCode) + "-" + featureCode + "-" + layerCode
	return codeResponse
}
