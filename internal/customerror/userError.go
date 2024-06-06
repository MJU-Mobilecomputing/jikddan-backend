package customerror

import "net/http"

var (
	USER_CREATION_FAIL = NewError(http.StatusUnprocessableEntity, "user creation fail.", "101")
	USER_NOT_FOUND     = NewError(http.StatusNotFound, "user not found", "102")
)

func UserCreationFail(err error) error {
	userCreationFail := USER_CREATION_FAIL
	userCreationFail.Data = err
	return userCreationFail
}

func UserNotFound(err error) error {
	userNotFound := USER_NOT_FOUND
	userNotFound.Data = err
	return userNotFound
}
