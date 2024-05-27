// Code generated by hertz generator.

package antennas

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	antennas "voacap/biz/model/antennas"
)

// GetAntennas .
// @router /getAntennas [GET]
func GetAntennas(ctx context.Context, c *app.RequestContext) {
	var err error
	var req antennas.GetAntennasReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(antennas.GetAntennasResp)

	c.JSON(consts.StatusOK, resp)
}
