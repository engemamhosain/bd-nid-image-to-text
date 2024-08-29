package models

type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var Response = make(map[string]APIResponse)

func init() {

	Response["catch_error"] = APIResponse{4000, "Error from catch block"}
	Response["token_not_allowed"] = APIResponse{403, "Token not allowed"}
	Response["text_not_found"] = APIResponse{4000, "Text not found"}
	Response["image_not_valid"] = APIResponse{4000, "Image not valid.Please upload correct image"}
	Response["field_missing"] = APIResponse{4001, "Field missing"}

}
func GetCode(message string) APIResponse {
	return Response[message]
}
