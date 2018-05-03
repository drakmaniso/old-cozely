/*
Package data provides a parser for the ".czl" configuration format


WORK IN PROGRESS

This is currently only a proof of concept, with no implementation.

Comments

Comments start with two semicolons ";;", and end with newline:

  ;; This whole line is a comment
  (1, 2, 3) ;; Comment on the same line as meaningful data

An isolated semicolon is considered ambiguous:

  this is ; Ambiguous! (but still interpreted as a comment)

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

The 13 special characters:

  ; , : = "
  ( ) [ ] { } < >

Note that backslash is not a special character! The following is allowed:

  C:\Program Files\Cozely

A quoted string is a character sequence between double quotes; any character is
allowed except newline and the double quote:

  "This (quoted) string: contains several special characters."

To insert a double quote character, use two in succession:

  "Here is a ""quote"" inside a string." ; Here is a "quote" inside a string.

To allow newlines, put a newline immediately after the opening double quote, and
another immediately before the closing one:

  "This is
  wrong..." ;; Ambiguous! (newline in quoted value)

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

  (first, , third)   ;; Ambiguous! (interpreted as three items)
  (first, "", third) ;; Second item is empty

It is also an error to use a coma just before a bracket, except if there is a
newline between them:

  (first, second, ) ;; Ambiguous! (interpreted as *two* items)

  (
    first,
    second, ;; Correct (two items)
  )

On the other hand, multiple newlines between items are allowed:

  (
    first


    second ;; Correct (two items)
  )

Three different kind of brackets can be used, and are all equivalent: round
brackets "()", square brackets "[]", and curly brackets "{}". The following
lists are all the same:

  (first, second, third)
  [first, second, third]
  {first, second, third}

These different brackets are especially useful to write nested lists:

  (first, [s, e, c, o, n, d], third)

  (
    first
    [s, e, c, o, n, d]
    [{t, h}, i, r ,d]
  )

It is also possible to write lists with angle brackets "<>", but they have
special meaning, and are *not* equivalent to the same list with traditional
brackets:

  <first, second, third> ;; Not the same list as the three above!

In fact, most of the time those angle-bracketed list will simply be ignored. The
only thing the user need to know is that they obey the same parsing rules as the
other lists. For advanced use of the czl format, application can activate an
option in the parser to interpret them as macro lists (see the corresponding
section).


Labels

Both strings and lists can, optionally, be prefixed with a label. The label and
its value are separated by either a colon ":" or an equal sign "=":

  label: in front of a string
  (and, this, is, also, a, labeled=string, inside, a, list)

  label: (in, front, of, a, list)
  label (in, front, of, a, list)  ;; Ambiguous! (interpreted as above)

Unquoted labels cannot not contain any blank space, newline, or special
characters. Use quotes if you need them:

  "this is": correct
  this is: wrong     ;; Ambiguous! (interpreted as above)

  "this: is (also)" = correct

  "
  and this is a
  multi-line label!
  " : unusual

Both the label and its value are mandatory. To specify an empty value, use the
empty string "":

  lang=""
  label=   ;; Ambiguous (interpreted as above: empty string value)

The empty string can also be used as a label, but is equivalent to an unlabeled
value:

  value
  ""=value ;; Same as above
  =value   ;; Ambigous (interpreted as above: no label)

In a given list, the same label can occur multiple times. The correctness of
this is left to the application:

  (first, second, lang=british, lang=american) ;; Correct, depending on the application
  (first, _, third, count=2, count=3)          ;; Ambiguous, depending on the application
  (first, first, first)                        ;; Correct: items don't need to be unique


Shortcuts

  <pi=3.14159>  ;; Discarded: defines a string variable
  <pi>          ;; Evaluates to "3.14157"

  <vowels=(a, e, i, o, u, y)>  ;; Discarded: defines a list variable
  <vowels>                     ;; (a, e, i, o, u, y)

  <foo=$>     ;; Discarded: defines an anchor for the current string
  <foo=$$>    ;; Discarded: defines an anchor for the current list
  <foo=$$$>   ;; Discarded: defines an anchor for the parent list
  <foo=$...>  ;; Discarded: defines an anchor for the top-level ancestor list
  <foo=$$...> ;; Discarded: defines an anchor for the current item of the top-level ancestor list
  <foo>       ;; Copy of the anchored string or list

  <#> <#> <#> <#>                   ;; "0 1 2 3"
  (<#>, <#>, <#>, <#>)              ;; (0, 1, 2, 3)
  (<#>, <#>, [<#>, <#>])            ;; (0, 1, [2, 3])

  (<#>, <##>, <#>, <##>, <#>, <##>) ;; (0, 0, 1, 1, 2, 2)
  (<#>, <##>, <##>, <##>, <#>)      ;; (0, 0, 1, 2, 0, 1)

  (<#=1>, <#>, <#>, <#>)            ;; (1, 2, 3, 4)
  (<#>, <#=10>, <#>, <#=100>)       ;; (0, 10, 11, 100)
  <#=1><#><#><#=1><#><#>            ;; "123123"
  (<#="a">, <#>, <#>, <#>)          ;; (a, b, c, d)
  (<#=foo>, <#>, <#>, <#>)          ;; (<value of foo>, <value of foo+1>, ...)
  (<#=0xf0>, <#>, <#>, <#>)         ;; (0xf0, 0xf1, 0xf2)
  {
    (color <#=1>, pink)             ;; (color 1, pink)
    (color <#>, yellow)             ;; (color 2, yellow)
    (color <#>, blue)               ;; (color 3, blue)
  }
  <#=0+3> <#> <#> <#> <#> <#> <#>   ;; "0 3 6 9 12 15 18"
  <#=0-1> <#> <#> <#> <#>           ;; "0 -1 -2 -3 -4"
  <#=0*2> <#> <#> <#> <#> <#> <#>   ;; "0 1 2 4 8 16 32"
  <#=1/2> <#> <#> <#>               ;; "1 0.5 0.25 0.125"
  <#=5//2> <#> <#> <#>              ;; "5 2 1 0"
  <#=0%3> <#> <#> <#> <#> <#> <#>   ;; "0 1 2 0 1 2 0"
  <#=2%3> <#> <#> <#> <#> <#> <#>   ;; "2 0 1 2 0 1 2"
  <#=0%3+2> <#> <#> <#> <#> <#> <#> ;; "2 3 4 2 3 4 2"

  <?:foo>33                          ;; "33" if foo is defined, string is discarded othewise
  (1, <?:foo>33, 2)                  ;; (1, 33, 2) if foo is defined, (1, 2) otherwise
  <?:foo>(1, 33, 2)                  ;; (1, 33, 2) if foo is defined, list is discarded otherwise
  <!:foo>33                          ;; "44" if foo is *not* defined
  <!:"">33                           ;; Always discarded

  <*:4>hop                           ;; "hophophophop"
  <*:3>(1, 2)                        ;; (1, 2), (1, 2), (1, 2)
  <*:8><#>                           ;; "01234567"
  <*:8><#=1>                         ;; "12345678"

*/
package data
