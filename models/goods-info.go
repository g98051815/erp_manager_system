package models

import (

	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"erp_manager_system/structures"
	"log"
)


type GoodsInfo struct{

	structures.GoodsInfo

}

func init(){

	dsn := beego.AppConfig.String("dsn")

	orm.RegisterDriver("mysql",orm.DRMySQL)

	orm.RegisterDataBase("default","mysql",dsn)

	orm.RegisterModel( new( GoodsInfo ) )

}


func (g *GoodsInfo) Input() error{

	orm := orm.NewOrm()

	_ , err := orm.Insert(g)

	if err!= nil{

		return err
	}

	return nil
}


func (g *GoodsInfo)List( limit int , page int )(num int64 , goodsInfo []GoodsInfo , err error){

	result := []GoodsInfo{}

	orm := orm.NewOrm()

	if num , err := orm.Raw("select * from "+g.TableName()).QueryRows( &result ); err!=nil{

		log.Println(err)

		return num , result , err

	}else{

		return num , result , nil

	}



}