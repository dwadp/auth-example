package response

type (
	JSON map[string]interface{}
)

func OK(message string, data interface{}) JSON {
	return JSON{
		"status":  true,
		"message": message,
		"data":    data,
	}
}

func Error(message string, err error) JSON {
	errMessage := ""

	if err != nil {
		errMessage = err.Error()
	}

	return JSON{
		"status":  false,
		"message": message,
		"error":   errMessage,
	}
}

func WithStatus(message string, status bool, data interface{}) JSON {
	return JSON{
		"status":  status,
		"message": message,
		"data":    data,
	}
}
