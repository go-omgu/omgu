# Progress

Simple progress displayer.

## Usage

```go
// New progress with max value
p := NewProgress(100)

// Run with a title
p.Run("Running...")

// Dynamic update title and max
p.SetTitle("Look...")
p.SetMax(120)

// Increse progress
p.Incr(1)

// Finish progress
p.Finish()
```
