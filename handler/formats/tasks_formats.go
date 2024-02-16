package formats

// format for post request
var JsonFormat struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			Task_name string `json:"task_name"`
			Completed bool   `json:"Completed"`
		} `json:"attributes"`
		Relationships struct {
			User struct {
				Id_User *int `json:"id"`
			} `json:"user"`
		} `json:"relationships"`
	} `json:"data"`
}
