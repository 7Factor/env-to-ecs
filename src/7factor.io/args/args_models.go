package args

// Place your Args models here and call them in args.go
// Include your cli short and long flags, default value for arg, and a usage string to show up in --help calls.
var opt struct {
	Infile    string   `short:"i" long:"infile" description:"The infile to parse." required:"true"`
	Outfile   string   `short:"o" long:"outfile" description:"The outfile to write to." default:"stdout" required:"false"`
	Variables []string `short:"v" long:"variable" description:"Optional variable to pass ie A=B." required:"false"`
}
