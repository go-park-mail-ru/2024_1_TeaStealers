package logger

/*
const permissions = 0o644

func openFile(name string) (*os.File, error) {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, permissions)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func New(cfg config.Logger) (*slog.Logger, error) {
	var level slog.Level
	switch cfg.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	}

	var out io.Writer
	switch cfg.Sink {
	case "stdout":
		out = os.Stdout
	default:
		file, err := openFile(cfg.Sink)
		if err != nil {
			return nil, err
		}
		out = file
	}

	opts := HandlerOpts{
		level: level,
		out:   out,
	}
	handler := NewHandler(&opts)
	l := slog.New(handler)

	return l, nil
}
*/
