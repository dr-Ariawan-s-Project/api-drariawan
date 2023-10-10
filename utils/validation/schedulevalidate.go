package validation

import (
	"errors"
	"strings"

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

func CoreToRegVal(data schedule.Core) ScheduleValidate {
	return ScheduleValidate{
		UserId:            data.UserId,
		HealthcareAddress: data.HealthcareAddress,
		Day:               data.Day,
		TimeStart:         data.TimeStart,
		TimeEnd:           data.TimeEnd,
	}
}
func CreateValidate(data schedule.Core) error {
	validate := validator.New()
	val := CoreToRegVal(data)
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
