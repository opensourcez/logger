package logger

// TODO: finish this
func PrintAppoloVariables(vars map[string]interface{}) {
	var toPrint []interface{}
	for i, v := range vars {
		toPrint = append(toPrint, i, v)
	}
	//PrintINFO("APPOLO VARIBLES", toPrint)
}

// TODO: finish this
func PrintAppoloQuery(query map[string]interface{}) {
	var toPrint []interface{}
	toPrint = append(toPrint, "query", query["query"].(string))
	//sPrintINFO("APPOLO QUERY", toPrint)
}
