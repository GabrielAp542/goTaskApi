package mappermissions

var ConfigPermissions = map[string]map[string]string{
	"/tasks": {
		"GET":    "public",
		"POST":   "tasks:create",
		"PUT":    "tasks:update",
		"DELETE": "tasks:delete",
	},
	"/tasks:id": {
		"GET": "tasks:get_id",
	},
}
