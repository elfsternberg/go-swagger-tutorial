package timeofday

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
	t := time.Now()
	utc, err := time.LoadLocation(*tz)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Time zone not found: %s", *tz))
	}

	thetime := t.In(utc).String()
	return &thetime, nil
}


func GetTime(timezone *Timezone) func(operations.TimeGetParams) middleware.Responder{
	defaultTZ := timezone.Timezone
	
	return func(params operations.TimeGetParams) middleware.Responder {
		var tz *string = &defaultTZ
		if (params.Timezone != nil) {
			tz = params.Timezone
		}
		
		thetime, err := getTimeOfDay(tz)
		if err != nil {
			return operations.NewTimeGetNotFound().WithPayload(
				&models.ErrorResponse {
					int32(operations.TimeGetNotFoundCode),
					swag.String(fmt.Sprintf("%s", err)),
				})
		}
		
		return operations.NewTimeGetOK().WithPayload(
			&models.Timeofday{
				Timeofday: *thetime,
			})
	}
}

func PostTime(timezone *Timezone) func(operations.TimePostParams) middleware.Responder{
	defaultTZ := timezone.Timezone

	return func(params operations.TimePostParams) middleware.Responder {
		var tz *string = &defaultTZ
		if (params.Timezone != nil) {
			tz = params.Timezone.Timezone
		}
		
		thetime, err := getTimeOfDay(tz)
		if err != nil {
			return operations.NewTimePostNotFound().WithPayload(
				&models.ErrorResponse {
					int32(operations.TimePostNotFoundCode),
					swag.String(fmt.Sprintf("%s", err)),
				})
		}
		
		return operations.NewTimePostOK().WithPayload(
			&models.Timeofday{
				Timeofday: *thetime,
			})
	}
}
