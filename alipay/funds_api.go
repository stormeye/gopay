package alipay

import (
	"encoding/json"
	"fmt"

	"github.com/stormeye/gopay"
)

// alipay.fund.trans.uni.transfer(单笔转账接口)
//	文档地址：https://opendocs.alipay.com/apis/api_28/alipay.fund.trans.uni.transfer
func (a *Client) FundTransUniTransfer(bm gopay.BodyMap) (aliRsp *FundTransUniTransferResponse, err error) {
	err = bm.CheckEmptyError("out_biz_no", "trans_amount", "product_code", "payee_info")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.fund.trans.uni.transfer"); err != nil {
		return nil, err
	}
	aliRsp = new(FundTransUniTransferResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.fund.account.query(支付宝资金账户资产查询接口)
//	文档地址：https://opendocs.alipay.com/apis/api_28/alipay.fund.account.query
func (a *Client) FundAccountQuery(bm gopay.BodyMap) (aliRsp *FundAccountQueryResponse, err error) {
	err = bm.CheckEmptyError("alipay_user_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.fund.account.query"); err != nil {
		return nil, err
	}
	aliRsp = new(FundAccountQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.fund.trans.common.query(转账业务单据查询接口)
//	文档地址：https://opendocs.alipay.com/apis/api_28/alipay.fund.trans.common.query
func (a *Client) FundTransCommonQuery(bm gopay.BodyMap) (aliRsp *FundTransCommonQueryResponse, err error) {
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.fund.trans.common.query"); err != nil {
		return nil, err
	}
	aliRsp = new(FundTransCommonQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}
