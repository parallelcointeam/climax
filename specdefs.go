package climax

import (
	"errors"
	"time"
)

// Address is an alias used to filter for string representation of a network address
type Address *string

// Bool is an alias used to filter for string representation of a boolean
type Bool *bool

// Duration is an alias used to filter for string representation of a duration (time.Duration is nanoseconds)
type Duration *time.Duration

// Float is an alias used to filter for string representation of a float64 as used in json
type Float *float64

// Handler is a
type Handler func() int

// Int is an alias used to filter for string representation of a
type Int *int

// KV is a named container that can contain other data types and including pointers to other KV's, allowing the creation of sparse trees.
/*
	Arrays of this type are used to declare collections of key/values which can store other KV's under a tag, or lists of strings for collections of string values such as examples

	The reason for using arrays is that the invocations of oneshot flags (non-variables in climax's base structure) are order-sensitive since they can terminate execution or optionally be triggered at shutdown and other ways of encoding ordering are cumbersome to declare.
*/
type KV struct {
	Key  string
	Data interface{}
}

// List is a simple list of strings, specifically used to handle example invocation lists, and arrays of List can be used for Topics since they have only 3 fields, they can be interpreted to be in this fixed order without requiring tags to denote the field's type.
type List []string

// PathVariant is a go-equivalent of the 'variant' type that contains one of multiple other types inside it
type PathVariant struct {
	KV   *KV
	Tags *Tags
}

// String is an alias used to filter for string representation of a
type String *string

// Tags is an array of KVs
type Tags []KV

// URL is an alias used to filter for string representation of a
type URL *string

var Parsers = map[string]func(string, interface{}) bool{
	"bool":     ParseBool,
	"int":      ParseInt,
	"float":    ParseFloat,
	"duration": ParseDuration,
	"string":   ParseString,
	"address":  ParseAddress,
	"url":      ParseURL,
}

var Types map[string]interface{}

// NewVar is a quick way to create a new, typed variable, suitable for the Type field of Spec.Commands.Params that defines how the value specified in a flag will be parsed.
// The recommended best way for doing this is to put the address of a configuration struct field that will be filled by a command line option
func NewVar(typeOf string) (in interface{}) {
	switch typeOf {
	case "bool":
		return new(Bool)
	case "int":
		return new(Int)
	case "float":
		return new(Float)
	case "duration":
		return new(Duration)
	case "string":
		return new(String)
	case "address":
		return new(Address)
	case "url":
		return new(URL)
	}
	return errors.New("'" + typeOf + "' is not a valid climax field type")
}

/*
// Commands are an array of arrays that are implicitly grouped by their membership in the second row of the slice. If a group name is not given, the default label is simply the index of the first level of the slice, so as not to be overly distracting in its non-informativeness, while giving the desired structure.
type Commands [][]struct {
	// The name field of the first element of the slice of structs is interpreted to mean the subsequent elements are grouped with this name
	Name string
	// Short is an arbitrary shorter string that also triggers the same command
	Short string
	// Brief description, preferably under 80 characters, that will be printed in the default flag/definition help text.
	Brief string
	// Usage is an example, typical use of the flag
	Usage string
	// Help is a longer text that provides more detail about the command.
	Help string
	Params
	SubCommands
	Handler func() int
}

// Params are the parameters for the command.
// These are called SubCommands but they are the 'NonVariable' type in Context - they appear here as a separate array field in Spec but they are defined in the output Application and Command structs the Spec parser processes.
// These are for one-off processes that may cause the application to exit immediately afterwards, performed before normal startup, or scheduled to be run on cleanup. These are also automatically, implicitly grouped under the default group name of "oneshots" as they are not configuration values, but rather, triggers that execute prior to the main config handler, in order of specification
type Params []struct {
	// Same basic structure applies to parameters as the same named fields above. These go into the 'Variable' map in Context, and can have a defined default value and arbitrary sanitiser function that checks any aspect of the variable for sanity and will trigger abort and error message.
	Name  string
	Brief string
	Usage string
	Help  string
	// Type must be a pointer to one of the types
	Type    interface{}
	Default interface{}
}

// Spec is an extension to the climax package that provides a structured definition syntax that specifies all of the textual information, uses a type switch
type Spec struct {
	// The name of the application, usually should be also the name of the executable (ie the folder name containing the main, but this is not mandatory)
	Name string
	// Brief one sentence description
	Brief string
	// It is recommended to use semver conventions in this string.
	Version string
	Commands
	// Handler is run if climax finds its name is found
	Handler func() int
}

type SubCommands []struct {
	Name  string
	Short string
	Brief string
	Usage string
	Help  string
	// If implied is true, the presence of this flag indicates the value is false
	Implied bool
	// If this value is set to true, the program will exit at completion of the execution of the handler, if the flag appears, the value is false and the subcommand does not execute.
	Terminate bool
	// If this handler is to be chained to the main defer of the handler, this flag is set to true
	RunAfter bool
	// Handler returns a 0 for success and nonzero for failure, and is intended to be passed through to os.Exit() to terminate and return a status to the shell.
	Handler func() int
}


*/
