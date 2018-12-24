package controllers

import (
	"github.com/astaxie/beego"
	"erp_manager_system/models"
	"log"
	"encoding/json"
	"fmt"
)

type GoodsInfoController struct{

	beego.Controller

}

//@Title Create
//@Description 添加商品的信息
//@Success 200 {string}
//@Failure 400 bad request
//@router / [post]
func (g *GoodsInfoController) Post(){

	context :=g.Ctx.Input.RequestBody

	parseForm :=models.GoodsInfo{}

	response := map[string]interface{}{}

	message:= map[string]interface{}{

		"code":"200",

		"message":"request success",

		"response":response,

	}

	if err := json.Unmarshal(context,&parseForm); err!= nil{

		log.Println("error display:",err)
		message["code"]=400
		message["message"] = err

	}

	if err :=g.ParseForm( &parseForm ); err!=nil{

		log.Println(err)
		message["code"]=400
		message["message"] = err

	}

	if err :=parseForm.Input(); err!=nil{
		log.Println(err)
		message["code"]=400
		message["message"] = err


	}

	fmt.Println(message)
	g.Data["json"] = message

	g.ServeJSON()
}

//@Title Get
//@Description 获取商品的信息
//@Success 200 {string}
//@Failure 400 bad request
//@router /:postid [get]
func (g *GoodsInfoController) Get(){

	g.Data["json"] = map[string]string{
		"name":"获取商品的信息",
	}

	g.ServeJSON()

}

//@Title GetAll
//@Description 查询商品列表
//Success 200 {string}
//@router / [get]
func (g *GoodsInfoController) GetAll(){

	var goodsInfoList []models.GoodsInfo

	result := models.GoodsInfo{}

	num , goodsInfoList , err :=result.List(1,1)

	message := map[string]interface{}{

		"code":200,
		"message":"request success",
		"response":map[string]interface{}{
			"counter":num,
			"list":goodsInfoList,
		},
	}

	if err!=nil{

		message["code"] = 400
		message["message"] = "request success"

	}
	g.Data["json"] = message
	g.ServeJSON()

}