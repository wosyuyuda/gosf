package sf

import "testing"

func TestF下单(t *testing.T) {
	client := F新建链接("QQBXX1pcUCKV", "HdOrk7Wu9uQiXnjXO6kaBBYV40emI7dN", "7551234567", 1)
	client.F下单(body)
}

var (
	body = `{
		"cargoDetails": [{
			"amount": 308.0,
			"count": 1.0,
			"name": "君宝牌地毯",
			"unit": "个",
			"volume": 0.0,
			"weight": 0.1
		}],
		"contactInfoList": [{
			"address": "十堰市丹江口市公园路155号",
			"city": "十堰市",
			"company": "清雅轩保健品专营店",
			"contact": "张三丰",
			"contactType": 1,
			"county": "武当山风景区",
			"mobile": "17006805888",
			"province": "湖北省"
		}, {
			"address": "湖北省襄阳市襄城区环城东路122号",
			"city": "襄阳市",
			"contact": "郭襄阳",
			"county": "襄城区",
			"contactType": 2,
			"mobile": "18963828829",
			"province": "湖北省"
		}],
		"customsInfo": {},
		"expressTypeId": 1,
		"isReturnRoutelabel": 1,
		"extraInfoList": [],
		"isOneselfPickup": 0,
		"language": "zh-CN",
		"monthlyCard": "7551234567",
		"orderId": "QIAO-2020069-005",
		"parcelQty": 1,
		"payMethod": 1,
		"totalWeight": 6
	}`
)
