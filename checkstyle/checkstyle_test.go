package checkstyle

import (
	"testing"
)

func TestBuildCheckStyle(t *testing.T) {
	checkStyle := New()

	checkStyle.AddError("path/to/file", Error{
		Line:10,
		Column: 5,
		Severity: "error",
		Message: "msg1",
		Source: "test",
	})

	checkStyle.AddError("path/to/file", Error{
		Line:11,
		Column: 58,
		Severity: "debug",
		Message: "msg2",
		Source: "test",
	})

	// Check the output
	want:= `<checkstyle version="1.0.0"><file name="path/to/file"><error line="10" column="5" severity="error" message="msg1" source="test"></error><error line="11" column="58" severity="debug" message="msg2" source="test"></error></file></checkstyle>`
	if checkStyle.String() != want {
		t.Errorf("Wrong output for String()\n\twant: %s\n\tgot: %s",want,checkStyle.String())
	}
}
