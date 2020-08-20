package logger

const (
	LogLevelEmergency = "EMERGENCY"
	LogLevelError     = "ERROR"
	LogLevelCritical  = "CRITICAL"
	LogLevelAlert     = "ALERT"
	LogLevelWarning   = "WARNING"
	LogLevelNotice    = "NOTICE"
	LogLevelInfo      = "INFO"
)

const (
	// examples, we need NewObject codes.
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

func BadRequest(original error, message string) *InformationConstruct {
	error := NewObject(message, HTTPCodeBadRequest)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}

func MissingCookie(original error) *InformationConstruct {
	error := NewObject(MessageMissingCookie, HTTPCodeUnAuthorized)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}

func MissingJWTHeader(original error) *InformationConstruct {
	error := NewObject(MessageMissingJWTHeader, HTTPCodeUnAuthorized)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}

func CouldNotGeneratePassword(original error) *InformationConstruct {
	error := NewObject(MessageCouldNotGeneratePassword, HTTPCodeBadRequest)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func RouteNotFound(original error) *InformationConstruct {
	error := NewObject(MessageRouteNotFound, HTTPCodeNotFound)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func InvalidUniqueIdentifier(original error) *InformationConstruct {
	error := NewObject(MessageInvalidUniqueIdentifier, HTTPCodeBadRequest)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func RecordNotFound404(original error) *InformationConstruct {
	error := NewObject(MessageRecordNotFound404, HTTPCodeNotFound)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func RecordNotFound(original error) *InformationConstruct {
	error := NewObject("", HTTPCodeNoContent)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func BadEmailOrPassword(original error) *InformationConstruct {
	error := NewObject(MessageBadEmailOrPassword, HTTPCodeUnAuthorized)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func Unauthorized(original error) *InformationConstruct {
	error := NewObject(MessageUnauthorized, HTTPCodeUnAuthorized)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func LoginExpired(original error) *InformationConstruct {
	error := NewObject(MessageLoginExpired, HTTPCodeLoginExpired)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func UnauthorizedCustomMessage(original error, message string) *InformationConstruct {
	error := NewObject(message, HTTPCodeUnAuthorized)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func DisabledUser(original error) *InformationConstruct {
	error := NewObject(MessageDisabledUser, HTTPCodeUnAuthorized)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func TooManyPasswordResetAttemps(original error) *InformationConstruct {
	error := NewObject(MessageTooManyPasswordResetAttemps, HTTPCodeUnAuthorized)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func ListNotAccepted(original error) *InformationConstruct {
	error := NewObject(MessageListNotAccepted, HTTPCodeBadRequest)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func GeneralDecoding(original error) *InformationConstruct {
	error := NewObject(MessageGeneralDecoding, HTTPCodeBadRequest)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func JSONDecoding(original error) *InformationConstruct {
	error := NewObject(MessageJSONDecoding, HTTPCodeBadRequest)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func XMLDecoding(original error) *InformationConstruct {
	error := NewObject(MessageXMLDecoding, HTTPCodeBadRequest)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func JSONEncoding(original error) *InformationConstruct {
	error := NewObject(MessageJSONEncoding, HTTPCodeInternalServerError)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func GenericEncoding(original error) *InformationConstruct {
	error := NewObject(MessageGenericEncoding, HTTPCodeInternalServerError)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func XMLEncoding(original error) *InformationConstruct {
	error := NewObject(MessageXMLEncoding, HTTPCodeInternalServerError)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func GenericError(original error) *InformationConstruct {
	error := NewObject("", HTTPCodeInternalServerError)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func GenericErrorWithMessage(original error, message string) *InformationConstruct {
	error := NewObject(message, HTTPCodeInternalServerError)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func GenericMessage(message string) *InformationConstruct {
	error := NewObject(message, 0)
	return &error
}
func UnexpectedJWTSigningMethod(signingMethodInUse string) *InformationConstruct {
	error := NewObject(MessageUnexpectedJWTSigningMethod+signingMethodInUse, HTTPCodeBadRequest)
	error.ReturnToClient = true
	error.OriginalError = nil
	return &error
}
func MissingPassword(original error) *InformationConstruct {
	error := NewObject(MessageMissingPassword, HTTPCodeBadRequest)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func UniqueKeyConstraint(original error) *InformationConstruct {
	error := NewObject(original.Error(), HTTPCodeBadRequest)
	error.ReturnToClient = true
	error.OriginalError = original
	return &error
}
func DatabaseConnectionErrror(original error) *InformationConstruct {
	error := NewObject(MessageDatabaseConnectionErrror, HTTPCodeInternalServerError)
	error.ReturnToClient = false
	error.OriginalError = original
	return &error
}
