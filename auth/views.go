package auth


func F() int {
	return 120
}

type WebResponce struct {
	Success 	bool 	`json:"success"`
}

// resp := WebResponce{true, nil}
