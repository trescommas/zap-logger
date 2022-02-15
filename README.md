# Zap Logger

Simple [zap](https://github.com/uber-go/zap) + [lumberjack](https://github.com/natefinch/lumberjack)

## Description

Easy to use logger that prints to console and rotates log files

### Usage

```
import "github.com/trescommas/zap-logger"

log := logger.NewLogger("./logs/test.log", 10, 2, 1, false)
log.Debug("works!")
```
