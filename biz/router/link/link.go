// Code generated by hertz generator. DO NOT EDIT.

package link

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	link "voacap/biz/handler/link"
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
		_links := root.Group("/links", _linksMw()...)
		_links.POST("/addLink", append(_addlinkMw(), link.AddLink)...)
		_links.POST("/deleteLinkById", append(_deletelinkbyidMw(), link.DeleteLinkById)...)
		_links.GET("/getLinks", append(_getlinksMw(), link.GetLinks)...)
		_links.POST("/updateLinkById", append(_updatelinkbyidMw(), link.UpdateLinkById)...)
	}
}
