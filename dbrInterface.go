package logger

// DBREventReceiver is a sentinel EventReceiver.
// Use it if the caller doesn't supply one.
type DBREventReceiver struct {
	LogTag     string
	ShowInfo   bool
	ShowErrors bool
	ShowTiming bool
}

// Event receives a simple notification when various events occur.
func (d *DBREventReceiver) Event(eventName string) {
	if d.ShowInfo {
		event := GenericMessage("QUERY EVENT")
		event.Query = eventName
		internalLogger.INFO(*event, d.LogTag)
	}
}

// EventKv receives a notification when various events occur along with
// optional key/value data.
func (d *DBREventReceiver) EventKv(eventName string, kvs map[string]string) {
	if d.ShowInfo {
		event := GenericMessage("QUERY EVENT")
		event.Query = eventName
		event.Labels = kvs
		internalLogger.INFO(*event, d.LogTag)
	}
}

// EventErr receives a notification of an error if one occurs.
func (d *DBREventReceiver) EventErr(eventName string, err error) error {
	if d.ShowErrors {
		event := GenericMessage("QUERY ERROR")
		event.Query = eventName
		event.OriginalError = err
		internalLogger.ERROR(*event, d.LogTag)
	}
	return err
}

// EventErrKv receives a notification of an error if one occurs along with
// optional key/value data.
func (d *DBREventReceiver) EventErrKv(eventName string, err error, kvs map[string]string) error {
	if d.ShowErrors {
		event := GenericMessage("QUERY ERROR")
		event.Query = eventName
		event.Labels = kvs
		event.OriginalError = err
		internalLogger.ERROR(*event, d.LogTag)
	}
	return err
}

// Timing receives the time an event took to happen.
func (d *DBREventReceiver) Timing(eventName string, nanoseconds int64) {
	if d.ShowTiming {
		event := GenericMessage("QUERY TIMING")
		event.Query = eventName
		event.QueryTiming = nanoseconds
		internalLogger.INFO(*event, d.LogTag)
	}
}

// TimingKv receives the time an event took to happen along with optional key/value data.
func (d *DBREventReceiver) TimingKv(eventName string, nanoseconds int64, kvs map[string]string) {
	if d.ShowTiming {
		event := GenericMessage("QUERY TIMING")
		event.Query = eventName
		event.QueryTiming = nanoseconds
		event.Labels = kvs
		internalLogger.INFO(*event, d.LogTag)
	}
}
