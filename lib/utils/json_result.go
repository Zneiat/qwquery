package utils

// JSONResult JSON 响应数据结构
type JSONResult struct {
	Success bool        `json:"success"` // 是否成功
	Code    int         `json:"code"`    // 相应代码
	Msg     string      `json:"msg"`     // 消息
	Data    interface{} `json:"data"`    // 数据
}

// PageResult JSON 分页数据结构体
type PageResult struct {
	Page     int           `json:"page"`
	PageSize int           `json:"pageSize"`
	Total    int           `json:"total"`
	LastPage int           `json:"lastPage"`
	List     []interface{} `json:"list"`
}

// JSON is normal json result
func JSON(code int, msg string, data interface{}, success bool) *JSONResult {
	return &JSONResult{
		Success: success,
		Code:    code,
		Msg:     msg,
		Data:    data,
	}
}

// JSONData is just response data
func JSONData(data interface{}) *JSONResult {
	return &JSONResult{
		Success: true,
		Code:    RespCodeOK,
		Data:    data,
	}
}

// JSONPageData is 响应 json 分页数据
func JSONPageData(data *PageResult) *JSONResult {
	return &JSONResult{
		Success: true,
		Code:    RespCodeOK,
		Data:    data,
	}
}

// JSONSuccess is just response success
func JSONSuccess() *JSONResult {
	return &JSONResult{
		Success: true,
		Code:    RespCodeOK,
	}
}

// JSONError is just response error
func JSONError(code int, msg string) *JSONResult {
	return &JSONResult{
		Success: false,
		Code:    code,
		Msg:     msg,
	}
}
