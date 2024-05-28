package service

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"voacap/biz/dal/db"
	"voacap/biz/model/common"
	"voacap/biz/model/link"
	"voacap/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

type LinkService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewLinkService create LinkService service
func NewLinkService(ctx context.Context, c *app.RequestContext) *LinkService {
	return &LinkService{ctx: ctx, c: c}
}

// GetAntennas 获取所有天线信息
func (s *LinkService) GetLinks() ([]*common.Link, error) {
	links, err := db.QueryAllLink()
	if err != nil {
		return nil, err
	}

	result := make([]*common.Link, len(links))
	for i, link := range links {
		result[i] = &common.Link{
			CircuitReliability: link.CircuitReliability,
			Coefficient:        link.Coefficient,
			Date:               link.Date.Format(time.DateOnly),
			IxAntennaFbandMax:  link.IxAntennaFbandMax,
			IxAntennaFbandMin:  link.IxAntennaFbandMin,
			IxAntennaFile:      link.IxAntennaFile,
			IxAntennaID:        link.IxAntennaID,
			IxAntennaName:      link.IxAntennaName,
			IxPower:            link.IxPower,
			IxStationID:        link.IxStationID,
			IxStationLat:       link.IxStationLat,
			IxStationLng:       link.IxStationLng,
			IxStationName:      link.IxStationName,
			LinkID:             link.LinkID,
			LinkType:           link.LinkType,
			Noise:              link.Noise,
			RxAntennaFbandMax:  link.RxAntennaFbandMax,
			RxAntennaFbandMin:  link.RxAntennaFbandMin,
			RxAntennaFile:      link.RxAntennaFile,
			RxAntennaID:        link.RxAntennaID,
			RxAntennaName:      link.RxAntennaName,
			RxStationID:        link.RxStationID,
			RxStationLat:       link.RxStationLat,
			RxStationLng:       link.RxStationLng,
			RxStationName:      link.RxStationName,
			SNR:                link.SNR,
			SunspotNum:         link.SunspotNum,
			TimeType:           link.TimeType,
			TxAntennaFbandMax:  link.TxAntennaFbandMax,
			TxAntennaFbandMin:  link.TxAntennaFbandMin,
			TxAntennaFile:      link.TxAntennaFile,
			TxAntennaID:        link.TxAntennaID,
			TxAntennaName:      link.TxAntennaName,
			TxPower:            link.TxPower,
			TxStationID:        link.TxStationID,
			TxStationLat:       link.TxStationLat,
			TxStationLng:       link.TxStationLng,
			TxStationName:      link.TxStationName,
		}
	}

	return result, nil
}

// AddLink 添加链路信息
func (s *LinkService) AddLink(req *link.AddLinkRequest) error {
	date, err := time.Parse(time.DateOnly, req.SysConfig.Date)
	if err != nil {
		return err
	}

	var linkType string
	if req.InterferStation.TransStation.StationID != -1 {
		linkType = "干扰"
	} else {
		linkType = "通信"
	}

	l := &db.Link{
		LinkType: linkType,

		// SysConfig
		Date:               date,
		TimeType:           req.SysConfig.TimeType,
		Coefficient:        req.SysConfig.Coefficient,
		SunspotNum:         req.SysConfig.SunspotNum,
		CircuitReliability: req.SysConfig.CircuitReliability,
		SNR:                req.SysConfig.SNR,
		Noise:              req.SysConfig.Noise,

		// TransStation
		TxStationID:   req.TransStation.TransStation.StationID,
		TxStationLat:  req.TransStation.TransStation.Slatitude,
		TxStationLng:  req.TransStation.TransStation.Slongitude,
		TxStationName: req.TransStation.TransStation.Sname,

		TxAntennaID:       req.TransStation.TransAntenna.AntennaID,
		TxAntennaName:     req.TransStation.TransAntenna.Aname,
		TxAntennaFile:     req.TransStation.TransAntenna.Afile,
		TxAntennaFbandMin: req.TransStation.TransAntenna.AfbandMin,
		TxAntennaFbandMax: req.TransStation.TransAntenna.AfbandMax,

		TxPower: req.TransStation.TransPower,

		// RecvStation
		RxStationID:   req.RecvStation.RecvStation.StationID,
		RxStationLat:  req.RecvStation.RecvStation.Slatitude,
		RxStationLng:  req.RecvStation.RecvStation.Slongitude,
		RxStationName: req.RecvStation.RecvStation.Sname,

		RxAntennaID:       req.RecvStation.RecvAntenna.AntennaID,
		RxAntennaName:     req.RecvStation.RecvAntenna.Aname,
		RxAntennaFile:     req.RecvStation.RecvAntenna.Afile,
		RxAntennaFbandMin: req.RecvStation.RecvAntenna.AfbandMin,
		RxAntennaFbandMax: req.RecvStation.RecvAntenna.AfbandMax,

		// InterferStation
		IxStationID:   req.InterferStation.TransStation.StationID,
		IxStationLat:  req.InterferStation.TransStation.Slatitude,
		IxStationLng:  req.InterferStation.TransStation.Slongitude,
		IxStationName: req.InterferStation.TransStation.Sname,

		IxAntennaID:       req.InterferStation.TransAntenna.AntennaID,
		IxAntennaName:     req.InterferStation.TransAntenna.Aname,
		IxAntennaFile:     req.InterferStation.TransAntenna.Afile,
		IxAntennaFbandMin: req.InterferStation.TransAntenna.AfbandMin,
		IxAntennaFbandMax: req.InterferStation.TransAntenna.AfbandMax,

		IxPower: req.InterferStation.TransPower,
	}

	return db.CreateLink(l)
}

// UpdateLinkById 更新链路信息
func (s *LinkService) UpdateLinkById(req *link.UpdateLinkByIdRequest) error {
	date, err := time.Parse(time.DateOnly, req.Date)
	if err != nil {
		return err
	}
	l := &db.Link{
		LinkID:             req.LinkID,
		LinkType:           req.LinkType,
		Date:               date,
		TimeType:           req.TimeType,
		Coefficient:        req.Coefficient,
		SunspotNum:         req.SunspotNum,
		CircuitReliability: req.CircuitReliability,
		SNR:                req.SNR,
		Noise:              req.Noise,
		TxStationID:        req.TxStationID,
		TxStationName:      req.TxStationName,
		TxStationLat:       req.TxStationLat,
		TxStationLng:       req.TxStationLng,
		TxAntennaID:        req.TxAntennaID,
		TxAntennaName:      req.TxAntennaName,
		TxAntennaFile:      req.TxAntennaFile,
		TxAntennaFbandMin:  req.TxAntennaFbandMin,
		TxAntennaFbandMax:  req.TxAntennaFbandMax,
		TxPower:            req.TxPower,
		RxStationID:        req.RxStationID,
		RxStationName:      req.RxStationName,
		RxStationLat:       req.RxStationLat,
		RxStationLng:       req.RxStationLng,
		RxAntennaID:        req.RxAntennaID,
		RxAntennaName:      req.RxAntennaName,
		RxAntennaFile:      req.RxAntennaFile,
		RxAntennaFbandMin:  req.RxAntennaFbandMin,
		RxAntennaFbandMax:  req.RxAntennaFbandMax,
		IxStationID:        req.IxStationID,
		IxStationName:      req.IxStationName,
		IxStationLat:       req.IxStationLat,
		IxStationLng:       req.IxStationLng,
		IxAntennaID:        req.IxAntennaID,
		IxAntennaName:      req.IxAntennaName,
		IxAntennaFile:      req.IxAntennaFile,
		IxAntennaFbandMin:  req.IxAntennaFbandMin,
		IxAntennaFbandMax:  req.IxAntennaFbandMax,
		IxPower:            req.IxPower,
	}
	return db.UpdateLinkById(l)
}

// DeleteLinkById 删除链路信息
func (s *LinkService) DeleteLinkById(req *link.DeleteLinkByIdRequest) error {
	return db.DeleteLinkById(req.LinkID)
}

// WriteLink2File 将链路信息写入文件
func (s *LinkService) WriteLink2File(req *link.WriteLink2FileRequest, trans bool) error {
	date, err := time.Parse(time.DateOnly, req.Date)
	if err != nil {
		return err
	}
	lines := []string{
		"COMMENT    Any VOACAP default cards may be placed in the file: VOACAP.DEF",
		"LINEMAX      55       number of lines-per-page",
		fmt.Sprintf("COEFFS    %s", req.Coefficient),
		"TIME          1   24    1    1",
		fmt.Sprintf("MONTH      %d %d.%d", date.Year(), int(date.Month()), date.Day()),
		fmt.Sprintf("SUNSPOT    %d.", req.SunspotNum),
	}

	if trans {
		lines = append(lines,
			fmt.Sprintf("LABEL     %s, %s", req.TxStationName, req.RxStationName),
			fmt.Sprintf("CIRCUIT%8.2fN %8.2fE %8.2fN %8.2fE  S     0", req.TxStationLat, req.TxStationLng, req.RxStationLat, req.RxStationLng),
		)
	} else {
		lines = append(lines,
			fmt.Sprintf("LABEL     %s, %s", req.IxStationName, req.RxStationName),
			fmt.Sprintf("CIRCUIT%8.2fN %8.2fE %8.2fN %8.2fE  S     0", req.IxStationLat, req.IxStationLng, req.RxStationLat, req.RxStationLng),
		)
	}

	lines = append(lines,
		fmt.Sprintf("SYSTEM       1. %v. 0.10  %d. %.1f 3.00 0.10", req.Noise, int(req.CircuitReliability*100), req.SNR),
		"FPROB      1.00 1.00 1.00 0.70",
	)

	if trans {
		lines = append(lines,
			fmt.Sprintf("ANTENNA       1    1    2   30     0.000[samples\\%-13s]106.3    %.4f", req.TxAntennaFile, req.TxPower),
		)
	} else {
		lines = append(lines,
			fmt.Sprintf("ANTENNA       1    1    2   30     0.000[samples\\%-13s]106.3    %.4f", req.IxAntennaFile, req.IxPower),
		)
	}

	lines = append(lines,
		fmt.Sprintf("ANTENNA       2    2    2   30     0.000[samples\\%-13s]289.1    0.0000", req.RxAntennaFile),
		"FREQUENCY  2.00 3.00 4.00 5.00 6.00 7.00 8.00 9.0010.00 0.00 0.00",
		"METHOD       26    0",
		"EXECUTE",
		"METHOD       30    0",
		"EXECUTE",
		"FREQUENCY 11.0012.0013.0014.0015.0016.0017.0018.0019.0020.00 0.00",
		"EXECUTE",
		"FREQUENCY 21.0022.0023.0024.0025.0026.0027.0028.0029.0030.00 0.00",
		"EXECUTE",
		"QUIT",
		"",
	)

	file, err := os.Create(utils.GetFilePath("C:/MyVoacap/myVOACAP/run/voacapx.dat"))
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()

	return err
}

// 调用计算exe进行链路计算
func (s *LinkService) CalculateLink() error {
	cmd := exec.Command(utils.GetFilePath("C:\\MyVoacap\\myVOACAP\\CheckMate\\Win32\\MyVoacap.exe"))
	return cmd.Run()
}

// GetLinkResult 获取链路计算结果
func (s *LinkService) GetLinkResult() ([]*link.CalculateResult, error) {
	filePath := utils.GetFilePath("C:/MyVoacap/myVOACAP/run/voacapx.out")
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	// 解析 LMT MUF LUF
	lines1 := lines[41:65]
	var transLink = make([]*link.CalculateResult, 0)
	for _, line := range lines1 {
		columns := strings.Fields(line)
		if len(columns) >= 7 {
			GMT, _ := strconv.ParseFloat(columns[0], 64)
			LMT, _ := strconv.ParseFloat(columns[1], 64)
			FOT, _ := strconv.ParseFloat(columns[2], 64)
			MUF, _ := strconv.ParseFloat(columns[5], 64)
			LUF, _ := strconv.ParseFloat(columns[6], 64)
			transLink = append(transLink, &link.CalculateResult{
				GMT:  GMT,
				LMT:  LMT,
				MUF:  MUF,
				LUF:  LUF,
				DBU:  []float64{0.0, 0.0},
				SDBW: []float64{0.0, 0.0},
				SNR:  []float64{0.0, 0.0},
				FOT:  FOT,
			})
		}
	}

	row := 66
	for i := 0; i < 36; i++ {
		lines2 := lines[row : row+54]

		parseLine := func(line string, endIndex int) []float64 {
			if i < 12 {
				endIndex--
			}
			fields := strings.Fields(line)
			fields = fields[1 : len(fields)+endIndex]
			values := make([]float64, len(fields))
			for i, field := range fields {
				values[i], _ = strconv.ParseFloat(field, 64)
			}
			return values
		}

		getParsedValues := func(offset int) (float64, []float64, []float64, []float64) {
			targetGMT, _ := strconv.ParseFloat(strings.Fields(lines2[offset])[0], 64)
			DBU := parseLine(lines2[offset+7], -2)
			SDBW := parseLine(lines2[offset+8], -3)
			SNR := parseLine(lines2[offset+10], -2)
			return targetGMT, DBU, SDBW, SNR
		}

		valuesArray := [2](func() (float64, []float64, []float64, []float64)){
			func() (float64, []float64, []float64, []float64) { return getParsedValues(9) },
			func() (float64, []float64, []float64, []float64) { return getParsedValues(32) },
		}

		for i := range transLink {
			for _, getValues := range valuesArray {
				targetGMT, DBU, SDBW, SNR := getValues()
				if transLink[i].GMT == targetGMT {
					transLink[i].DBU = append(transLink[i].DBU, DBU...)
					transLink[i].SDBW = append(transLink[i].SDBW, SDBW...)
					transLink[i].SNR = append(transLink[i].SNR, SNR...)
				}
			}
		}

		row += 55
	}

	return transLink, err
}
