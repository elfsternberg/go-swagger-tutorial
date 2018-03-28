package clock

import(
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/elfsternberg/timeofday/restapi/operations"
)

func GetClock(params operations.ClockGetParams) middleware.Responder {
	return middleware.NotImplemented("operation .ClockGet has not yet been implemented")
}

func PostClock(params operations.ClockPostParams) middleware.Responder {
	return middleware.NotImplemented("operation .ClockPost has not yet been implemented")
}
