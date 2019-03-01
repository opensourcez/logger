package loggernew

const (
	// examples, we need NewError codes.
	CodeTransactionCanceled = 1003
	CodeTransactionFailed   = 1002
	CodeTransactionSuccess  = 1001
	CodeTransactionPending  = 1000
)

const (
	HTTPCodeBadRequest          = 400
	HTTPCodeUnAuthorized        = 401
	HTTPCodeNotFound            = 404
	HTTPCodeNoContent           = 204
	HTTPCodeLoginExpired        = 400
	HTTPCodeInternalServerError = 500
)

const (
	MessageMissingCookie               = "You request has no authentication cookie"
	MessageMissingJWTHeader            = "You request has no JWT"
	MessageCouldNotGeneratePassword    = "We could not generate a password for you"
	MessageRouteNotFound               = "We could not find the path your are looking for"
	MessageInvalidUniqueIdentifier     = "Invalid Uniquer Identifier (UUID)"
	MessageRecordNotFound404           = "Record not found"
	MessageBadEmailOrPassword          = "Email or Password incorrect"
	MessageUnauthorized                = "You are not authorized to preform this action"
	MessageLoginExpired                = "You're login credentials have expired"
	MessageDisabledUser                = "This user has been disabled"
	MessageTooManyPasswordResetAttemps = "You have attemped to reset your password too many times"
	MessageListNotAccepted             = "This method does not accept an array/list"
	MessageGeneralDecoding             = "We could not decode your data"
	MessageJSONDecoding                = "We could not decode your json object"
	MessageXMLDecoding                 = "We could not decode your xml object"
	MessageJSONEncoding                = "Could not encode json"
	MessageGenericEncoding             = "Could not encode xml or json"
	MessageXMLEncoding                 = "Could not encode xml"
	MessageUnexpectedJWTSigningMethod  = "Unexpected JWT signing method, this system uses: "
	MessageMissingPassword             = "A password is required"
	MessageDatabaseConnectionErrror    = "Could not connect to database"
)

func BadRequest(original error, message string) *ErrorConstruct {
	error := NewError(message, HTTPCodeBadRequest)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}

func MissingCookie(original error) *ErrorConstruct {
	error := NewError(MessageMissingCookie, HTTPCodeUnAuthorized)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}

func MissingJWTHeader(original error) *ErrorConstruct {
	error := NewError(MessageMissingJWTHeader, HTTPCodeUnAuthorized)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}

func CouldNotGeneratePassword(original error) *ErrorConstruct {
	error := NewError(MessageCouldNotGeneratePassword, HTTPCodeBadRequest)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func RouteNotFound(original error) *ErrorConstruct {
	error := NewError(MessageRouteNotFound, HTTPCodeNotFound)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func InvalidUniqueIdentifier(original error) *ErrorConstruct {
	error := NewError(MessageInvalidUniqueIdentifier, HTTPCodeBadRequest)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func RecordNotFound404(original error) *ErrorConstruct {
	error := NewError(MessageRecordNotFound404, HTTPCodeNotFound)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func RecordNotFound(original error) *ErrorConstruct {
	error := NewError("", HTTPCodeNoContent)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func BadEmailOrPassword(original error) *ErrorConstruct {
	error := NewError(MessageBadEmailOrPassword, HTTPCodeUnAuthorized)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func Unauthorized(original error) *ErrorConstruct {
	error := NewError(MessageUnauthorized, HTTPCodeUnAuthorized)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func LoginExpired(original error) *ErrorConstruct {
	error := NewError(MessageLoginExpired, HTTPCodeLoginExpired)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func UnauthorizedCustomMessage(original error, message string) *ErrorConstruct {
	error := NewError(message, HTTPCodeUnAuthorized)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func DisabledUser(original error) *ErrorConstruct {
	error := NewError(MessageDisabledUser, HTTPCodeUnAuthorized)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func TooManyPasswordResetAttemps(original error) *ErrorConstruct {
	error := NewError(MessageTooManyPasswordResetAttemps, HTTPCodeUnAuthorized)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func ListNotAccepted(original error) *ErrorConstruct {
	error := NewError(MessageListNotAccepted, HTTPCodeBadRequest)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func GeneralDecoding(original error) *ErrorConstruct {
	error := NewError(MessageGeneralDecoding, HTTPCodeBadRequest)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func JSONDecoding(original error) *ErrorConstruct {
	error := NewError(MessageJSONDecoding, HTTPCodeBadRequest)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func XMLDecoding(original error) *ErrorConstruct {
	error := NewError(MessageXMLDecoding, HTTPCodeBadRequest)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func JSONEncoding(original error) *ErrorConstruct {
	error := NewError(MessageJSONEncoding, HTTPCodeInternalServerError)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func GenericEncoding(original error) *ErrorConstruct {
	error := NewError(MessageGenericEncoding, HTTPCodeInternalServerError)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func XMLEncoding(original error) *ErrorConstruct {
	error := NewError(MessageXMLEncoding, HTTPCodeInternalServerError)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func GenericError(original error) *ErrorConstruct {
	error := NewError("", HTTPCodeInternalServerError)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func GenericErrorWithMessage(original error, message string) *ErrorConstruct {
	error := NewError(message, HTTPCodeInternalServerError)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func UnexpectedJWTSigningMethod(signingMethodInUse string) *ErrorConstruct {
	error := NewError(MessageUnexpectedJWTSigningMethod+signingMethodInUse, HTTPCodeBadRequest)
	error.ReturnToClient = true
	error.OriginalError = nil
	return &error
}
func MissingPassword(original error) *ErrorConstruct {
	error := NewError(MessageMissingPassword, HTTPCodeBadRequest)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func UniqueKeyConstraint(original error) *ErrorConstruct {
	error := NewError(original.Error(), HTTPCodeBadRequest)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func DatabaseConnectionErrror(original error) *ErrorConstruct {
	error := NewError(MessageDatabaseConnectionErrror, HTTPCodeInternalServerError)
	error.ReturnToClient = false
	error.OriginalError = original
	return &error
}
