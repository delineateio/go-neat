package logging

func InitialiseLogging() {

	NewLogConfig().
		addConsole().
		addFile().
		configure().
        log()
}
