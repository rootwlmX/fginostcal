package service

import (
	"fginostcal/dao"
	"fginostcal/engine"
	"fginostcal/model"
	"time"
)

type SparService struct {
}

func (ss *SparService) Calculate(params model.SparParams) model.SparResponse {
	startTime := params.StartTime
	endTime := params.EndTime
	days := getDiffDays(endTime, startTime)

	sparResponse := model.SparResponse{
		Spars:  calculateTotalSpars(startTime, endTime, days),
		Coupon: calculateTotalCoupons(startTime, endTime, days),
	}

	return sparResponse
}

func calculateTotalSpars(start, end time.Time, days int) int {

	totalSpar := fiftyDaysSignInSpars(days)
	totalSpar += weeklyTaskSpars(days)
	totalSpar += weeklySignInSpars(days)
	totalSpar += getEventSpars(start, end)

	return totalSpar
}

func fiftyDaysSignInSpars(days int) int {
	return 30 * (days / 50)
}

func weeklyTaskSpars(days int) int {
	return 3 * (days / 7)
}

func weeklySignInSpars(days int) int {
	spars := 6 * (days / 7)
	extraDays := days % 7
	if extraDays >= 6 {
		spars += 4
	} else if extraDays >= 4 {
		spars += 2
	} else if extraDays >= 2 {
		spars += 1
	}

	return spars
}

func getEventSpars(start, end time.Time) int {
	spars := 0
	eventDao := dao.EventDao{DbEngine: engine.GetOrmEngine()}
	events := eventDao.GetEventsInDayRange(start, end)
	if events == nil {
		return spars
	}

	for _, event := range *events {
		spars += event.Spar
	}

	return spars
}

func calculateTotalCoupons(start, end time.Time, days int) int {
	totalCoupon := weeklySignInCoupons(days)
	totalCoupon += getEventCoupons(start, end)
	return totalCoupon
}

func weeklySignInCoupons(days int) int {
	return days / 7
}

func getEventCoupons(start, end time.Time) int {
	coupons := 0
	eventDao := dao.EventDao{DbEngine: engine.GetOrmEngine()}
	events := eventDao.GetEventsInDayRange(start, end)
	if events == nil {
		return coupons
	}
	for _, event := range *events {
		coupons += event.Coupon
	}

	return coupons
}

func getDiffDays(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}
