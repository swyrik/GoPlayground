package main

import (
	"fmt"
	s "strings" // you can have alias for an import
	"unicode"
)

var p = fmt.Println

func stringfuncstest() {
	p("Contains :", s.Contains("test", "es"))
	p("Count :", s.Count("test", "t"))
	p("HasPrefix :", s.HasPrefix("test", "te"))
	p("HasSuffix :", s.HasSuffix("test", "st"))
	// index starts from 0
	p("Index :", s.Index("test", "e"))
	p("LastIndex :", s.LastIndex("test", "t"))
	p("Join: ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat: ", s.Repeat("a", 5))
	p("Replace: ", s.Replace("foo", "o", "0", -1))
	p("Replace: ", s.Replace("foo", "o", "0", 2))
	p("Replace: ", s.Replace("foo", "", "0", -1))
	p("Split: ", s.Split("a-b-c-d-e", "-"))
	p("ToLower: ", s.ToLower("Test"))
	p("ToUpper: ", s.ToUpper("Test"))

	// Trim
	p("Trim: ", s.Trim(" GO Lang ", " "))
	p("Trim: ", s.Trim("!GO Lang!", "!"))
	p("Trim: ", s.Trim("$GO Lang$$$$", "$"))
	p("Trim: ", s.Trim("!$GO Lang$!$$", "$!"))

	p("Title: ", s.Title("hello world"))   // Hello World converts the first letter of each word to upper case
	p("Title: ", s.ToTitle("hello world")) // HELLO WORLD converts all the letters to upper case
	p("Title: ", s.ToTitleSpecial(unicode.SpecialCase{}, "Hello world prgrm"))

	p("Fields:", s.Fields("helloworld")) // returns array of string

}
