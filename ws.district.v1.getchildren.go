package qq

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"net/http"
)

type WsDistrictV1GetchildrenResponse struct {
	Status      int    `json:"status"`       // 状态码
	Message     string `json:"message"`      // 状态说明
	DataVersion string `json:"data_version"` // 行政区划数据版本，便于您判断更新
	Result      [][]struct {
		Id       string   `json:"id"`       // 行政区划唯一标识（adcode）
		Name     string   `json:"name"`     // 简称，如“内蒙古”
		Fullname string   `json:"fullname"` // 全称，如“内蒙古自治区”
		Pinyin   []string `json:"pinyin"`   // 行政区划拼音，每一下标为一个字的全拼
		Location struct {
			Lat float64 `json:"lat"` // 纬度
			Lng float64 `json:"lng"` // 经度
		} `json:"location"`
	} `json:"result"`
}

type WsDistrictV1GetchildrenResult struct {
	Result WsDistrictV1GetchildrenResponse // 结果
	Body   []byte                          // 内容
	Http   gorequest.Response              // 请求
}

func newWsDistrictV1GetchildrenResult(result WsDistrictV1GetchildrenResponse, body []byte, http gorequest.Response) *WsDistrictV1GetchildrenResult {
	return &WsDistrictV1GetchildrenResult{Result: result, Body: body, Http: http}
}

// WsDistrictV1Getchildren 获取下级行政区划
// https://lbs.qq.com/service/webService/webServiceGuide/webServiceDistrict#3
func (c *Client) WsDistrictV1Getchildren(ctx context.Context, id string, notMustParams ...gorequest.Params) (*WsDistrictV1GetchildrenResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("key", c.key)
	params.Set("id", id)
	params.Set("output", "JSON")
	// 请求
	request, err := c.request(ctx, apiUrl+"/ws/district/v1/getchildren", params, http.MethodGet)
	if err != nil {
		return newWsDistrictV1GetchildrenResult(WsDistrictV1GetchildrenResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WsDistrictV1GetchildrenResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWsDistrictV1GetchildrenResult(response, request.ResponseBody, request), err
}
