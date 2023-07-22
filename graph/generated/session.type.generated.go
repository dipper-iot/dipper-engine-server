// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dipper-iot/dipper-engine-server/ent"
	"github.com/dipper-iot/dipper-engine-server/graph/model"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type SessionInfoResolver interface {
	ChanID(ctx context.Context, obj *ent.Session) (uint64, error)
	Infinity(ctx context.Context, obj *ent.Session) (bool, error)
	Chan(ctx context.Context, obj *ent.Session) (*ent.RuleChan, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _ErrorEngine_message(ctx context.Context, field graphql.CollectedField, obj *model.ErrorEngine) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ErrorEngine_message(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Message, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ErrorEngine_message(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ErrorEngine",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ErrorEngine_error_detail(ctx context.Context, field graphql.CollectedField, obj *model.ErrorEngine) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ErrorEngine_error_detail(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ErrorDetail, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ErrorEngine_error_detail(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ErrorEngine",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ErrorEngine_code(ctx context.Context, field graphql.CollectedField, obj *model.ErrorEngine) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ErrorEngine_code(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Code, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*int)
	fc.Result = res
	return ec.marshalOInt2ᚖint(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ErrorEngine_code(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ErrorEngine",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ListSessionInfo_total(ctx context.Context, field graphql.CollectedField, obj *model.ListSessionInfo) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ListSessionInfo_total(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Total, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ListSessionInfo_total(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ListSessionInfo",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ListSessionInfo_sessions(ctx context.Context, field graphql.CollectedField, obj *model.ListSessionInfo) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ListSessionInfo_sessions(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Sessions, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]*ent.Session)
	fc.Result = res
	return ec.marshalOSessionInfo2ᚕᚖgithubᚗcomᚋdipperᚑiotᚋdipperᚑengineᚑserverᚋentᚐSession(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ListSessionInfo_sessions(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ListSessionInfo",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_SessionInfo_id(ctx, field)
			case "chan_id":
				return ec.fieldContext_SessionInfo_chan_id(ctx, field)
			case "infinity":
				return ec.fieldContext_SessionInfo_infinity(ctx, field)
			case "chan":
				return ec.fieldContext_SessionInfo_chan(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type SessionInfo", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _SessionInfo_id(ctx context.Context, field graphql.CollectedField, obj *ent.Session) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SessionInfo_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(uint64)
	fc.Result = res
	return ec.marshalNUint642uint64(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_SessionInfo_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SessionInfo",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Uint64 does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _SessionInfo_chan_id(ctx context.Context, field graphql.CollectedField, obj *ent.Session) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SessionInfo_chan_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.SessionInfo().ChanID(rctx, obj)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(uint64)
	fc.Result = res
	return ec.marshalNUint642uint64(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_SessionInfo_chan_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SessionInfo",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Uint64 does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _SessionInfo_infinity(ctx context.Context, field graphql.CollectedField, obj *ent.Session) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SessionInfo_infinity(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.SessionInfo().Infinity(rctx, obj)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_SessionInfo_infinity(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SessionInfo",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _SessionInfo_chan(ctx context.Context, field graphql.CollectedField, obj *ent.Session) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SessionInfo_chan(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.SessionInfo().Chan(rctx, obj)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*ent.RuleChan)
	fc.Result = res
	return ec.marshalOChan2ᚖgithubᚗcomᚋdipperᚑiotᚋdipperᚑengineᚑserverᚋentᚐRuleChan(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_SessionInfo_chan(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SessionInfo",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Chan_id(ctx, field)
			case "name":
				return ec.fieldContext_Chan_name(ctx, field)
			case "description":
				return ec.fieldContext_Chan_description(ctx, field)
			case "root_node":
				return ec.fieldContext_Chan_root_node(ctx, field)
			case "nodes":
				return ec.fieldContext_Chan_nodes(ctx, field)
			case "status":
				return ec.fieldContext_Chan_status(ctx, field)
			case "created_at":
				return ec.fieldContext_Chan_created_at(ctx, field)
			case "updated_at":
				return ec.fieldContext_Chan_updated_at(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Chan", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _SessionOutput_chan_id(ctx context.Context, field graphql.CollectedField, obj *model.SessionOutput) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SessionOutput_chan_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ChanID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(uint64)
	fc.Result = res
	return ec.marshalNUint642uint64(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_SessionOutput_chan_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SessionOutput",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Uint64 does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _SessionOutput_session_id(ctx context.Context, field graphql.CollectedField, obj *model.SessionOutput) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SessionOutput_session_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.SessionID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(uint64)
	fc.Result = res
	return ec.marshalNUint642uint64(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_SessionOutput_session_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SessionOutput",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Uint64 does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _SessionOutput_form_engine(ctx context.Context, field graphql.CollectedField, obj *model.SessionOutput) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SessionOutput_form_engine(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.FormEngine, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_SessionOutput_form_engine(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SessionOutput",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _SessionOutput_data(ctx context.Context, field graphql.CollectedField, obj *model.SessionOutput) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SessionOutput_data(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Data, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(map[string]interface{})
	fc.Result = res
	return ec.marshalNMap2map(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_SessionOutput_data(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SessionOutput",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Map does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _SessionOutput_type(ctx context.Context, field graphql.CollectedField, obj *model.SessionOutput) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SessionOutput_type(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(model.OutputType)
	fc.Result = res
	return ec.marshalNOutputType2githubᚗcomᚋdipperᚑiotᚋdipperᚑengineᚑserverᚋgraphᚋmodelᚐOutputType(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_SessionOutput_type(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SessionOutput",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type OutputType does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _SessionOutput_error(ctx context.Context, field graphql.CollectedField, obj *model.SessionOutput) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SessionOutput_error(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Error, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*model.ErrorEngine)
	fc.Result = res
	return ec.marshalOErrorEngine2ᚖgithubᚗcomᚋdipperᚑiotᚋdipperᚑengineᚑserverᚋgraphᚋmodelᚐErrorEngine(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_SessionOutput_error(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SessionOutput",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "message":
				return ec.fieldContext_ErrorEngine_message(ctx, field)
			case "error_detail":
				return ec.fieldContext_ErrorEngine_error_detail(ctx, field)
			case "code":
				return ec.fieldContext_ErrorEngine_code(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type ErrorEngine", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputDebugFilter(ctx context.Context, obj interface{}) (model.DebugFilter, error) {
	var it model.DebugFilter
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"chan_id"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "chan_id":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("chan_id"))
			it.ChanID, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputSessionListRequest(ctx context.Context, obj interface{}) (model.SessionListRequest, error) {
	var it model.SessionListRequest
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"infinity", "skip", "limit"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "infinity":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("infinity"))
			it.Infinity, err = ec.unmarshalNBoolean2bool(ctx, v)
			if err != nil {
				return it, err
			}
		case "skip":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("skip"))
			it.Skip, err = ec.unmarshalNInt2int(ctx, v)
			if err != nil {
				return it, err
			}
		case "limit":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("limit"))
			it.Limit, err = ec.unmarshalNInt2int(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputSessionRequest(ctx context.Context, obj interface{}) (model.SessionRequest, error) {
	var it model.SessionRequest
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"chan_id", "is_test", "data"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "chan_id":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("chan_id"))
			it.ChanID, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "is_test":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("is_test"))
			it.IsTest, err = ec.unmarshalNBoolean2bool(ctx, v)
			if err != nil {
				return it, err
			}
		case "data":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("data"))
			it.Data, err = ec.unmarshalOMap2map(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var errorEngineImplementors = []string{"ErrorEngine"}

func (ec *executionContext) _ErrorEngine(ctx context.Context, sel ast.SelectionSet, obj *model.ErrorEngine) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, errorEngineImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("ErrorEngine")
		case "message":

			out.Values[i] = ec._ErrorEngine_message(ctx, field, obj)

		case "error_detail":

			out.Values[i] = ec._ErrorEngine_error_detail(ctx, field, obj)

		case "code":

			out.Values[i] = ec._ErrorEngine_code(ctx, field, obj)

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var listSessionInfoImplementors = []string{"ListSessionInfo"}

func (ec *executionContext) _ListSessionInfo(ctx context.Context, sel ast.SelectionSet, obj *model.ListSessionInfo) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, listSessionInfoImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("ListSessionInfo")
		case "total":

			out.Values[i] = ec._ListSessionInfo_total(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "sessions":

			out.Values[i] = ec._ListSessionInfo_sessions(ctx, field, obj)

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var sessionInfoImplementors = []string{"SessionInfo"}

func (ec *executionContext) _SessionInfo(ctx context.Context, sel ast.SelectionSet, obj *ent.Session) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, sessionInfoImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("SessionInfo")
		case "id":

			out.Values[i] = ec._SessionInfo_id(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
			}
		case "chan_id":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._SessionInfo_chan_id(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return innerFunc(ctx)

			})
		case "infinity":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._SessionInfo_infinity(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return innerFunc(ctx)

			})
		case "chan":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._SessionInfo_chan(ctx, field, obj)
				return res
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return innerFunc(ctx)

			})
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var sessionOutputImplementors = []string{"SessionOutput"}

func (ec *executionContext) _SessionOutput(ctx context.Context, sel ast.SelectionSet, obj *model.SessionOutput) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, sessionOutputImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("SessionOutput")
		case "chan_id":

			out.Values[i] = ec._SessionOutput_chan_id(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "session_id":

			out.Values[i] = ec._SessionOutput_session_id(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "form_engine":

			out.Values[i] = ec._SessionOutput_form_engine(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "data":

			out.Values[i] = ec._SessionOutput_data(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "type":

			out.Values[i] = ec._SessionOutput_type(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "error":

			out.Values[i] = ec._SessionOutput_error(ctx, field, obj)

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNListSessionInfo2githubᚗcomᚋdipperᚑiotᚋdipperᚑengineᚑserverᚋgraphᚋmodelᚐListSessionInfo(ctx context.Context, sel ast.SelectionSet, v model.ListSessionInfo) graphql.Marshaler {
	return ec._ListSessionInfo(ctx, sel, &v)
}

func (ec *executionContext) marshalNListSessionInfo2ᚖgithubᚗcomᚋdipperᚑiotᚋdipperᚑengineᚑserverᚋgraphᚋmodelᚐListSessionInfo(ctx context.Context, sel ast.SelectionSet, v *model.ListSessionInfo) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._ListSessionInfo(ctx, sel, v)
}

func (ec *executionContext) unmarshalNOutputType2githubᚗcomᚋdipperᚑiotᚋdipperᚑengineᚑserverᚋgraphᚋmodelᚐOutputType(ctx context.Context, v interface{}) (model.OutputType, error) {
	var res model.OutputType
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNOutputType2githubᚗcomᚋdipperᚑiotᚋdipperᚑengineᚑserverᚋgraphᚋmodelᚐOutputType(ctx context.Context, sel ast.SelectionSet, v model.OutputType) graphql.Marshaler {
	return v
}

func (ec *executionContext) marshalNSessionInfo2githubᚗcomᚋdipperᚑiotᚋdipperᚑengineᚑserverᚋentᚐSession(ctx context.Context, sel ast.SelectionSet, v ent.Session) graphql.Marshaler {
	return ec._SessionInfo(ctx, sel, &v)
}

func (ec *executionContext) marshalNSessionInfo2ᚖgithubᚗcomᚋdipperᚑiotᚋdipperᚑengineᚑserverᚋentᚐSession(ctx context.Context, sel ast.SelectionSet, v *ent.Session) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._SessionInfo(ctx, sel, v)
}

func (ec *executionContext) unmarshalNSessionListRequest2githubᚗcomᚋdipperᚑiotᚋdipperᚑengineᚑserverᚋgraphᚋmodelᚐSessionListRequest(ctx context.Context, v interface{}) (model.SessionListRequest, error) {
	res, err := ec.unmarshalInputSessionListRequest(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNSessionOutput2ᚖgithubᚗcomᚋdipperᚑiotᚋdipperᚑengineᚑserverᚋgraphᚋmodelᚐSessionOutput(ctx context.Context, sel ast.SelectionSet, v *model.SessionOutput) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._SessionOutput(ctx, sel, v)
}

func (ec *executionContext) unmarshalNSessionRequest2githubᚗcomᚋdipperᚑiotᚋdipperᚑengineᚑserverᚋgraphᚋmodelᚐSessionRequest(ctx context.Context, v interface{}) (model.SessionRequest, error) {
	res, err := ec.unmarshalInputSessionRequest(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalODebugFilter2ᚖgithubᚗcomᚋdipperᚑiotᚋdipperᚑengineᚑserverᚋgraphᚋmodelᚐDebugFilter(ctx context.Context, v interface{}) (*model.DebugFilter, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputDebugFilter(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOErrorEngine2ᚖgithubᚗcomᚋdipperᚑiotᚋdipperᚑengineᚑserverᚋgraphᚋmodelᚐErrorEngine(ctx context.Context, sel ast.SelectionSet, v *model.ErrorEngine) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._ErrorEngine(ctx, sel, v)
}

func (ec *executionContext) marshalOSessionInfo2ᚕᚖgithubᚗcomᚋdipperᚑiotᚋdipperᚑengineᚑserverᚋentᚐSession(ctx context.Context, sel ast.SelectionSet, v []*ent.Session) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalOSessionInfo2ᚖgithubᚗcomᚋdipperᚑiotᚋdipperᚑengineᚑserverᚋentᚐSession(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	return ret
}

func (ec *executionContext) marshalOSessionInfo2ᚖgithubᚗcomᚋdipperᚑiotᚋdipperᚑengineᚑserverᚋentᚐSession(ctx context.Context, sel ast.SelectionSet, v *ent.Session) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._SessionInfo(ctx, sel, v)
}

func (ec *executionContext) marshalOSessionOutput2ᚕᚖgithubᚗcomᚋdipperᚑiotᚋdipperᚑengineᚑserverᚋgraphᚋmodelᚐSessionOutputᚄ(ctx context.Context, sel ast.SelectionSet, v []*model.SessionOutput) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNSessionOutput2ᚖgithubᚗcomᚋdipperᚑiotᚋdipperᚑengineᚑserverᚋgraphᚋmodelᚐSessionOutput(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

// endregion ***************************** type.gotpl *****************************