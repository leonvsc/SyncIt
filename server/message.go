package main

func parseMessage(msg string) string {

	switch msg {
	case "STOP":
		return "STOP"
	case "SAVE":
		return "File saved!"
	case "GET":
		return "Returning file..."
	case "UPLOAD":
		return "Uploading file..."
	default:
		return ""
	}
}
