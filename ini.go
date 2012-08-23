// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package ini

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

// File allows reading/writing of ini configuration data.
type File struct {
	Sections map[string]Section
}

// New creates a new, empty file.
func New() *File {
	f := new(File)
	f.Sections = make(map[string]Section)
	return f
}

// Section returns a section for the given name.
// If none exists, one is created.
//
// Use an empty name ("") to access the global section.
func (i *File) Section(name string) Section {
	if s, ok := i.Sections[name]; ok {
		return s
	}

	s := make(Section)
	i.Sections[name] = s
	return s
}

// Load loads ini data from the given file.
func (i *File) Load(file string) (err error) {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		return
	}

	var r Reader
	var key string

	section := i.Section("")

	r.Read(data, func(tt TokenType, value []byte) {
		str := string(bytes.TrimSpace(value))

		switch tt {
		case TokSection:
			if len(key) > 0 {
				key = ""
			}

			section = i.Section(str)

		case TokKey:
			key = str
			section[key] = ""

		case TokKeyList:
			key = str
			if _, ok := section[key]; !ok {
				section[key] = []string{}
			}

		case TokValue:
			if len(key) > 0 {
				if l, ok := section[key].([]string); ok {
					l = append(l, str)
					section[key] = l
				} else {
					section[key] = str
				}

				key = ""
			}

		}
	})

	return
}

// Save saves ini data to the given file.
func (i *File) Save(file string) (err error) {
	var w bytes.Buffer

	// Write Global section first.
	writeSection(&w, "", i.Section(""))

	// Write remaining sections.
	for name, section := range i.Sections {
		if len(name) == 0 {
			continue // Already done global section.
		}

		writeSection(&w, name, section)
	}

	return ioutil.WriteFile(file, w.Bytes(), 0600)
}

// WriteSection writes the given section to the supplied writer.
func writeSection(w io.Writer, name string, s Section) {
	if len(name) > 0 {
		fmt.Fprintf(w, "[%s]\n", name)
	}

	for key, value := range s {
		if list, ok := value.([]string); ok {
			for _, item := range list {
				fmt.Fprintf(w, "%s < %s\n", key, item)
			}
			continue
		}

		fmt.Fprintf(w, "%s = %s\n", key, value)
	}

	fmt.Fprintln(w)
}
