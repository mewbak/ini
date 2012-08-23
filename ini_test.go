// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package ini

import (
	"strings"
	"testing"
)

func TestRoundtrip(t *testing.T) {
	var err error

	a := New()
	s := a.Section("")
	s["randomkey"] = "abc def"

	s = a.Section("graphics")
	s.Set("width", 320)
	s.Set("height", 240)
	s.Set("depth", 32)

	s = a.Section("sound")
	s.Set("volume-master", 100)
	s.Set("volume-left", 0.6)
	s.Set("volume-right", 0.65)

	s = a.Section("misc")
	s["url"] = "http://www.server.com/page?var=value"

	if err = a.Save("testdata/test.ini"); err != nil {
		t.Fatal(err)
	}

	b := New()
	if err = b.Load("testdata/test.ini"); err != nil {
		t.Fatal(err)
	}

	if len(a.Sections) != len(b.Sections) {
		t.Fatalf("len(a.Sections) != len(b.Sections)")
	}

	for name, sa := range a.Sections {
		sb := b.Section(name)

		if len(sa) != len(sb) {
			t.Fatalf("%q: len(sa) != len(sb)", name)
		}

		for key := range sa {
			va := sa.S(key, "")
			vb := sb.S(key, "")

			if vb == "" {
				t.Fatalf("%q: Missing key %q", name, key)
			}

			if !strings.EqualFold(va, vb) {
				t.Fatalf("%q:%q: Value mismatch: %q - %q", name, key, va, vb)
			}
		}
	}
}
