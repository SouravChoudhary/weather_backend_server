package apperror

func ReturnStoreErr(errCode string, err error) Error {
	if storeErr, exist := ErrorMap[errCode]; exist {
		storeErr.Details = map[string]interface{}{"db_err": err.Error()}
		return storeErr
	}
	return ErrorCodeNotImplemented
}

func ReturnServiceErr(errCode string, err error) Error {
	if svcErr, exist := ErrorMap[errCode]; exist {
		svcErr.Details = map[string]interface{}{"svc_err": err.Error()}
		return svcErr
	}
	return ErrorCodeNotImplemented
}
