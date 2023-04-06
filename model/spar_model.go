package model

import "time"

type SparParams struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

type SparResponse struct {
	Spars  int `json:"spars"`
	Coupon int `json:"coupon"`
}
