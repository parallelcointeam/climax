package climax

import (
	"testing"
)

var exampleApp = &KV{
	"app", Tags{
		{"name", "example app"},
		{"brief", "brief text describing app in one sentence"},
		{"usage", "example invocation of the flag"},
		{"version", "v0.0.0a"},
		{"default", ""},
		{"commands",
			Tags{
				{"ctl",
					Tags{
						{"short", ""},
						{"brief", ""},
						{"usage", ""},
						{"help", ""},
						{"group", ""},
						{"handler",
							func() int {
								return 0
							},
						},
						{"vars",
							Tags{
								{"rpcserver",
									Tags{
										{"short", ""},
										{"brief", ""},
										{"usage", ""},
										{"help", ""},
										{"group", ""},
										{"type", "address"},
										{"slot",
											"here goes a pointer to the config variables"},
										{"default",
											"here goes the default value for this field"},
									},
								},
								{"datadir",
									Tags{
										{"short", ""},
										{"brief", ""},
										{"usage", ""},
										{"help", ""},
										{"group", ""},
										{"type", "string"},
										{"slot",
											"here goes a pointer to the config variable"},
										{"default",
											"here goes the default value for this field"},
									},
								},
							},
						},
						{"oneshots",
							Tags{
								{"init",
									Tags{
										{"short", ""},
										{"brief", ""},
										{"usage", ""},
										{"help", ""},
										{"group", ""},
										{"implicit", false},
										{"runafter", false},
										{"terminates", true},
										{"handler",
											func() int {
												return 0
											},
										}},
								},
							},
						},
						{"examples",
							List{
								"example flag invocation1",
							},
						},
					},
				},
			},
		},
	},
}

func TestGetPath(t *testing.T) {
	exampleApp.Get("name")
}

/*
// Below is an example minimal declaration for a Spec containing two Commands

var exampleSpec = Spec{
	Name:    "name",
	Brief:   "Brief",
	Version: "v0.0.0",
	Commands: Commands{
		{
			{ // Commands[0][0]
				Name:  "name",
				Short: "short",
				Brief: "brief",
				Usage: "usage",
				Help:  "help",
				Params: Params{
					{
						Name:    "noderpc",
						Brief:   "short",
						Usage:   "usage",
						Help:    "help",
						Type:    mockAddress,
						Default: "127.0.0.1:11047",
					},
				},
				SubCommands: SubCommands{
					{ // SubCommands[0]
						Name:      "name",
						Short:     "short",
						Brief:     "brief",
						Usage:     "usage",
						Help:      "help",
						Implied:   false,
						Terminate: false,
						RunAfter:  false,
						Handler: func() int {
							// handle subcommand
							return 0
						},
					},
					// SubCommands[1]...
					{},
				},
				Handler: func() int {
					// handle command
					return 0
				},
			},
			{ // Command[0][1]
				Name:        "node",
				Short:       "",
				Brief:       "",
				Usage:       "",
				Help:        "",
				Params:      Params{},
				SubCommands: SubCommands{},
				Handler:     func() int { return 0 },
			},
		},
		{ // Commands[1][0]
			// ...
		},
	},
	Handler: func() int {
		// handle default
		return 0
	},
}

var mockAddress = ToAddress(&mockstring)

var mockstring = "127.0.0.1:11047"
*/
