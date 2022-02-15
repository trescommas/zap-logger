# Zap Logger

Easy to use logger that prints to console and rotates log files


Built on top of [zap](https://github.com/uber-go/zap) + [lumberjack](https://github.com/natefinch/lumberjack)


### Usage

```
import "github.com/trescommas/zap-logger/logger"

log := logger.NewLogger("./logs/test.log", 10, 2, 1, false)
log.Debug("works!")
```
