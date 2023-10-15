package validation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/schedule"
	"github.com/go-playground/validator/v10"
)

type ScheduleValidate struct {
	UserId            int    `validate:"required,numeric"`
	HealthcareAddress string `validate:"required"`
	Day               string `validate:"required"`
	TimeStart         string `validate:"required"`
	TimeEnd           string `validate:"required"`
}

func CoreToRegValSche(data schedule.Core) ScheduleValidate {
	return ScheduleValidate{
		UserId:            data.UserId,
		HealthcareAddress: data.HealthcareAddress,
		Day:               data.Day,
		TimeStart:         data.TimeStart,
		TimeEnd:           data.TimeEnd,
	}
}
func ScheCreateValidate(data schedule.Core) error {
	validate := validator.New()
	val := CoreToRegValSche(data)
	if err := validate.Struct(val); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			// vlderror := ""
			if e.Field() == "Password" && e.Value() != "" {
				// vlderror = fmt.Sprintf("%s is not %s", e.Value(), e.Tag())
				return errors.New(err.Error())
			}
			if e.Value() == "" {
				// vlderror = fmt.Sprintf("%s is %s", e.Field(), e.Tag())
				return errors.New(err.Error())
			} else {
				// vlderror = fmt.Sprintf("%s is not %s", e.Value(), e.Tag())
				return errors.New(err.Error())
			}
		}
	}
	return nil
}

func TimeValidation(time string) error {
	validate := validator.New()
	timeCheck := strings.Replace(time, ":", "", -1)
	if time != "" && strings.Contains(time, ":") {
		err := validate.Var(timeCheck, "numeric")
		if err != nil {
			// e := err.(validator.ValidationErrors)[0]
			// vlderror := fmt.Sprintf("%s is not %s", time, e.Tag())
			return errors.New(err.Error())
		}
	}
	return nil
}

func TimesValidation(timeStart, timeEnd string) error {
	err := TimeValidation(timeStart)
	if err != nil {
		return errors.New(err.Error())
	}
	err = TimeValidation(timeEnd)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func UpdateScheduleCheckValidation(data schedule.Core) error {
	validate := validator.New()
	if data.UserId == 0 && data.Day == "" && data.HealthcareAddress == "" && data.TimeStart == "" && data.TimeEnd == "" {
		return nil
	}
	if data.UserId != 0 {
		err := validate.Var(data.UserId, "numeric")
		if err != nil {
			// e := err.(validator.ValidationErrors)[0]
			// vlderror := fmt.Sprintf("UserId %d is not %s", data.UserId, e.Tag())
			return errors.New(err.Error())
		}
	}
	if data.TimeStart != "" {
		err := TimeValidation(data.TimeStart)
		if err != nil {
			return errors.New(err.Error())
		}
	}
	if data.TimeEnd != "" {
		err := TimeValidation(data.TimeEnd)
		if err != nil {
			return errors.New(err.Error())
		}
	}

	return nil
}

func TimeCheckerVal(timeStart, timeEnd string) error {
	strPemisah := ":"

	// Memisahkan hour dan minute time start
	tsHour := timeStart[:strings.Index(timeStart, strPemisah)]
	tsMin := timeStart[(strings.Index(timeStart, strPemisah) + 1):]
	if len(tsHour) < 2 {
		tsHour = fmt.Sprintf("0%s", tsHour)
	}
	if len(tsMin) < 2 {
		tsMin = fmt.Sprintf("0%s", tsMin)
	}

	// Memisahkan hour dan minute time end
	teHour := timeEnd[:strings.Index(timeEnd, strPemisah)]
	teMin := timeEnd[(strings.Index(timeEnd, strPemisah) + 1):]
	if len(teHour) < 2 {
		teHour = fmt.Sprintf("0%s", teHour)
	}
	if len(teMin) < 2 {
		teMin = fmt.Sprintf("0%s", teMin)
	}

	// cek apakah hour lebih dari 23 dan min lebih dari 59
	intTSHour, _ := strconv.Atoi(tsHour)
	intTEHour, _ := strconv.Atoi(teHour)
	intTSMin, _ := strconv.Atoi(tsMin)
	intTEMin, _ := strconv.Atoi(teMin)
	if int(intTSHour) > 23 || int(intTEHour) > 23 || int(intTEHour) < 0 || int(intTSHour) < 0 {
		return errors.New(config.TIME_ERR_FORMAT_HOUR)
	}
	if int(intTSMin) > 59 || int(intTEMin) > 59 || int(intTSMin) < 0 || int(intTEMin) < 0 {
		return errors.New(config.TIME_ERR_FORMAT_MINUTE)
	}

	// cek apakah time end lebih kecil dari time start
	intTimeStart, _ := strconv.Atoi(strings.Replace(timeStart, ":", "", -1))
	intTimeEnd, _ := strconv.Atoi(strings.Replace(timeEnd, ":", "", -1))
	if intTimeEnd <= intTimeStart {
		return errors.New(config.TIME_ERR_INVALID_TIME)
	}
	return nil
}
