package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Reference: https://blog.rapid7.com/2016/08/04/build-a-simple-cli-tool-with-golang/

// Manual tests
// $ stringparse list --text "Hackers beware." --metric words --unique
// textPtr: Hackers beware., metricPtr: words, uniquePtr: true
// $ stringparse count --text "She sells sea shells by the sea shore" --metric substring --substringList se,sh,ea,he
// textPtr: She sells sea shells by the sea shore, metricPtr: substring, substringPtr: , substringListPtr: [se sh ea he], uniquePtr: false
// $ stringparse count --text "Komand Security"
// textPtr: Komand Security, metricPtr: chars, substringPtr: , substringListPtr: [], uniquePtr: false

// In the context of its flag package, all strings beginning with dashes (single or double) are considered flags,
//	and all strings without dashes that are not user input are considered arguments. All strings after the first argument are considered arguments.

// Flag syntax:
// 	-example -example=text -example text // Non-boolean flags only
// 	- and -- may be used for flags

// Choice Flags
//   A choice flag must have user input from within a certain set. To implement a choice flag, create a set of values,
//   and check that the user input is in the set. If input is not in the set, notify the user accordingly and exit.

// Help Text
// It’s important for your CLI tool to have a message directing users on how to run it.
// It’s good practice to include a --help or -h flag with your command, which will print usage instructions.
// Go will create this usage message for you, by combining the the help text fields of all your flags.
// The message can be printed explicitly using flag.PrintDefaults(). Usage, help, and man pages commonly use a small syntax to describe the valid command form:
//   angle brackets for required parameters: ping <hostname>
//   square brackets for optional parameters: mkdir [-p] <dirname>
//   ellipses for repeated items: cp <source1> [source2...] <dest>
//   vertical bars for choice of items: netstat {-t|-u}

// Create a new type for a list of Strings
type stringList []string

// Implemenet the flag.Value interface
func (s *stringList) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *stringList) Set(value string) error {
	*s = strings.Split(value, ",")
	return nil
}

func main() {
	// Subcommands
	countCommand := flag.NewFlagSet("count", flag.ExitOnError)
	listCommand := flag.NewFlagSet("list", flag.ExitOnError)

	// Count subcommand flag pointers
	// Adding a new choice for --metric of 'substring' and a new --substring flag
	countTextPtr := countCommand.String("text", "", "Text to parse. (Required)")
	countMetricPtr := countCommand.String("metric", "chars", "Metric {chars|words|lines|substring}. (Required)")
	countSubstringPtr := countCommand.String("substring", "", "The substring to be counted. Required for --metric=substring")
	countUniquePtr := countCommand.Bool("unique", false, "Measure unique values of a metric.")

	// Use flag.Var to create a falg of our new FlagType
	// Default value is the current value at countStringListPtr (currently a nil value)
	var countStringList stringList
	countCommand.Var(&countStringList, "substringList", "A comma separated list of substrings to be counted")

	// List subcommand flag pointers
	listTextPtr := listCommand.String("text", "", "Text to parse. (Required)")
	listMetricPtr := listCommand.String("metric", "chars", "Metric <chars|words|lines>. (Required)")
	listUniquePtr := listCommand.Bool("unique", false, "Measure unique values of a metric.")

	// Verify that a subcommand has been provided
	// os.Arg[0] is the main command
	// os.Arg[1] will be the subcommand
	if len(os.Args) < 2 {
		fmt.Println("List or count subcommand is required!")
		os.Exit(1)
	}

	// Switch on the subcommand
	// Parse the flags for appropriate FlagSet
	// FlagSet.Parse() requires a set of arguments to parse as input
	// os.Args[2:] will be all arguments starting after the subcommand at os.Args[1]
	switch os.Args[1] {
	case "list":
		listCommand.Parse(os.Args[2:])
	case "count":
		countCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Check which subcommand was Parsed using the FlagSet.Parsed() function.
	// FlagSet.Parse() will evaluate to false if no flags were parsed (i.e. the user did not provide any flags)
	if listCommand.Parsed() {
		// Required Flags
		if *listTextPtr == "" {
			listCommand.PrintDefaults()
			os.Exit(1)
		}

		// Choice flag
		metricChoices := map[string]bool{"chars": true, "words": true, "lines": true}
		if _, validChoice := metricChoices[*listMetricPtr]; !validChoice {
			listCommand.PrintDefaults()
			os.Exit(1)
		}

		// Print
		fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *listTextPtr, *listMetricPtr, *listUniquePtr)
	}

	if countCommand.Parsed() {
		// Required Flags
		if *countTextPtr == "" {
			countCommand.PrintDefaults()
			os.Exit(1)
		}

		// If the metric flag is substring, the substring flag is required
		if *countMetricPtr == "substring" && *countSubstringPtr == "" && (&countStringList).String() == "[]" {
			countCommand.PrintDefaults()
			os.Exit(1)
		}

		// If the metric flag is not substring, the substring flag must not be used
		if *countMetricPtr != "substring" && (*countSubstringPtr != "" || (&countStringList).String() != "[]") {
			fmt.Println("--substring and --substringList may only be used with --metric=substring.")
			countCommand.PrintDefaults()
			os.Exit(1)
		}

		// Choice flag
		metricChoices := map[string]bool{"chars": true, "words": true, "lines": true, "substring": true}
		if _, validChoice := metricChoices[*countMetricPtr]; !validChoice {
			listCommand.PrintDefaults()
			os.Exit(1)
		}

		// Print
		fmt.Printf("textPtr: %s, metricPtr: %s, substringPtr: %v, substringListPtr: %v, uniquePtr: %t\n",
			*countTextPtr, *countMetricPtr, *countSubstringPtr, (&countStringList).String(), *countUniquePtr)
	}

	// textPtr := flag.String("text", "", "Text to parse.")
	// metricPtr := flag.String("metric", "chars", "Metric {chars|words|lines};.")
	// uniquePtr := flag.Bool("unique", false, "Measure unique values of a metric.")
	// flag.Parse()

	// // Making the flag required
	// if *textPtr == "" {
	// 	flag.PrintDefaults()
	// 	os.Exit(1)
	// }

	// fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *textPtr, *metricPtr, *uniquePtr)
}
