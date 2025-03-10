package applicationError

import "net/http"

var defaultMessageMap = map[int]string{
	http.StatusBadRequest:                   "Bad Request",
	http.StatusUnauthorized:                 "Unauthorized",
	http.StatusPaymentRequired:              "Payment Required",
	http.StatusForbidden:                    "Forbidden",
	http.StatusNotFound:                     "Not Found",
	http.StatusMethodNotAllowed:             "Method Not Allowed",
	http.StatusNotAcceptable:                "Not Acceptable",
	http.StatusProxyAuthRequired:            "Proxy Authentication Required",
	http.StatusRequestTimeout:               "Request Timeout",
	http.StatusConflict:                     "Conflict",
	http.StatusGone:                         "Gone",
	http.StatusLengthRequired:               "Length Required",
	http.StatusPreconditionFailed:           "Precondition Failed",
	http.StatusRequestEntityTooLarge:        "Request Entity Too Large",
	http.StatusRequestURITooLong:            "Request-URI Too Long",
	http.StatusUnsupportedMediaType:         "Unsupported Media Type",
	http.StatusRequestedRangeNotSatisfiable: "Requested Range Not Satisfiable",
	http.StatusExpectationFailed:            "Expectation Failed",
	http.StatusInternalServerError:          "Internal Server Error",
	http.StatusNotImplemented:               "Not Implemented",
	http.StatusBadGateway:                   "Bad Gateway",
	http.StatusServiceUnavailable:           "Service Unavailable",
	http.StatusGatewayTimeout:               "Gateway Timeout",
	http.StatusHTTPVersionNotSupported:      "HTTP Version Not Supported",
}

func getDefaultErrorMessage(status int, msg string) string {
	if defaultMessage, ok := defaultMessageMap[status]; ok {
		return defaultMessage
	}

	return msg
}
