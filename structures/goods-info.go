package structures

type GoodsInfo struct{
	//唯一标识
	PostID string `json:"post_id" orm:"column(post_id); pk;" form:"post_id"`
	//商品的标题
	GoodsTitle string `json:"title" orm:"column(title)" form:"title"`
	//商品介绍
	GoodsDescription	string `json:"g_description" orm:"column(g_description)" form:"g_description"`
	//商品的副标题
	GoodsSmallTitle string `json:"small_title" orm:"column(small_title)" form:"small_title"`
	//商品的备注
	GoodsNote	string `json:"note" orm:"column(note)" form:"note"`
}

func(g *GoodsInfo) TableName() string{

	return "goods_info"

}