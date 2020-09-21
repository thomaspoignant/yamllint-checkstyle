package main

import (
	"fmt"
	"github.com/bitfield/script"
	"strings"

	"github.com/thomaspoignant/yamllint-checkstyle/checkstyle"
	"github.com/thomaspoignant/yamllint-checkstyle/extractor"
)

func main() {
	report := checkstyle.New()
	script.Stdin().MatchRegexp(extractor.ExtractRegex).EachLine(func(inputLine string, out *strings.Builder) {
		filename, yamlError, err := extractor.YamlLintExtractor(inputLine)
		if err!=nil{
			panic(err)
		}
		report.AddError(filename, yamlError)
	})
	// Output XML
	fmt.Print(report)
}