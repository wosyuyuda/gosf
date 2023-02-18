package sf

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/wosyuyuda/gosf/model"
)

var (
	SF *Client
)

func init() {
	SF = F新建链接("QQBXX1pcUCKV", "HdOrk7Wu9uQiXnjXO6kaBBYV40emI7dN", "7551234567", V沙箱环境)

}

func TestF下单(t *testing.T) {
	order := model.SFOrder{
		IsReturnRoutelabel: 1,
		CargoDetails: []model.CargoDetail{
			{Name: "苹果手机"},
		},
		ContactInfoList: []model.ContactInfo{
			{ContactType: V寄件人, Address: "十堰市丹江口市公园路155号", City: "十堰市", Contact: "张三丰",
				County: "武当山风景区", Mobile: "17888888888", Province: "湖北省"},
			{ContactType: V收件人, Address: "湖北省襄阳市襄城区环城东路122号", City: "襄阳市", Contact: "郭襄阳",
				County: "襄城区", Mobile: "18963828829", Province: "湖北省"},
		},
	}

	//json.Unmarshal([]byte(body), &order)
	res, err := SF.F下单(&order)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	fmt.Println("res:", res)
}

var (
	db = `{"apiErrorMsg":"","apiResponseID":"000186644AFEBA3F9365AD375264A03F","apiResultCode":"A1000","apiResultData":"{\"success\":true,\"errorCode\":\"S0000\",\"errorMsg\":null,\"msgData\":{\"orderId\":\"SF192961670303744\",\"originCode\":\"719\",\"destCode\":\"710\",\"filterResult\":2,\"remark\":\"\",\"url\":null,\"paymentLink\":null,\"isUpstairs\":null,\"isSpecialWarehouseService\":null,\"mappingMark\":null,\"agentMailno\":null,\"returnExtraInfoList\":null,\"waybillNoInfoList\":[{\"waybillType\":1,\"waybillNo\":\"SF7444463102821\",\"boxNo\":null,\"length\":null,\"width\":null,\"height\":null,\"weight\":null,\"volume\":null}],\"routeLabelInfo\":[{\"code\":\"1000\",\"routeLabelData\":{\"waybillNo\":\"SF7444463102821\",\"sourceTransferCode\":\"719\",\"sourceCityCode\":\"719\",\"sourceDeptCode\":\"719\",\"sourceTeamCode\":\"\",\"destCityCode\":\"710\",\"destDeptCode\":\"710BF\",\"destDeptCodeMapping\":\"\",\"destTeamCode\":\"001\",\"destTeamCodeMapping\":\"\",\"destTransferCode\":\"710\",\"destRouteLabel\":\"710BF-011\",\"proName\":\"\",\"cargoTypeCode\":\"T6\",\"limitTypeCode\":\"T6\",\"expressTypeCode\":\"B1\",\"codingMapping\":\"S47\",\"codingMappingOut\":\"\",\"xbFlag\":\"0\",\"printFlag\":\"000000000\",\"twoDimensionCode\":\"MMM={'k1':'710','k2':'710BF','k3':'001','k4':'T4','k5':'SF7444463102821','k6':'A','k7':'dd297476'}\",\"proCode\":\"T  标快\",\"printIcon\":\"00010000\",\"abFlag\":\"A\",\"destPortCode\":\"\",\"destCountry\":\"\",\"destPostCode\":\"\",\"goodsValueTotal\":\"\",\"currencySymbol\":\"\",\"cusBatch\":\"\",\"goodsNumber\":\"\",\"errMsg\":\"\",\"checkCode\":\"dd297476\",\"proIcon\":\"\",\"fileIcon\":\"\",\"fbaIcon\":\"\",\"icsmIcon\":\"\",\"destGisDeptCode\":\"710BF\",\"newIcon\":null,\"sendAreaCode\":null,\"destinationStationCode\":null,\"sxLabelDestCode\":null,\"sxDestTransferCode\":null,\"sxCompany\":null,\"newAbFlag\":null,\"destAddrKeyWord\":\"\",\"rongType\":null,\"waybillIconList\":null},\"message\":\"SF7444463102821:\"}],\"contactInfoList\":null,\"sendStartTm\":null,\"customerRights\":null,\"expressTypeId\":null}}"}`
)

func TestF解析(t *testing.T) {
	res := new(model.Response)
	if err := json.Unmarshal([]byte(db), res); err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	fmt.Println("数据是:", res.ApiResultData)
	rdata := new(model.OrderResponse)
	if err := json.Unmarshal([]byte(res.ApiResultData), rdata); err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	fmt.Printf("数据是:%+v", rdata)
}
