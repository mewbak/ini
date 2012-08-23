// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

/*
This package allows us to read and write .ini configuration files.

Standard features include:

    * Named sections.
    * Global key/value pairs (those not associated with a section).
      Every file comes with a default, global section in which all of these
      are stored. It can be accessed using the empty section name ("").
    * Builtin methods for easy type conversions of values.

Non-standard features include:

    * Appending multiple values to the same key using the `<` operator
      as opposed to `=`.

Example of writing an ini file:

	file := ini.New()
	gfx := file.Section("graphics")

	// Directly assigning a value.
	gfx["title"] = "My window title"

	// Indirectly setting a value.
	gfx.Set("width", width)
	gfx.Set("height", height)
	gfx.Set("fullscreen", fullscreen)

	// Write contents to a file.
	err := file.Write("myfile.ini")


Example of reading an ini file:

	file := ini.New()
	err := file.Read("myfile.ini")

	if err != nil {
		return
	}

	// Read a section directly.
	gfx, ok := file.Sections["graphics"]
	if !ok {
		return
	}

	// Or use the Section method. This creates a new section
	// implicitely if it does not exist.
	gfx = file.Section("graphics")

	// Read a key/value pair directly
	title := gfx["title"].(string)

	// Or use the type conversion methods.
	width := gfx.I32("width")
	height := gfx.I32("height")
	fullscreen := gfx.B("fullscreen")


Enumerate all sections and key/value pairs:

	file := ini.New()

	// ... Load up data ...

	for name, section := range file.Sections {
		if len(name) == 0 {
			fmt.Println("Section <global>")
		} else {
			fmt.Printf("Section %q\n", name)
		}

		for key, value := range section {
			fmt.Printf(" %s = %q\n", key, value)
		}

		fmt.Println()
	}

Using the append operator:

	[keys]
	auth < xxxxxxxx
	auth < yyyyyyyy
	auth < zzzzzzzz

This creates a list of string values for the `auth` key.
This is useful if you want to append new values later on without having
to modify the parsing code for new keys. In the example above, we use it
to add new authentication keys for key rotation purposes.

	keys := file.Section("keys")
	auth := keys.List("auth") // auth == []string{...}

*/
package ini
