package clock

import(
	"time"
	"errors"
	"fmt"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/elfsternberg/timeofday/restapi/operations"
	"github.com/elfsternberg/timeofday/models"
	"github.com/go-openapi/swag"
)



func getTimeOfDay(tz *string) (*string, error) {
	defaultTZ := "UTC"
	
	t := time.Now()
	if tz == nil {
		tz = &defaultTZ
	}

	utc, err := time.LoadLocation(*tz)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Time zone not found: %s", *tz))
	}

	thetime := t.In(utc).String()
	return &thetime, nil
}


func GetClock(params operations.ClockGetParams) middleware.Responder {
	var tz *string = nil

	if (params.Timezone != nil) {
		tz = params.Timezone
	}
	
	thetime, err := getTimeOfDay(tz)
	if err != nil {
		return operations.NewClockGetNotFound().WithPayload(
			&models.ErrorResponse {
				int32(operations.ClockGetNotFoundCode),
				swag.String(fmt.Sprintf("%s", err)),
			})
	}

	return operations.NewClockGetOK().WithPayload(
		&models.Timeofday{
			Timeofday: *thetime,
		})
}


func PostClock(params operations.ClockPostParams) middleware.Responder {
	var tz *string = nil

	if (params.Timezone != nil) {
		tz = params.Timezone.Timezone
	}

	thetime, err := getTimeOfDay(tz)
	if err != nil {
		return operations.NewClockPostNotFound().WithPayload(
			&models.ErrorResponse {
				int32(operations.ClockPostNotFoundCode),
				swag.String(fmt.Sprintf("%s", err)),
			})
	}

	return operations.NewClockPostOK().WithPayload(
		&models.Timeofday{
			Timeofday: *thetime,
		})
}
