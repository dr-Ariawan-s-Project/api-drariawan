package service

import (
	"errors"
	"strings"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/schedule"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/helpers"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/validation"
)

type ScheduleService struct {
	qry schedule.ScheduleData
}

func New(sd schedule.ScheduleData) schedule.ScheduleService {
	return &ScheduleService{
		qry: sd,
	}
}

// for pagination
// GetPagination implements schedule.ScheduleService.
func (ss *ScheduleService) GetPagination(page int, perPage int) (map[string]any, error) {
	totalRows, err := ss.qry.CountByFilter()
	response := map[string]any{
		"page":          0,
		"limit":         0,
		"total_pages":   0,
		"total_records": 0,
	}
	if err != nil {
		return response, err
	}
	paginationRes := helpers.CountPagination(totalRows, page, perPage)
	response["page"] = paginationRes.Page
	response["limit"] = paginationRes.Limit
	response["total_pages"] = paginationRes.TotalPages
	if perPage == 0 {
		// di schedule feature, jika limit tidak diinputkan maka tampilkan semua datanya.
		response["limit"] = 0
		response["total_pages"] = 1
	}
	response["total_records"] = paginationRes.TotalRecords
	return response, nil
}

// Create implements schedule.ScheduleService.
func (ss *ScheduleService) Create(data schedule.Core, role string) error {
	if strings.ToLower(role) != config.VAL_DokterAccess && strings.ToLower(role) != config.VAL_AdminAccess && strings.ToLower(role) != config.VAL_SuperAdminAccess {
		return errors.New(config.VAL_Unauthorized)
	}
	err := validation.TimesValidation(data.TimeStart, data.TimeEnd)
	if err != nil {
		return errors.New(err.Error())
	}
	err = validation.CreateScheduleValidate(data)
	if err != nil {
		return errors.New(err.Error())
	}
	err = validation.TimeCheckerValidate(data.TimeStart, data.TimeEnd)
	if err != nil {
		return errors.New(err.Error())
	}
	err = ss.qry.Create(data)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// Update implements schedule.ScheduleService.
func (ss *ScheduleService) Update(id int, data schedule.Core, role string) error {

	err := validation.TimesValidation(data.TimeStart, data.TimeEnd)
	if err != nil {
		return errors.New(err.Error())
	}
	err = validation.UpdateScheduleCheckValidation(data)
	if err != nil {
		return errors.New(err.Error())
	}
	err = ss.qry.Update(id, data)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// Delete implements schedule.ScheduleService.
func (ss *ScheduleService) Delete(id int, role string) error {

	err := ss.qry.Delete(id)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// GetAll implements schedule.ScheduleService.
func (ss *ScheduleService) GetAll(role string, page int, perPage int) ([]schedule.Core, error) {
	// log.Println(strings.ToLower(role), config.VAL_SusterAccess)
	if strings.ToLower(role) != config.VAL_SusterAccess && strings.ToLower(role) != config.VAL_DokterAccess && strings.ToLower(role) != config.VAL_PatientAccess && strings.ToLower(role) != config.VAL_SuperAdminAccess && strings.ToLower(role) != config.VAL_AdminAccess {
		return []schedule.Core{}, errors.New(config.VAL_Unauthorized)
	}
	// if perPage <= 0 {
	// 	perPage = 10
	// }
	if page <= 0 {
		page = 1
	}
	offset := (page * perPage) - perPage

	if offset < 0 {
		offset = 0
	}
	// log.Println("SERVICE OK")
	res, err := ss.qry.GetAll(offset, perPage)
	if err != nil {
		return []schedule.Core{}, errors.New(err.Error())
	}
	return res, nil
}
