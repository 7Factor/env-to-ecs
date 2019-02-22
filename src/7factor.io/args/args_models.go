package args

// Place your Args models here and call them in args.go
// Include your cli short and long flags, default value for arg, and a usage string to show up in --help calls.

var inFileArgs = map[string]string{
	"shortFlag":    "i",
	"longFlag":     "infile",
	"defaultValue": "",
	"usage":        "The infile to parse",
}

var outFileArgs = map[string]string{
	"shortFlag":    "o",
	"longFlag":     "outfile",
	"defaultValue": "stdout",
	"usage":        "The outfile to parse",
}
