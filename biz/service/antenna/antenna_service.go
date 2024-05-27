package service

import (
	"context"
	"voacap/biz/dal/db"
	"voacap/biz/model/common"

	"github.com/cloudwego/hertz/pkg/app"
)

type AntennaService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewAntennaService create antenna service
func NewAntennaService(ctx context.Context, c *app.RequestContext) *AntennaService {
	return &AntennaService{ctx: ctx, c: c}
}

func (s *AntennaService) GetAntennas() ([]*common.Antennas, error) {
	antennas, err := db.QueryAllAntennas()
	if err != nil {
		return nil, err
	}

	result := make([]*common.Antennas, len(antennas))
	for i, antenna := range antennas {
		result[i] = &common.Antennas{
			AntennaID: antenna.AntennaID,
			Aname:     antenna.Aname,
			Afile:     antenna.Afile,
			AfbandMin: antenna.AfbandMin,
			AfbandMax: antenna.AfbandMax,
		}
	}

	return result, nil
}
