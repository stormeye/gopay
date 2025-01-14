package wechat

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/stormeye/gopay"
	"github.com/stormeye/gopay/pkg/util"
)

// 申请退款API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_14.shtml
func (c *ClientV3) V3Refund(bm gopay.BodyMap) (wxRsp *RefundRsp, err error) {
	var (
		ts       = time.Now().Unix()
		nonceStr = util.GetRandomString(32)
		url      = v3DomesticRefund
	)
	authorization, err := c.authorization(MethodPost, url, nonceStr, ts, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &RefundRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(RefundOrderResponse)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询单笔退款API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_2_10.shtml
func (c *ClientV3) V3RefundQuery(outRefundNo string) (wxRsp *RefundQueryRsp, err error) {
	var (
		ts       = time.Now().Unix()
		nonceStr = util.GetRandomString(32)
		uri      = fmt.Sprintf(v3DomesticRefundQuery, outRefundNo)
	)
	authorization, err := c.authorization(MethodGet, uri, nonceStr, ts, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &RefundQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(RefundQueryResponse)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}
