package app

const (
	SUCCESS         = 200
	ERROR           = 500
	INVALID_PARAMS  = 400
)

var Message = map[int]string{
	SUCCESS:         "ok",
	ERROR:           "fail",
	INVALID_PARAMS:  "invalid params",
}

var Status = map[int]string{
	SUCCESS:         "success",
	ERROR:           "fail",
	INVALID_PARAMS:  "invalid params",
}

func GetStatus(code int) string   {
	return Status[GetHttpStatus(code)]
}

func GetMsg(code int) string {
	if(code==0){
		return Message[SUCCESS]
	}else{
		return Message[code]
	}
}

func GetHttpStatus(code int)  int {
	if(code==0){
		return SUCCESS
	}
	if(code<500){
		return INVALID_PARAMS
	}else{
		return ERROR
	}
}