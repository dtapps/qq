package qq

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"net/http"
)

type WsDistrictV1ListResponse struct {
	Status      int    `json:"status"`       // 状态码
	Message     string `json:"message"`      // 状态说明
	DataVersion string `json:"data_version"` // 行政区划数据版本，便于您判断更新
	Result      [][]struct {
		Id       string   `json:"id"`               // 行政区划唯一标识（adcode）
		Name     string   `json:"name,omitempty"`   // 简称，如“内蒙古”
		Fullname string   `json:"fullname"`         // 全称，如“内蒙古自治区”
		Pinyin   []string `json:"pinyin,omitempty"` // 行政区划拼音，每一下标为一个字的全拼
		Location struct {
			Lat float64 `json:"lat"` // 纬度
			Lng float64 `json:"lng"` // 经度
		} `json:"location"` // 经纬度
		Cidx []int `json:"cidx,omitempty"`
	} `json:"result"`
}

type WsDistrictV1ListResult struct {
	Result WsDistrictV1ListResponse // 结果
	Body   []byte                   // 内容
	Http   gorequest.Response       // 请求
}

func newWsDistrictV1ListResult(result WsDistrictV1ListResponse, body []byte, http gorequest.Response) *WsDistrictV1ListResult {
	return &WsDistrictV1ListResult{Result: result, Body: body, Http: http}
}

// WsDistrictV1List 获取省市区列表
// https://lbs.qq.com/service/webService/webServiceGuide/webServiceDistrict#2
func (c *Client) WsDistrictV1List(ctx context.Context, notMustParams ...gorequest.Params) (*WsDistrictV1ListResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("key", c.key)
	params.Set("output", "JSON")
	// 请求
	request, err := c.request(ctx, apiUrl+"/ws/district/v1/list", params, http.MethodGet)
	if err != nil {
		return newWsDistrictV1ListResult(WsDistrictV1ListResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WsDistrictV1ListResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWsDistrictV1ListResult(response, request.ResponseBody, request), err
}
