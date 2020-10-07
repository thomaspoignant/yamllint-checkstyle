package extractor

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/thomaspoignant/yamllint-checkstyle/checkstyle"
)
var ExtractRegex = regexp.MustCompile(`^([^:]*):([0-9]*):([0-9]*): \[(.*)\] (.*) \((.*)\)$`)
func YamlLintExtractor(inputLine string) (string, checkstyle.Error, error){
	extractedRegex := ExtractRegex.FindAllStringSubmatch(inputLine, -1)
	if len(extractedRegex) >0 {
		extract := extractedRegex[0]
		line, _ := strconv.Atoi(extract[2]) // We can ignore the errors because we extract only numbers
		column, _ := strconv.Atoi(extract[3]) // We can ignore the errors because we extract only numbers
		return extract[1], checkstyle.Error{
			Line: line,
			Column: column,
			Severity: extract[4],
			Message: extract[5],
			Source: extract[6],
		}, nil
	}
	return "", checkstyle.Error{}, fmt.Errorf("impossible to extract this line: %s",inputLine)
}