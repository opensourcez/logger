package logger

// DBREventReceiver is a sentinel EventReceiver.
// Use it if the caller doesn't supply one.
type DBREventReceiver struct {
	LogTag     string
	ShowInfo   bool
	ShowErrors bool
	ShowTiming bool
	OPID       string
	AddToChain bool
}

// Event receives a simple notification when various events occur.
func (d *DBREventReceiver) Event(eventName string) {
	if d.ShowInfo {
		event := GenericMessage("QUERY EVENT")
		event.Query = eventName
		event.LogLevel = "INFO"
		event.LogTag = d.LogTag
		event.Operation = Operation{ID: d.OPID}
		if d.AddToChain {
			internalLogger.AddToChain(d.OPID, *event)
		} else {
			internalLogger.INFO(*event, d.LogTag)
		}
	}
}

// EventKv receives a notification when various events occur along with
// optional key/value data.
func (d *DBREventReceiver) EventKv(eventName string, kvs map[string]string) {
	if d.ShowInfo {
		event := GenericMessage("QUERY EVENT")
		event.Query = eventName
		event.Labels = kvs
		event.LogLevel = "INFO"
		event.Operation = Operation{ID: d.OPID}
		event.LogTag = d.LogTag
		if d.AddToChain {
			internalLogger.AddToChain(d.OPID, *event)
		} else {
			internalLogger.INFO(*event, d.LogTag)
		}
	}
}

// EventErr receives a notification of an error if one occurs.
func (d *DBREventReceiver) EventErr(eventName string, err error) error {
	event := ParsePGError(err)
	event.Query = eventName
	event.LogLevel = "ERROR"
	event.Operation = Operation{ID: d.OPID}
	event.LogTag = d.LogTag
	if d.ShowErrors {
		if d.AddToChain {
			internalLogger.AddToChain(d.OPID, *event)
		} else {
			internalLogger.INFO(*event, d.LogTag)
		}
	}
	return event
}

// EventErrKv receives a notification of an error if one occurs along with
// optional key/value data.
func (d *DBREventReceiver) EventErrKv(eventName string, err error, kvs map[string]string) error {
	event := ParsePGError(err)
	event.Query = eventName
	event.Labels = kvs
	event.Operation = Operation{ID: d.OPID}
	event.LogLevel = "ERROR"
	event.LogTag = d.LogTag
	if d.ShowErrors {
		if d.AddToChain {
			internalLogger.AddToChain(d.OPID, *event)
		} else {
			internalLogger.INFO(*event, d.LogTag)
		}
	}
	return event
}

// Timing receives the time an event took to happen.
func (d *DBREventReceiver) Timing(eventName string, nanoseconds int64) {
	if d.ShowTiming {
		event := GenericMessage("QUERY TIMING")
		event.Query = eventName
		event.QueryTiming = nanoseconds
		event.LogLevel = "INFO"
		event.Operation = Operation{ID: d.OPID}
		event.LogTag = d.LogTag
		if d.AddToChain {
			internalLogger.AddToChain(d.OPID, *event)
		} else {
			internalLogger.INFO(*event, d.LogTag)
		}
	}
}

// TimingKv receives the time an event took to happen along with optional key/value data.
func (d *DBREventReceiver) TimingKv(eventName string, nanoseconds int64, kvs map[string]string) {
	if d.ShowTiming {
		event := GenericMessage("QUERY TIMING")
		event.Query = eventName
		event.QueryTiming = nanoseconds
		event.Labels = kvs
		event.Operation = Operation{ID: d.OPID}
		event.LogLevel = "INFO"
		event.LogTag = d.LogTag
		if d.AddToChain {
			internalLogger.AddToChain(d.OPID, *event)
		} else {
			internalLogger.INFO(*event, d.LogTag)
		}
	}
}
