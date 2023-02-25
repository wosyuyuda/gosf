package sf

import (
	"encoding/json"
	"errors"

	"github.com/wosyuyuda/gosf/util"
)

type SearchRoute struct {
	TrackingNumber []string `json:"trackingNumber"` //订单id
	TrackingType   int      `json:"trackingType"`   //1:根据顺丰运单号查询,2:根据客户订单号查询
}

//下单前有些参数要检验一下最好。。。
func (a *Client) F下单(body *SFOrder) (res *OrderResponse, err error) {
	body.MonthlyCard = a.V月结账号
	body.Language = "zh-CN"
	if body.OrderId == "" {
		body.OrderId = util.Idw.F获取字符串id()
	}
	if body.PayMethod == 0 {
		body.PayMethod = 1 //默认寄方付
	}
	if body.ExpressTypeId == 0 {
		//https://open.sf-express.com/developSupport/734349?activeIndex=324604
		body.ExpressTypeId = 2 //默认2为标快
	}

	bd, err := a.F发送请求(V下单服务, body)
	if err != nil {
		return
	}
	res = new(OrderResponse)
	if err = json.Unmarshal([]byte(bd), res); err != nil {
		return
	}
	if !res.Success {
		err = errors.New(res.ErrorMsg)
	}
	return
}

//不传为顺丰的单号，传了值则为商家单号
func (a *Client) F快递查询(id []string, 顺丰单号 ...int) (res *OrderResponse, err error) {
	search := new(SearchRoute)
	search.TrackingNumber = id
	//判断id是顺丰的单号还是自带的单号
	if len(顺丰单号) == 0 {
		search.TrackingType = 1
	} else {
		search.TrackingType = 2
	}
	bd, err := a.F发送请求(V路由查询接口, search)
	if err != nil {
		return
	}
	res = new(OrderResponse)
	if err = json.Unmarshal([]byte(bd), res); err != nil {
		return
	}
	if !res.Success {
		err = errors.New(res.ErrorMsg)
	}
	//fmt.Println("结果：", string(bd))
	return
}
