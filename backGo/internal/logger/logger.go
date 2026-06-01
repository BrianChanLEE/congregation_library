package logger

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

// GinWriter는 GIN 로그에 색상을 적용합니다.
type GinWriter struct{}

func (w *GinWriter) Write(p []byte) (n int, err error) {
	msg := string(p)
	// GIN-debug 노란색, WARNING은 붉은색
	if strings.Contains(msg, "[GIN-debug]") {
		msg = "\033[33m" + msg + "\033[0m"
	} else if strings.Contains(msg, "[WARNING]") {
		msg = "\033[31m" + msg + "\033[0m"
	}
	return os.Stdout.Write([]byte(msg))
}

// Log는 프로젝트 전역에서 사용할 구조화된 로거입니다.
var Log = slog.Default()

// GinLog는 GIN 전용 단순 로거입니다.
var GinLog = slog.Default()

// Init은 로거를 초기화합니다.
func Init() {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// 시간 형식 간소화
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(a.Value.Time().Format("15:04:05"))
			}
			// 소스 정보 파일명만 표시
			if a.Key == slog.SourceKey {
				source := a.Value.Any().(*slog.Source)
				parts := strings.Split(source.File, "/")
				if len(parts) > 0 {
					source.File = parts[len(parts)-1]
				}
				a.Value = slog.StringValue(fmt.Sprintf("%s:%d", source.File, source.Line))
			}
			// 로그 레벨 패딩
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				var color string
				var levelStr string
				switch level {
				case slog.LevelError:
					color = "\033[31m"
					levelStr = "ERROR"
				case slog.LevelWarn:
					color = "\033[33m"
					levelStr = "WARN "
				case slog.LevelInfo:
					color = "\033[32m"
					levelStr = "INFO "
				}
				a.Value = slog.StringValue(fmt.Sprintf("%s%s\033[0m", color, levelStr))
			}
			return a
		},
	}

	handler := slog.NewTextHandler(os.Stdout, opts)
	Log = slog.New(handler)

	// GIN 전용 로거 (포맷팅 없이 메시지만 출력)
	GinLog = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey || a.Key == slog.LevelKey || a.Key == slog.SourceKey {
				return slog.Attr{} // 필드 제거
			}
			return a
		},
	}))
}
