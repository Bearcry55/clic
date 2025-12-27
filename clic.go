package clic

import (
	"fmt"
	"os"
)

type CLI struct {
	flags      map[string]Flag
	appName    string
	appDesc    string
	appVersion string
}

type Flag struct {
	name         string
	desc         string
	takesValue   bool
	handler      func()
	valueHandler func(string)
}

// New creates a new CLI instance with default -h and -v flags
func New(appName, appDesc, appVersion string) *CLI {
	c := &CLI{
		flags:      make(map[string]Flag),
		appName:    appName,
		appDesc:    appDesc,
		appVersion: appVersion,
	}
	
	// Auto-register help flag
	c.flags["-h"] = Flag{
		name:       "-h",
		desc:       "Show all available flags",
		takesValue: false,
		handler: func() {
			c.showHelp()
		},
	}
	
	// Auto-register version flag
	c.flags["-v"] = Flag{
		name:       "-v",
		desc:       "Show version",
		takesValue: false,
		handler: func() {
			fmt.Printf("%s v%s\n", c.appName, c.appVersion)
		},
	}
	
	return c
}

// Flag registers a flag without value
func (c *CLI) Flag(name, desc string, handler func()) {
	c.flags[name] = Flag{
		name:       name,
		desc:       desc,
		takesValue: false,
		handler:    handler,
	}
}

// FlagWithValue registers a flag that takes a value
func (c *CLI) FlagWithValue(name, desc string, handler func(string)) {
	c.flags[name] = Flag{
		name:         name,
		desc:         desc,
		takesValue:   true,
		valueHandler: handler,
	}
}

// showHelp prints all registered flags professionally
func (c *CLI) showHelp() {
	fmt.Printf("\n%s\n", c.appName)
	if c.appDesc != "" {
		fmt.Printf("%s\n", c.appDesc)
	}
	fmt.Println("\nUSAGE:")
	fmt.Printf("  %s [flag] [value]\n", c.appName)
	fmt.Println("\nFLAGS:")
	
	for _, flag := range c.flags {
		if flag.takesValue {
			fmt.Printf("  %-15s %s (requires value)\n", flag.name, flag.desc)
		} else {
			fmt.Printf("  %-15s %s\n", flag.name, flag.desc)
		}
	}
	fmt.Println()
}

// Parse checks command line args and executes matching handlers in order
func (c *CLI) Parse() {
	if len(os.Args) < 2 {
		fmt.Printf("No flag provided. Use -h for help\n")
		return
	}
	
	i := 1
	for i < len(os.Args) {
		arg := os.Args[i]
		
		if flag, exists := c.flags[arg]; exists {
			if flag.takesValue {
				if i+1 >= len(os.Args) {
					fmt.Printf("Error: %s requires a value\n", arg)
					return
				}
				flag.valueHandler(os.Args[i+1])
				i += 2 // Skip flag and its value
			} else {
				flag.handler()
				i++ // Skip flag only
			}
		} else {
			fmt.Printf("Unknown flag: %s. Use -h for help\n", arg)
			return
		}
	}
}