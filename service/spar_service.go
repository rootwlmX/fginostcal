package service

import (
	"fginostcal/model"
	"time"
)

type SparService struct {
}

func (ss *SparService) Calculate(params model.SparParams) model.SparResponse {

	days := getDiffDays(params.EndTime, params.StartTime)

	sparResponse := model.SparResponse{
		Spars:  calculateTotalSpars(days),
		Coupon: calculateTotalCoupon(days),
	}

	return sparResponse
}

func calculateTotalSpars(days int) int {
	totalSpar := fiftyDaysSignInSpars(days)
	totalSpar += weeklyTaskSpars(days)
	totalSpar += weeklySignInSpars(days)

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

func calculateTotalCoupon(days int) int {
	totalCoupon := weeklySignInCoupons(days)
	return totalCoupon
}

func weeklySignInCoupons(days int) int {
	return days / 7
}

func getDiffDays(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}
