# clic

A minimal, clean CLI library for Go. Build command-line tools with simple, readable code.

## Features

- ✅ Simple flags (`-t`, `-h`)
- ✅ Flags with values (`-n John`, `-p 8080`)
- ✅ Multiple flags in one command
- ✅ Auto-generated help (`-h`)
- ✅ Auto-generated version (`-v`)
- ✅ Professional output formatting
- ✅ Zero dependencies
- ✅ No bloat

## Installation

```bash
go get github.com/Bearcry55/clic
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/Bearcry55/clic"
)

func main() {
    cli := clic.New("myapp", "My awesome CLI tool", "1.0.0")
    
    // Simple flag
    cli.Flag("-t", "Run tests", func() {
        fmt.Println("Running tests...")
    })
    
    // Flag with value
    cli.FlagWithValue("-n", "Set your name", func(name string) {
        fmt.Printf("Hello, %s!\n", name)
    })
    
    cli.Parse()
}
```

## Usage Examples

**Show help:**
```bash
$ myapp -h

myapp
My awesome CLI tool

USAGE:
  myapp [flag] [value]

FLAGS:
  -h              Show all available flags
  -v              Show version
  -t              Run tests
  -n              Set your name (requires value)
```

**Show version:**
```bash
$ myapp -v
myapp v1.0.0
```

**Use flags:**
```bash
$ myapp -t
Running tests...

$ myapp -n John
Hello, John!
```

**Multiple flags:**
```bash
$ myapp -t -n John
Running tests...
Hello, John!
```

## API

### `New(appName, appDesc, appVersion string) *CLI`
Creates a new CLI instance with auto-generated `-h` and `-v` flags.

### `Flag(name, desc string, handler func())`
Registers a flag that doesn't take a value.

### `FlagWithValue(name, desc string, handler func(string))`
Registers a flag that takes a value.

### `Parse()`
Parses command-line arguments and executes matching handlers.

## Why clic?

Most CLI libraries are over-engineered. clic focuses on:
- **Simplicity** - 10 lines to build a working CLI
- **Clarity** - Clean, readable code
- **Minimalism** - No unnecessary features

Perfect for small tools, scripts, and utilities where you want something that just works.

## License

MIT

## Contributing

Issues and PRs welcome at [github.com/Bearcry55/clic](https://github.com/Bearcry55/clic)
