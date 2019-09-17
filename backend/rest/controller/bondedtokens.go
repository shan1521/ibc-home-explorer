package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/vo"
)

func RegisterBondedTokens(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerBondedTokensValidators,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func registerBondedTokensValidators(r *mux.Router) error {
	doApi(r, types.UrlRegisterBondedTokensValidators, "GET", func(request vo.IrisReq) interface{} {
		bondedtokens.SetTid(request.TraceId)
		vtype := QueryParam(request, "type")
		result, err := bondedtokens.QueryBondedTokensValidator(vtype)
		if err != nil {
			return vo.NewResponse("-1", err.Error(), nil)
		}
		return vo.NewResponse(types.CodeSuccess.Code, types.CodeSuccess.Msg, result)
	})

	return nil
}