package customerror

import "net/http"

var DIARYMENU_CREATION_FAIL = NewError(http.StatusUnprocessableEntity, "diary menu creation fail", "201")

func DiaryMenuCreationFail(err error) error {
	diaryMenuCreationFail := DIARYMENU_CREATION_FAIL
	diaryMenuCreationFail.Data = err
	return diaryMenuCreationFail
}
