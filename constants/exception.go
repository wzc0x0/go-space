package constants

import "net/http"

const (
	SUCCESS       = 200
	Redirect	  = http.StatusMovedPermanently
	ERROR         = 500
	InvalidParams = 400
	ErrorAuth = 401
	ErrorToken = 402
	FORBIDDEN 	= 403
	NotFound  = 404
	MethodNotAllowed = 405
)