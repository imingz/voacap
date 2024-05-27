// Code generated by hertz generator. DO NOT EDIT.

package station

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	station "voacap/biz/handler/station"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_stations := root.Group("/stations", _stationsMw()...)
		_stations.POST("/addStation", append(_addstationMw(), station.AddStation)...)
		_stations.GET("/getStations", append(_getstationsMw(), station.GetStations)...)
		_stations.POST("/updateStationById", append(_updatestationbyidMw(), station.UpdateStationById)...)
	}
}
