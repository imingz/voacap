package service

import (
	"context"
	"voacap/biz/dal/db"
	"voacap/biz/model/antennas"
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

// GetAntennas 获取所有天线信息
func (s *AntennaService) GetAntennas() ([]*common.Antenna, error) {
	antennas, err := db.QueryAllAntennas()
	if err != nil {
		return nil, err
	}

	result := make([]*common.Antenna, len(antennas))
	for i, antenna := range antennas {
		result[i] = &common.Antenna{
			AntennaID: antenna.AntennaID,
			Aname:     antenna.Aname,
			Afile:     antenna.Afile,
			AfbandMin: antenna.AfbandMin,
			AfbandMax: antenna.AfbandMax,
		}
	}

	return result, nil
}

// AddAntenna 添加天线
func (s *AntennaService) AddAntenna(req *antennas.AddAntennaRequest) error {
	return db.AddAntenna(db.Antenna{
		Aname:     req.Aname,
		Afile:     req.Afile,
		AfbandMin: req.AfbandMin,
		AfbandMax: req.AfbandMax,
	})
}

// UpdateAntennaById 更新天线信息
func (s *AntennaService) UpdateAntennaById(req *antennas.UpdateAntennaByIdRequest) error {
	return db.UpdateAntennaById(db.Antenna{
		AntennaID: req.AntennaId,
		Aname:     req.Aname,
		Afile:     req.Afile,
		AfbandMin: req.AfbandMin,
		AfbandMax: req.AfbandMax,
	})
}
