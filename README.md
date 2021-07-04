# log

### Example

```go
  if err := log.Init("example.log", log.CommonLog, false); err != nil {
    fmt.Printf("Failed init log: %v", err)
    os.Exit(1)
  }
  
  log.Info.Println("Info text")
  log.Warning.Println("Warning text")
  log.Error.Print("Error text")
  log.Fatal.Println("Fatal text")
  log.Trace.Println("Trace text")
  log.Debug.Println("Debug text")
```

### Output

```
INFO:    2021/07/04 19:28:55 main.go:14: Info text
WARNING: 2021/07/04 19:28:55 main.go:15: Warning text
ERROR:   2021/07/04 19:28:55 main.go:16: Error text
FATAL:   2021/07/04 19:28:55 main.go:17: Fatal text
```

### Level logging

There are three levels of logging.

- **CommonLog** - Normal logging mode. uses info, warning, error, fatal logs.
- **DebugLog** - Usually only enabled when debugging. Shows all the logs that CommonLog mode and more Debug.
- **TraceLog** - Designates finer-grained informational events than the Debug. Shows all the above logs and more Trace.
