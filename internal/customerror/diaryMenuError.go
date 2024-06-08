package customerror

import "net/http"

var DIARYMENU_CREATION_FAIL = NewError(http.StatusUnprocessableEntity, "diary menu creation fail", "201")

var DIARYMENU_NOT_FOUND = NewError(http.StatusNotFound, "diary menu not found", "404")

func DiaryMenuCreationFail(err error) error {
	diaryMenuCreationFail := DIARYMENU_CREATION_FAIL
	diaryMenuCreationFail.Data = err
	return diaryMenuCreationFail
}

func DiaryMenuNotFound(err error) error {
	diaryMenuNotFound := DIARYMENU_NOT_FOUND
	diaryMenuNotFound.Data = err
	return nil
}
