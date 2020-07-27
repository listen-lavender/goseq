package api

import "errors"

var (
	ErrParseParameter         = errors.New("参数解析错误")
	ErrNotLogin               = errors.New("您当前未登录,请登录后再试!")
	ErrSystemErrorCode uint32 = 1
	ErrSystemErrorMsg         = "系统繁忙,请稍后再试!" // 不是自定义的错误码 不能直接返回给用户
)
