package extractor

import (
	"reflect"
	"testing"

	"github.com/thomaspoignant/yamllint-checkstyle/checkstyle"
)

func TestYamlLintExtractor(t *testing.T) {
	type want struct {
		filename string
		error checkstyle.Error
	}
	tests := []struct {
		name    string
		inputLine    string
		want    want
		wantErr bool
	}{
		{
			name: "yamllint valid line",
			inputLine:"test.yaml:1:8: [error] no new line character at the end of file (new-line-at-end-of-file)",
			wantErr: false,
			want: want{
				filename: "test.yaml",
				error: checkstyle.Error{
					Line: 1,
					Column: 8,
					Severity: "error",
					Message: "no new line character at the end of file",
					Source: "new-line-at-end-of-file",
				},
			},
		},
		{
			name: "yamllint invalid line",
			inputLine:"test.yaml:t:8: [error] no new line character at the end of file (new-line-at-end-of-file)",
			wantErr: true,
			want: want{
				filename: "",
				error: checkstyle.Error{},
			},
		},
		{
			name: "yamllint invalid column",
			inputLine:"test.yaml:1:t: [error] no new line character at the end of file (new-line-at-end-of-file)",
			wantErr: true,
			want: want{
				filename: "",
				error: checkstyle.Error{},
			},
		},
		{
			name: "yamllint invalid format",
			inputLine:"[error] test:1,2 no new line character at the end of file (new-line-at-end-of-file)",
			wantErr: true,
			want: want{
				filename: "",
				error: checkstyle.Error{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filename, checkstyleErr, err := YamlLintExtractor(tt.inputLine)
			if (err != nil) != tt.wantErr {
				t.Errorf("YamlLintExtractor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if filename != tt.want.filename {
				t.Errorf("YamlLintExtractor() filename = %v, want %v", filename, tt.want.filename)
			}
			if !reflect.DeepEqual(checkstyleErr, tt.want.error) {
				t.Errorf("YamlLintExtractor() checkstyleErr = %v, want %v", checkstyleErr, tt.want.error)
			}
		})
	}
}