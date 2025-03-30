package applicationError

import "net/http"

var defaultMessageMap = map[int]string{
	// 400
	http.StatusBadRequest:                   "ERRC400",
	http.StatusUnauthorized:                 "ERRC401",
	http.StatusPaymentRequired:              "ERRC402",
	http.StatusForbidden:                    "ERRC403",
	http.StatusNotFound:                     "ERRC404",
	http.StatusMethodNotAllowed:             "ERRC405",
	http.StatusNotAcceptable:                "ERRC406",
	http.StatusProxyAuthRequired:            "ERRC407",
	http.StatusRequestTimeout:               "ERRC408",
	http.StatusConflict:                     "ERRC409",
	http.StatusGone:                         "ERRC410",
	http.StatusLengthRequired:               "ERRC411",
	http.StatusPreconditionFailed:           "ERRC412",
	http.StatusRequestEntityTooLarge:        "ERRC413",
	http.StatusRequestURITooLong:            "ERRC414",
	http.StatusUnsupportedMediaType:         "ERRC415",
	http.StatusRequestedRangeNotSatisfiable: "ERRC416",
	http.StatusExpectationFailed:            "ERRC417",
	// 500
	http.StatusInternalServerError:     "ERRC500",
	http.StatusNotImplemented:          "ERRC501",
	http.StatusBadGateway:              "ERRC502",
	http.StatusServiceUnavailable:      "ERRC503",
	http.StatusGatewayTimeout:          "ERRC504",
	http.StatusHTTPVersionNotSupported: "ERRC505",
}

func getDefaultErrorCode(status int) string {
	if errorCode, ok := defaultMessageMap[status]; ok {
		return errorCode
	}

	return "ERRC500"
}
