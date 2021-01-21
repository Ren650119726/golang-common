package response

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
)

// 定义错误码
const (
	RESPONSE_SUCCESS        int = 0     //成功
	RESPONSE_UNKNOWN        int = 20000 //未知错误,服务不可用
	RESPONSE_AUTH_ERROR     int = 20001 //授权权限错误
	RESPONSE_PARAM_ERROR    int = 20002 //非法的参数
	RESPONSE_PERMIT_ERROR   int = 20003 //权限不足
	RESPONSE_BUSINESS_ERROR int = 30000 //业务处理失败
	RESPONSE_LIVE_ERROR     int = 40000 //处理失败
)
const (
	ERR_ALL_NUMBER_REGISTER = 10001 // 注册账号为纯数字
	ERR_HASE_AT_REGISTER    = 10002 // 注册账号含@符号
	ERR_BIND_PHONE          = 10003 // 该手机号已被绑定
	ERR_BIND_EMAIL          = 10004 // 该邮箱已被绑定
	ERR_PHONE_ILLEGAL       = 10005 // 该手机号不规范
	ERR_USERNAME_LENGTH     = 10006 // 账号长度不合法
	ERR_PASSWORD_LENGTH     = 10007 // 密码长度不合法
)

const (
	InnerErrorMsg         = "系统繁忙,请稍后再试"
	InnerErrorCode        = http.StatusServiceUnavailable
	NeedSuperAdmin        = "需要超级管理员权限"
	AccountAbnormal       = "账号状态异常,请尝试重新登陆"
	AccountWrongISS       = "错误的颁发机构"
	AlreadyRegistered     = "用户名已被注册"
	LoginFail             = "用户名或密码错误;或用户状态异常"
	LoginCaptchaFail      = "请输入正确的验证码"
	ModifySelfAvatar404   = "uid不匹配,可能需要重新登陆"
	ModifySelfNickName404 = ModifySelfAvatar404
	Common404             = "资源未找到"
	Admin404              = "没有符合条件的管理员; 1.可能是超级管理员; 2.已被删除; 3.传入的uid错误"
	Frozen404             = "没有符合条件的管理员; 1.可能是超级管理员; 2.已被删除或冻结; 3.传入的uid错误"
	Unfrozen404           = "没有符合条件的管理员; 1.可能是超级管理员; 2.不是冻结状态; 3.传入的uid错误"
	PasswordIllegal       = "密码只能是6到20位的字母和数字"
	NicknameIllegal       = "昵称只能是1到20位的字母和数字"
	PasswordBase64Illegal = "错误的密码格式"
	UsernameIllegal       = "用户名只能是6到20位的字母和数字,且不能全为数字"
)

//从struct里获取某个字段转为json时的字段名
func getJsonTagName(arg interface{}, fieldName string) string {
	field, ok := reflect.TypeOf(arg).Elem().FieldByName(fieldName)
	if !ok {
		return fieldName
	}
	fieldJSONName, ok := field.Tag.Lookup("form")
	if !ok {
		return fieldName
	}
	return fieldJSONName
}

//返回验证失败的错误信息，从错误信息里提取错误字段，msg为返回给前端的提示信息
func WrapValidateFailed(object interface{}, err error, msg string) HttpInnerResult {
	m := make(map[string]interface{})
	for _, item := range err.(validator.ValidationErrors) {
		m[getJsonTagName(object, item.Field())] = item.Value()
	}
	return HttpInnerResult{http.StatusBadRequest, m, msg}
}
