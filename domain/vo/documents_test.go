package vo

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type docTest struct {
	doc        string
	shouldPass bool
}

func (tst docTest) run(t *testing.T) {
	doc, err := ParseDocument(tst.doc)
	if (err != nil) && (tst.shouldPass) {
		t.Fatalf("this should work with doc: %v", doc)
	}

	if (err == nil) && !tst.shouldPass {
		t.Fatalf("this should not work with doc: %v", doc)
	}

	if err != nil {
		require.Equal(t, doc, Document{})
		require.Equal(t, err, ErrInvalidDocument)
	}
}

var docTests = []docTest{
	{"10579592944", true},
	{"105.795.929-44", false},
	{"abcdefghijk", false},
	{"12.345.678/0001-95", false},
	{"12345678000195", true},
}

func TestParseDocument(t *testing.T) {
	for _, tst := range docTests {
		t.Run(tst.doc, tst.run)
	}
}
