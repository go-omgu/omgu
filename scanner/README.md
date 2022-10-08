# Scanner

Line scanner auto filter comments and empty line.

## Usage

```go
// Scan from STDIN
s := scanner.NewScanner(os.Stdin)

// Add new filter prefix
// Default has '#' and '//'
s.SkipPrefix = append(s.SkipPrefix, []byte("--"))

// Scan a byte slice
buf, err := s.Scan()

// Scan a string
str, err := s.ScanString()

// Use with for
for s, err := s.ScanString(); err == nil; str, err = s.ScanString() {
	// Do something...
}
```
