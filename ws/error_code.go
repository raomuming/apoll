package ws

const (
	OK                 = 200
	NotLoggedIn        = 1000
	ParameterIllegal   = 1001
	UnauthorizedUserId = 1002
	Unauthorized       = 1003
	ServerError        = 1004
	NotData            = 1005
	ModelAddError      = 1006
	ModelDeleteError   = 1007
	ModelStoreError    = 1008
	OperationFailure   = 1009
	RoutingNotExist    = 1010
)

func GetErrorMessage(code uint32, message string) string {
	var codeMessage string
	codeMap := map[uint32]string{
		OK:                 "Success",
		NotLoggedIn:        "not logged in",
		ParameterIllegal:   "parameter illegal",
		UnauthorizedUserId: "unauthorized user id",
		Unauthorized:       "unauthorized",
		ServerError:        "server error",
		NotData:            "not data",
		ModelAddError:      "model add error",
		ModelDeleteError:   "model delete error",
		ModelStoreError:    "model store error",
		OperationFailure:   "operation failure",
		RoutingNotExist:    "routing not exist",
	}

	if message == "" {
		if value, ok := codeMap[code]; ok {
			codeMessage = value
		} else {
			codeMessage = "not defined error type!"
		}
	} else {
		codeMessage = message
	}

	return codeMessage
}
