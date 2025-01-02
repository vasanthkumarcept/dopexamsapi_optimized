package handlers

import "os"

var Action string
var Remarks string
var UserErrorRemarks string = os.Getenv("USER_ERROR_REMARKS")
var emptyObject = struct{}{}
