package ecs

import "sync"

// EndpointMap Endpoint Data
var once sync.Once
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	once.Do(func() {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "ecs-cn-hangzhou.aliyuncs.com",
			"cn-beijing-gov-1":            "ecs.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "ecs-cn-hangzhou.aliyuncs.com",
			"cn-shanghai-inner":           "ecs.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "ecs-cn-hangzhou.aliyuncs.com",
			"cn-haidian-cm12-c01":         "ecs-cn-hangzhou.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "ecs-cn-hangzhou.aliyuncs.com",
			"cn-north-2-gov-1":            "ecs.aliyuncs.com",
			"cn-yushanfang":               "ecs.aliyuncs.com",
			"cn-hongkong-finance-pop":     "ecs.aliyuncs.com",
			"cn-shanghai-finance-1":       "ecs-cn-hangzhou.aliyuncs.com",
			"cn-beijing-finance-pop":      "ecs.aliyuncs.com",
			"cn-wuhan":                    "ecs.aliyuncs.com",
			"cn-zhangbei":                 "ecs.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "ecs.cn-qingdao-nebula.aliyuncs.com",
			"rus-west-1-pop":              "ecs.aliyuncs.com",
			"cn-shanghai-et15-b01":        "ecs-cn-hangzhou.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "ecs-cn-hangzhou.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "ecs-cn-hangzhou.aliyuncs.com",
			"eu-west-1-oxs":               "ecs.cn-shenzhen-cloudstone.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "ecs-cn-hangzhou.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "ecs-cn-hangzhou.aliyuncs.com",
			"cn-shenzhen-finance-1":       "ecs-cn-hangzhou.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "ecs-cn-hangzhou.aliyuncs.com",
			"cn-hangzhou-test-306":        "ecs-cn-hangzhou.aliyuncs.com",
			"cn-huhehaote-nebula-1":       "ecs.cn-qingdao-nebula.aliyuncs.com",
			"cn-shanghai-et2-b01":         "ecs-cn-hangzhou.aliyuncs.com",
			"cn-hangzhou-finance":         "ecs.aliyuncs.com",
			"cn-beijing-nu16-b01":         "ecs-cn-hangzhou.aliyuncs.com",
			"cn-edge-1":                   "ecs.cn-qingdao-nebula.aliyuncs.com",
			"cn-fujian":                   "ecs-cn-hangzhou.aliyuncs.com",
			"ap-northeast-2-pop":          "ecs.aliyuncs.com",
			"cn-shenzhen-inner":           "ecs.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "ecs.cn-zhangjiakou.aliyuncs.com",
			"cn-hangzhou":                 "ecs-cn-hangzhou.aliyuncs.com",
		}
	})
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
