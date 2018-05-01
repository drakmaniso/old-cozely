/*
Package data provides a parser for the ".czl" configuration format


WORK IN PROGRESS

This is currently only a proof of concept, with no implementation.

Comments

Comments start with two pipe characters "||", and end with newline:

  || This whole line is a comment
  (1, 2, 3) || But comments can also occur at end of line

An isolated pipe character is considered ambiguous:

  Red | Blue || Ambiguous! (interpreted as "Red | Blue"; see quoted strings)

Strings

A String is any character sequence not interpreted by the parser, i.e. a piece
of data considered atomic.

  First
  Second string
  This line is a single string!
  1001
  3.14
  12/03/2001
  1280 x 720

Unquoted strings cannot contain newlines or special characters. Blank spaces are
allowed, but are stripped at the begininng and the end of the value.

The 16 special characters:

  | , : = " @ & $
  ( ) [ ] { } < >

Note that backslash is not a special character! The following is allowed:

  C:\Program Files\Cozely

A quoted string is a character sequence between double quotes; any character is
allowed except newline and the double quote:

  "This (quoted) string: contains several special characters."

To insert a double quote character, use two in succession:

  "Here is a ""quote"" inside a string." || Here is a "quote" inside a string.

To allow newlines, put a newline immediatelu after the opening double quote, and
another immediately before the closing one:

  "This is
  wrong..." || Ambiguous! (newline in quoted value)

  "
  This, however, is a
  single string spanning on *two* lines!
  "

Lists

A list is a sequence of coma-separated items surrounded by brackets. Each item
is either a string or a list. In other words, list can be nested:

  (first item, second item, third item)
  (first item is a string, (second, item, is, a, list))

Items can be separated by a coma, a newline, or both. In other words, comas at
end-of-line are optional:

  (first, second,
  third)

  (
    first,
    second,
    third
  )

  (
    first
    second
    third
  )

  (
    first,
    second,
    third,
  )

It is however an error to use two comas in a row. To correctly insert an empty
item, an empty quoted string "" should be used:

  (first, , third) || Ambiguous! (interpreted as three items)
  (first, "", third) || Second item is empty

It is also an error to use a coma just before a bracket, except if there is a
newline between them:

  (first, second, ) || Ambiguous! (interpreted as *two* items)

  (
    first,
    second, || Correct (two items)
  )

On the other hand, multiple newlines between items are allowed:

  (
    first


    second || Correct (two items)
  )

Four different kind of brackets can be used, and are all equivalent: round
brackets "()", square brackets "[]", curly brackets "{}", and angle brackets
"<>". The following lists are all the same:

  (first, second, third)
  [first, second, third]
  {first, second, third}
  <first, second, third>

The different brackets are especially useful to write nested lists:

  (first, [s, e, c, o, n, d], third)

  (
    first
    [s, e, c, o, n, d]
    [<t, h>, i, r ,d]
  )

In addition to simple items, lists can also contain key-value pairs. A pair is
distinguished from a normal item by the presence of a colon ":" or an equal sign
"=" between the key and the value:

  (first item, second item, an: option, another: option)

Key-value pairs are not ordered, but they must occur after the last item:

  (first, second, lang=english, count=2)
  (first, second, count=2, lang=english) // Same list as above
  (lang=english, first, count=2, second) // Ambiguous! (interpreted as above)

It is an error to omit the key or the value of a pair. To specify an empty key
or an empty value, use a quoted empty string:

  (1, 2, lang=) // Ambiguous (interpreted as an empty value)
  (1, 2, lang="")
  (1, 2, =arabic) // Ambigous (interpreted as an empty key)
  (1, 2, ""=arabic)

In a key-value pair, the value can be either a string or a list, but the key
must be a string:

  (first, second, lang=english)
  (first, second, lang=[british, american])
  (first, second, [lang, language]=english) || Ambiguous! (the pair is ignored)

In a given list, all key-value pairs must have unique keys:

  (first, second, lang=british, lang=american) || Ambiguous! (second pair overrides the first)
  (first, first, first) || Correct: items don't need to be unique

*/
package data
