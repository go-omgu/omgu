# Flags help printer

## Usage

```go
// Usage std flag mod
import "flag"
import "gopkg.in/omgu.v0/flags"

func Example1() {
		// Define args
		var param string
		flag.StringVar(&param, "s", "", "A param")
		flag.Parse()
		
		// Print all flags usage
		flags.PrintAll()
		
		// Print with sections
		section1 := &flags.FlagSection{
				Name: "general",
				Flags: []string{"s"},
		}
		flags.PrintHelp([]*FlagSection{section1})
}

func Example2() {
		// Multistring
		var names MultiString
		flag.Var(&names, "n", "names")
		flag.Parse()
		
		// xx -n 1 -n 2 -n 3
		len(names) // 3, got 1, 2, 3
		
		// Multiword
		// xx -w foo,bar -w ace
		// got foo, bar, ace
		
		// MultiInt
		// xx -i 100,200 -i 300
		// got 100, 200, 300
		
		// MultiIntRange
		// xx -r 1-10,11 -i 13
		// got 1,2,...,11,13
}
```
