package service

import (
	"context"
	"voacap/biz/dal/db"
	"voacap/biz/model/common"

	"github.com/cloudwego/hertz/pkg/app"
)

type StationService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewStationService create StationService service
func NewStationService(ctx context.Context, c *app.RequestContext) *StationService {
	return &StationService{ctx: ctx, c: c}
}

// GetAntennas 获取所有天线信息
func (s *StationService) GetStation() ([]*common.Station, error) {
	stations, err := db.QueryAllStation()
	if err != nil {
		return nil, err
	}

	result := make([]*common.Station, len(stations))
	for i, station := range stations {
		result[i] = &common.Station{
			StationID:  station.StationID,
			Slatitude:  station.Slatitude,
			Slongitude: station.Slongitude,
			Sname:      station.Sname,
		}
	}

	return result, nil
}
