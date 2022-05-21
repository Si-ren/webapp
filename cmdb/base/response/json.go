package response

type JsonResponse struct {
	Code   int         `json:"code,omitempty"`
	Msg    string      `json:"msg,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

func NewJsonResponse(code int, msg string, result interface{}) *JsonResponse {
	return &JsonResponse{
		Code:   code,
		Msg:    msg,
		Result: result,
	}
}

var (
	Unauthorzation = NewJsonResponse(401, "Unauthorzation", nil)
	OK             = NewJsonResponse(200, "OK", nil)
	BadRequest     = NewJsonResponse(400, "Bad Request", nil)
)
