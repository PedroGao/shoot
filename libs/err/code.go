package err

import "net/http"

const (
	OKCode              = 0
	UnKnownCode         = 10000
	NoRouteMatchedCode  = 10001
	NotFoundCode        = 10002
	NoMethodMatchedCode = 20000
	UserNotFoundCode    = 30000
	ParamsErrCode       = 40000
	BookNotFoundCode    = 50000
	ForbiddenCode       = 60000
)

var (
	OK              = NewHTTPErr(OKCode, http.StatusOK, "OK!")
	UnKnown         = NewHTTPErr(UnKnownCode, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	NoRouteMatched  = NewHTTPErr(NoRouteMatchedCode, http.StatusNotFound, "sorry,there is no route the url matched")
	NotFound        = NewHTTPErr(NotFoundCode, http.StatusNotFound, "not found")
	NoMethodMatched = NewHTTPErr(NoMethodMatchedCode, http.StatusNotFound, "sorry,there is no methods the action matched")
	UserNotFound    = NewHTTPErr(UserNotFoundCode, http.StatusNotFound, "sorry,no user found")
	ParamsErr       = NewHTTPErr(ParamsErrCode, http.StatusBadRequest, "params err")
	BookNotFound    = NewHTTPErr(BookNotFoundCode, http.StatusNotFound, "sorry,no book found")
	Forbidden       = NewHTTPErr(ForbiddenCode, http.StatusOK, "forbidden！you can't access the resource")
)
