package sf

import (
	"encoding/json"
	"errors"

	"github.com/wosyuyuda/gosf/util"
)

//下单前有些参数要检验一下最好。。。
func (a *Client) F下单(body *SFOrder) (res *OrderResponse, err error) {
	body.MonthlyCard = a.V月结账号
	body.Language = "zh-CN"
	body.OrderId = "SF" + util.Idw.F获取字符串id()
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
