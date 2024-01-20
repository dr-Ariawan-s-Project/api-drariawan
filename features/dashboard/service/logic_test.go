package service

import (
	"errors"
	"testing"
	"time"

	"github.com/dr-ariawan-s-project/api-drariawan/features/dashboard"
	"github.com/dr-ariawan-s-project/api-drariawan/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetDashboardStatistics(t *testing.T) {
	dashboardRepo := new(mocks.DashboardData)
	t.Run("success get dashboard data", func(t *testing.T) {
		dashboardRepo.On("CountQuestionerAttempt").Return(1, nil).Once()
		dashboardRepo.On("CountAttemptByStatusAssessment", "validated").Return(1, nil).Once()
		timeNow := time.Now()
		dashboardRepo.On("CountAttemptByMonth", int(timeNow.Month())).Return(1, nil).Once()
		dashboardRepo.On("CountAllPatient").Return(10, nil).Once()

		dashboardService := New(dashboardRepo)
		response, err := dashboardService.GetDashboardStatistics()
		assert.NoError(t, err)
		assert.Equal(t, 10, response.AllPatient)
		dashboardRepo.AssertExpectations(t)
	})

	t.Run("error get data CountQuestionerAttempt", func(t *testing.T) {
		dashboardRepo.On("CountQuestionerAttempt").Return(0, errors.New("error")).Once()

		dashboardService := New(dashboardRepo)
		_, err := dashboardService.GetDashboardStatistics()
		assert.Error(t, err)
		dashboardRepo.AssertExpectations(t)
	})
	t.Run("error get data CountAttemptByStatusAssessment", func(t *testing.T) {
		dashboardRepo.On("CountQuestionerAttempt").Return(1, nil).Once()
		dashboardRepo.On("CountAttemptByStatusAssessment", "validated").Return(0, errors.New("error")).Once()

		dashboardService := New(dashboardRepo)
		_, err := dashboardService.GetDashboardStatistics()
		assert.Error(t, err)
		dashboardRepo.AssertExpectations(t)
	})
	t.Run("error get data CountAttemptByMonth", func(t *testing.T) {
		dashboardRepo.On("CountQuestionerAttempt").Return(1, nil).Once()
		dashboardRepo.On("CountAttemptByStatusAssessment", "validated").Return(1, nil).Once()
		timeNow := time.Now()
		dashboardRepo.On("CountAttemptByMonth", int(timeNow.Month())).Return(0, errors.New("error")).Once()

		dashboardService := New(dashboardRepo)
		_, err := dashboardService.GetDashboardStatistics()
		assert.Error(t, err)
		dashboardRepo.AssertExpectations(t)
	})
	t.Run("error get data CountAllPatient", func(t *testing.T) {
		dashboardRepo.On("CountQuestionerAttempt").Return(1, nil).Once()
		dashboardRepo.On("CountAttemptByStatusAssessment", "validated").Return(1, nil).Once()
		timeNow := time.Now()
		dashboardRepo.On("CountAttemptByMonth", int(timeNow.Month())).Return(1, nil).Once()
		dashboardRepo.On("CountAllPatient").Return(0, errors.New("error")).Once()

		dashboardService := New(dashboardRepo)
		_, err := dashboardService.GetDashboardStatistics()
		assert.Error(t, err)
		dashboardRepo.AssertExpectations(t)
	})

}

func TestGetQuestionerAttemptPerMonth(t *testing.T) {
	dashboardRepo := new(mocks.DashboardData)
	response := []dashboard.DashboardAttemptCore{
		{
			Month: "januari",
			Count: 10,
		},
		{
			Month: "februari",
			Count: 10,
		},
		{
			Month: "maret",
			Count: 10,
		},
		{
			Month: "april",
			Count: 10,
		},
		{
			Month: "mei",
			Count: 10,
		},
		{
			Month: "juni",
			Count: 10,
		},
		{
			Month: "juli",
			Count: 10,
		},
		{
			Month: "agustus",
			Count: 10,
		},
		{
			Month: "september",
			Count: 10,
		},
		{
			Month: "oktober",
			Count: 10,
		},
		{
			Month: "november",
			Count: 10,
		},
		{
			Month: "desember",
			Count: 10,
		},
	}
	t.Run("test success get data", func(t *testing.T) {
		dashboardRepo.On("CountQuestionerAttemptPerMonth").Return(response, nil).Once()
		dashboardService := New(dashboardRepo)
		response, err := dashboardService.GetQuestionerAttemptPerMonth()
		assert.NoError(t, err)
		assert.Equal(t, 1, len(response))
	})
}
