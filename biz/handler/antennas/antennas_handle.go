// Code generated by hertz generator.

package antennas

import (
	"context"

	antennas "voacap/biz/model/antennas"
	service "voacap/biz/service/antenna"
	"voacap/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetAntennas .
// @router /getAntennas [GET]
func GetAntennas(ctx context.Context, c *app.RequestContext) {
	var err error
	var req antennas.GetAntennasReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, antennas.GetAntennasResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	data, err := service.NewAntennaService(ctx, c).GetAntennas()

	resp := utils.BuildBaseResp(err)
	c.JSON(consts.StatusOK, antennas.GetAntennasResp{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		Data:       data,
	})
}
