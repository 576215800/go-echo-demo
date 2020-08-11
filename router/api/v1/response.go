package v1

type RestResponse struct{
	Status  int
	Code    int
	Msg     string
	Data 	interface{}
}