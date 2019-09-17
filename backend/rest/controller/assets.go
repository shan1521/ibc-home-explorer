package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/vo"
	"github.com/irisnet/explorer/backend/orm/document"
	"strings"
)

var assets service.AssetsService

func RegisterAssets(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryNativeAsset,
		registerQueryGatewayAsset,
		registerAssetTokens,
		registerAssetGatewayInfo,
		registerAssetTokenInfo,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func registerQueryNativeAsset(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryNativeAsset, "GET", func(request vo.IrisReq) interface{} {
		assets.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 5))
		total := QueryParam(request, "total")
		txtype := QueryParam(request, "tx_type")
		symbol := QueryParam(request, "symbol")
		istotal := true
		if total == "false" {
			istotal = false
		}
		res, err := assets.GetNativeAsset(symbol, txtype, page, size, istotal)
		if err != nil {
			return vo.NewResponse("-1", err.Error(), nil)
		}
		return vo.NewResponse(types.CodeSuccess.Code, types.CodeSuccess.Msg, res)
	})
	return nil
}

func registerQueryGatewayAsset(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryGatewayAsset, "GET", func(request vo.IrisReq) interface{} {
		assets.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 5))
		txtype := QueryParam(request, "tx_type")
		symbol := QueryParam(request, "symbol")
		gateway := QueryParam(request, "gateway")
		total := QueryParam(request, "total")
		istotal := true
		if total == "false" {
			istotal = false
		}
		var res vo.AssetsRespond
		var err error
		if txtype == document.Tx_Asset_TxType_TransferGatewayOwner {
			moniker := QueryParam(request, "moniker")
			res, err = assets.GetTransferGatewayOwner(moniker, page, size, istotal)
		} else {
			res, err = assets.GetGatewayAsset(symbol, gateway, txtype, page, size, istotal)
		}
		if err != nil {
			return vo.NewResponse("-1", err.Error(), nil)
		}
		return vo.NewResponse(types.CodeSuccess.Code, types.CodeSuccess.Msg, res)
	})
	return nil
}

func registerAssetTokens(r *mux.Router) error {
	doApi(r, types.UrlRegisterAssetTokens, "GET", func(request vo.IrisReq) interface{} {
		assets.SetTid(request.TraceId)
		//tokenid := QueryParam(request, "tokenid")
		source := QueryParam(request, "source")
		result, err := assets.QueryAssetTokens(source)
		if err != nil {
			return vo.NewResponse("-1", err.Error(), nil)
		}
		return vo.NewResponse(types.CodeSuccess.Code, types.CodeSuccess.Msg, result)
	})
	return nil
}

func registerAssetTokenInfo(r *mux.Router) error {
	doApi(r, types.UrlRegisterAssetTokenInfo, "GET", func(request vo.IrisReq) interface{} {
		assets.SetTid(request.TraceId)
		tokenid := Var(request, "token_id")
		result, err := assets.QueryAssetTokenDetail(strings.ToLower(tokenid))
		if err != nil {
			return vo.NewResponse("-1", err.Error(), nil)
		}
		return vo.NewResponse(types.CodeSuccess.Code, types.CodeSuccess.Msg, result)
	})
	return nil
}

func registerAssetGatewayInfo(r *mux.Router) error {
	doApi(r, types.UrlRegisterAssetGatewayInfo, "GET", func(request vo.IrisReq) interface{} {
		assets.SetTid(request.TraceId)
		moniker := Var(request, "moniker")
		result, err := assets.QueryAssetGatewayDetail(moniker)
		if err != nil {
			return vo.NewResponse("-1", err.Error(), nil)
		}
		return vo.NewResponse(types.CodeSuccess.Code, types.CodeSuccess.Msg, result)
	})
	return nil
}
