package errorx

import (
	"moodMoment-server/common/global"
	"net/http"

	"go.uber.org/zap"
)

/** Errorx 对接前端antpro框架的接口格式规范，用于实现error接口
 */
type Errorx struct {
	Success      bool   `json:"success"`      // 【必需】请求状态
	ErrorCode    int    `json:"errorCode"`    // 【选填】错误类型代码：400错误请求，401未授权，500服务器内部错误，200成功
	ErrorMessage string `json:"errorMessage"` // 【选填】向用户显示消息
	TraceMessage string `json:"traceMessage"` // 【选填】调试错误信息，请勿在生产环境下使用，可有可无
	ShowType     int    `json:"showType"`     // 【选填】错误显示类型：0.不提示错误;1.警告信息提示；2.错误信息提示；4.通知提示；9.页面跳转
	TraceId      string `json:"traceId"`      // 【选填】方便后端故障排除：唯一的请求ID
	Host         string `json:"host"`         // 【选填】方便后端故障排除：当前访问服务器的主机
}

/** ResponseError 对接前端antpro框架的接口格式规范，用于实现展现错误
 */
type ResponseError struct {
	Success      bool   `json:"success"`      // 【必需】请求状态
	ErrorCode    int    `json:"errorCode"`    // 【选填】错误类型代码：400错误请求，401未授权，500服务器内部错误，200成功
	ErrorMessage string `json:"errorMessage"` // 【选填】向用户显示消息
	TraceMessage string `json:"traceMessage"` // 【选填】调试错误信息，请勿在生产环境下使用，可有可无
	ShowType     int    `json:"showType"`     // 【选填】错误显示类型：0.不提示错误;1.警告信息提示；2.错误信息提示；4.通知提示；9.页面跳转
	TraceId      string `json:"traceId"`      // 【选填】方便后端故障排除：唯一的请求ID
	Host         string `json:"host"`         // 【选填】方便后端故障排除：当前访问服务器的主机
}

const (
	ERROR_CODE_INTERNAL = 500 //内部错误
	ERROR_CODE_EXTERNAL = 400 //外部错误（默认）
	ERROR_CODE_AUTH     = 401 //外部错误（默认）
	SHOW_TYPE           = 2   //信息错误
)

// var TRACE_MODEL = isTraceModel() //调试模式
var ServerError = InternalError
var RequestError = ExternalError

//Error 返回错误信息，主要用于实现原生error接口，
//注意：打印errorx类型的时候，会调用该函数值，所以返回errorx，等于返回该Error()的结果
func (e *Errorx) Error() string {
	return e.ErrorMessage
}

//Data 返回errorx的值
//如果需要返回errorx的值怎么办？
//那就需要我们自己定义一个展现errorx结果的函数，使用ResponseError结构体
func (e *Errorx) Data() *ResponseError {
	return &ResponseError{
		Success:      e.Success,
		ErrorCode:    e.ErrorCode,
		ErrorMessage: e.ErrorMessage,
		TraceMessage: e.TraceMessage,
		ShowType:     e.ShowType,
		TraceId:      e.TraceId,
		Host:         e.Host,
	}
}

/** InternalError 内部错误 500
 * alertMsg string 提示信息
 * traceErr string 调试信息
 */
func InternalError(alertMsg string, traceErr error) error {
	return internalError(alertMsg, traceErr)
}

/** ExternalError 外部错误 400
 * alertMsg string 提示信息
 * traceMsg string 调试信息
 */
func ExternalError(alertMsg string, traceErr error) error {
	return externalError(alertMsg, traceErr)
}

/** AuthError 身份校验失败 401
 * alertMsg string 提示信息
 * traceMsg string 调试信息
 */
func AuthError(alertMsg string, traceErr error) error {
	return authError(alertMsg, traceErr)
}

func internalError(alertMsg string, traceErr error) *Errorx {
	errorx := newErrorx(alertMsg, traceErr)
	if len(errorx.ErrorMessage) == 0 {
		errorx.ErrorMessage = "服务器错误"
	}
	errorx.ErrorCode = ERROR_CODE_INTERNAL
	//非调试模式，统一报：服务器错误
	if !isTraceModel() {
		errorx.ErrorMessage = "服务器错误"
		errorx.TraceMessage = "服务器出现错误，请联系管理员licat233@gmail.com"
	}
	//内部错误要做日志处理
	global.LOG.Error("服务器错误", zap.Error(traceErr))
	return errorx
}

func externalError(alertMsg string, traceErr error) *Errorx {
	errorx := newErrorx(alertMsg, traceErr)
	if len(errorx.ErrorMessage) == 0 {
		errorx.ErrorMessage = "请求错误"
	}
	errorx.ErrorCode = ERROR_CODE_EXTERNAL
	return errorx
}

func authError(alertMsg string, traceErr error) *Errorx {
	errorx := newErrorx(alertMsg, traceErr)
	if len(errorx.ErrorMessage) == 0 {
		errorx.ErrorMessage = "身份校验失败"
	}
	errorx.ErrorCode = ERROR_CODE_AUTH
	return errorx
}

//newErrorx 新建Errorx实例，默认为外部请求错误
func newErrorx(alertMsg string, traceErr error) *Errorx {
	traceMsg := ""
	//调试模式下，展现错误调试信息
	if isTraceModel() && traceErr != nil {
		traceMsg = traceErr.Error()
	}
	errorx := &Errorx{
		Success:      false,
		ErrorCode:    ERROR_CODE_EXTERNAL,
		ErrorMessage: alertMsg,
		TraceMessage: traceMsg,
		ShowType:     SHOW_TYPE,
		TraceId:      "",
		Host:         "",
	}
	return errorx
}

/** ErrorToErrorX 对接go-zero的错误处理函数，将原生错误转化为 Errorx，默认为ExternalError
 * inerr 传入的错误
 */
func ErrorToErrorX(inerr error) (int, interface{}) {
	if inerr == nil {
		//没有错误，应该正常返回
		return http.StatusOK, &struct {
			Success bool   `json:"success"` // 请求状态。
			Message string `json:"message"` // 响应消息
		}{
			Success: true,
			Message: "请求成功",
		}
	}
	//类型断言看是否为 Errorx 类型
	if errorx, ok := inerr.(*Errorx); ok {
		//无需处理，直接返回
		return http.StatusOK, errorx.Data()
	}

	//原生错误，我们默认把原生error作为ExternalError返回
	// return http.StatusOK, externalError(inerr.Error(), nil)
	return http.StatusOK, externalError(inerr.Error(), nil).Data()
}

func isTraceModel() bool {
	return true
}
