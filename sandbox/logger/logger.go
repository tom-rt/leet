package logger

import (
	"log/slog"
	"os"
)

func GetReqLogger(path string) *slog.Logger {
	attrs := []slog.Attr{
		{Key: "path", Value: slog.StringValue(path)},
	}
	handler := slog.NewTextHandler(os.Stdout, nil).WithAttrs(attrs)
	logger := slog.New(handler)
	return logger
}

// func GetReqLogger(path string) *slog.Logger {
// 	reqID := uuid.New().String()
// 	loggerAttributes := []slog.Attr{
// 		slog.Attr{
// 			Key:   "requestID",
// 			Value: slog.StringValue(reqID),
// 		},
// 		slog.Attr{
// 			Key:   "path",
// 			Value: slog.StringValue(path),
// 		},
// 	}
// 	handler := slog.NewJSONHandler(os.Stdout, nil).WithAttrs(loggerAttributes)
// 	logger := slog.New(handler)
// 	return logger
// }
