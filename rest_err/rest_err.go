package rest_err

type RestErr struct {
	Message string `json:"message"`
	Err     string `json:"error"`
	Status  int    `json:"status"`
	Causes  []Causes
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewRestErr(message string, status int, err string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Status:  status,
		Err:     err,
		Causes:  causes,
	}
}
