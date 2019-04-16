package logger

// TODO: add more codes.
func ParsePGError(code string, customMessage string) (outError *InformationConstruct) {
	switch code {
	case "42703": // Column not found error
		outError = BadRequest(nil, customMessage)
		outError.Hint = "Unknown column"
	case "42601": // bad syntax error
		outError = BadRequest(nil, customMessage)
		outError.Hint = "Your syntax might be off, review all your column and table references."
	case "22P02": // bad syntax error for UUID
		outError = BadRequest(nil, customMessage)
		outError.Hint = "You have an inalid UUID in your database transaction"
	case "42P01": // table not found
		outError = BadRequest(nil, customMessage)
		outError.Hint = "You are trying to save data to a table that does not exist, double check your table names"
	default:
		outError = BadRequest(nil, code+":"+customMessage)
	}
	return
}
