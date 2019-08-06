package logger

import (
	"log"
	"os"
)

func (f *FileClient) new(config *LoggingConfig) (err error) {

	f.Loggers = make(map[string]string)
	f.fileChannels = make(map[string]chan []byte)
	for _, v := range config.Logs {
		f.Loggers[v] = v
		f.fileChannels[v] = make(chan []byte)
	}
	f.Config = config
	return nil
}

func (f *FileClient) StartAndWatchFileChannels() {
	fileWatcherChannel := make(chan string)
	for _, tag := range f.Loggers {
		go f.startBufferToFile(tag, fileWatcherChannel)
	}
	for {
		tag := <-fileWatcherChannel
		if f.Loggers[tag] != "" {
			go f.startBufferToFile(tag, fileWatcherChannel)

		}
	}
}

func (f *FileClient) log(object *InformationConstruct, severity string, logTag string) {
	if object.StackTrace == "" {
		err := GetStack(f.Config, object)
		if err != nil {
			object.StackTrace = "Could not get stacktrace, error:" + err.Error()
		}
	}
	f.fileChannels[logTag] <- []byte(severity + ": " + object.JSON())
}

func (f *FileClient) close() {
	// no op
	for _, v := range f.Loggers {
		close(f.fileChannels[v])
	}
}

func (f *FileClient) makeSureFileIsOpen(tag string) (logfile *os.File) {
	if f.BaseLogFilePermissions == 0 {
		f.BaseLogFilePermissions = 0660
	}
	logfile, err := os.OpenFile(f.BaseLogFolder+"/"+tag, os.O_RDWR|os.O_APPEND|os.O_CREATE, f.BaseLogFilePermissions)
	if err != nil {
		// Starting software without a logger is a cardinal sin.
		log.Println("The logger could not open a file to log in .. this requires your attention !")
		panic(err)
	}
	return
}
func (f *FileClient) startBufferToFile(tag string, watcherChannel chan string) {
	file := f.makeSureFileIsOpen(tag)
	defer func(watcherChan chan string, tag string, file *os.File) {
		file.Close()
		if r := recover(); r != nil {
			watcherChan <- tag
		}
	}(watcherChannel, tag, file)
	for {
		file.Write(<-f.fileChannels[tag])
	}
}
