package controllers

import (
	"fmt"
	"os"
	"tl_mlkit/models"
	"tl_mlkit/services"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/manucorporat/try"
)

// MlkitController operations for Mlkit
type MlkitController struct {
	beego.Controller
}

// URLMapping ...
func (c *MlkitController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Mlkit
// @Param	body		body 	models.Mlkit	true		"body for Mlkit content"
// @Success 201 {object} models.Mlkit
// @Failure 403 body is empty
// @router / [post]
func (c *MlkitController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Mlkit by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Mlkit
// @Failure 403 :id is empty
// @router /:id [get]
func (c *MlkitController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Mlkit
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Mlkit
// @Failure 403
// @router / [get]
func (c *MlkitController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Mlkit
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Mlkit	true		"body for Mlkit content"
// @Success 200 {object} models.Mlkit
// @Failure 403 :id is not int
// @router /:id [put]
func (c *MlkitController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Mlkit
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *MlkitController) Delete() {

}

// GetFiles ...
// @Title GetFiles
// @Description get Mlkit by file
// @Param	file		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Mlkit
// @Failure 403 :file is empty
// @router /:file [post]
func (c *MlkitController) GetFiles() {
	fmt.Print("******file read start ******************")
	file, header, er := c.GetFile("image")
	if file != nil {

		try.This(func() {
			mlkit, err := services.DetectText(os.Stdout, file)
			if mlkit != nil {
				c.Data["json"] = mlkit
				c.ServeJSON()
			} else {
				fmt.Println(err)
				c.Data["json"] = models.GetCode("image_not_valid")
				c.ServeJSON()
			}

		}).Finally(func() {

		}).Catch(func(e try.E) {
			fmt.Println(e)
			c.Data["json"] = models.GetCode("image_not_valid")
			c.ServeJSON()
		})

		fmt.Print("header", header.Filename)
	} else {
		fmt.Print("image error field missing")
		fmt.Print(er)
		c.Data["json"] = models.GetCode("field_missing")
		c.ServeJSON()
	}
}
