package service

import (
	"fginostcal/model"
	"time"
)

type SparService struct {
}

func (ss *SparService) Calculate(params model.SparParams) int {
	startTime := params.StartTime
	endTime := params.EndTime
	totalSpar := fiftyDaysSignInReward(startTime, endTime)

	return totalSpar
}

func fiftyDaysSignInReward(start, end time.Time) int {
	days := getDiffDays(end, start)
	return 30 * (days / 50)
}

func getDiffDays(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}
