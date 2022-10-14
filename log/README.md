# log

Features:

- Grepable JSON
- JSON
- CSV

## Usage

### Logger

```go
import "xxx/log"

// Create a Logger
// Log level: Normal > Verbose > Debug
l := log.New(os.Stdout, log.Verbose)

l.Log("aaa", "bbb") // this will displayed
l.Verbose("aaa", "bbb") // this will displayed
l.Debug("aaa", "bbb") // this will NOT displayed

// Dump a struct
l.Dump(xxx) // Dump is debug level
```

### JSON

```go
import "xxx/log"

log.PrintJSON(xxx)

// pretty
log.PrintJSONPretty(xxx)
```

### CSV

```go
import "xxx/log"

// Example output struct
type Result struct {
	A string `csv:"a,0"`
	B string `csv:"b,1"`
	C string `csv:"c,2"`
}

logger := log.NewCSV[Result](os.Stdout)
// or start with a record
r1 := &Result{"", "", ""}
logger := log.NewCSVLike[Result](os.Stdout, r1)

// append new record
logger.Append(&Result{"", "", ""})
logger.AppendMany([]Result{...})
```
