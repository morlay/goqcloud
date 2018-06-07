package scf

// 运行函数的返回

type Result struct {
	// 表示函数的计费耗时，单位是毫秒，异步调用返回为空
	BillDuration int64 `json:"BillDuration"`
	// 表示执行函数的耗时，单位是毫秒，异步调用返回为空
	Duration float64 `json:"Duration"`
	// 表示执行函数的错误返回信息，异步调用返回为空
	ErrMsg string `json:"ErrMsg"`
	// 此次函数执行的Id
	FunctionRequestId string `json:"FunctionRequestId"`
	// 0为正确，异步调用返回为空
	InvokeResult int64 `json:"InvokeResult"`
	// 表示执行过程中的日志输出，异步调用返回为空
	Log string `json:"Log"`
	// 执行函数时的内存大小，单位为Byte，异步调用返回为空
	MemUsage int64 `json:"MemUsage"`
	// 表示执行函数的返回，异步调用返回为空
	RetMsg string `json:"RetMsg"`
}
