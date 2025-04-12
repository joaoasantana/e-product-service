package response

type (
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	Failure struct {
		Status Status `json:"status"`
		Error  string `json:"error"`
	}

	Success struct {
		Status Status `json:"status"`
		Data   any    `json:"data"`
	}
)
