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


Macros

Any list written with angle brackets is parsed as a macro or an expression. By
default, they will simply be ignored; if they have a label, it is also ignored.
An option in the parser has to be turned on in order to use their functionality.

  <pi=3.14159>  ;; Defines a variable (evaluates to nothing)
  <pi>          ;; Evaluates to "3.14157"

  <vowels=(a, e, i, o, u, y)>  ;; Defines a list constant (evaluates to nothing)
  <vowels>                     ;; Evaluates to the list (a, e, i, o, u, y)

  <$foo>     ;; Defines a named anchor (evaluates to nothing)

  /0         ;; The first top-level item
  /1         ;; The second top-level item
  /0/0       ;; The first item in the first top-level item (which must be a list)
  .          ;; The current list
  ./.        ;; The current item
  ..         ;; The parent list
  ../..      ;; The parent of the parent list
  ...        ;; The top-level ancestor list
  .../.      ;; The current item in the top-level ancestor list
  ./0        ;; The first item in the current list
  ./+1       ;; The next item in the current list
  ./-1       ;; The previous item in the current list
  ../0       ;; The first item in the parent list
  ../+1      ;; The item following the current list
  foo        ;; The item containing the anchor foo
  foo/..     ;; The list containing the anchor foo

  <#.>       ;; Index of the current list (position in the parent list)
  <#..>      ;; Index of the parent list (position in the grand-parent list)
  <#../+1>   ;; Index of the item following the current list
  <#...>     ;; Index of the top-level parent list
  <#.../.>   ;; Index of the current item of the top-level parent list
  <#foo>     ;; Index of the string containing anchor foo
  <#foo/..>  ;; Index of the list containing anchor foo

  <@/0/0>    ;; Copy of the first item of the first top-level item
  <@./0>     ;; Copy of the first item of the current list
  <@./0/0>   ;; Copy of the first item of the first item of the current list
  <@./+1>    ;; Copy of the following item
  <@../+1>   ;; Copy of the item following the current list
  <@.../+1>  ;; Copy of the next top-level item
  <@foo>     ;; Copy of the item following anchor foo
  <@foo/..>  ;; Copy of the list containing anchor foo


Expressions

  <1 2 +>         ;; "3"
  <a=2 a>       ;; "2"
  <1 1 a=+ a>   ;; "2"
  <a=2 a ++>    ;; "11"
  <1 2 3 +>       ;; Error! (the stack contains 2 items at the end of the expression)
  <1 2 3 + !>     ;; Two items: "1", "5"
  <pi 2 *>       ;; "6.28318"
  <"foo" "bar" ~>     ;; "foobar"
  <"foo" 33 ~>      ;; "foo33"
  <111 33 ~>      ;; "11133"
  <# 1 +>         ;; Index of the current item plus one
  <1 # +>         ;; Same as above
  <# 1 + 2 *>     ;; Double of the index of the current item

  <pi 3.14159 ?eq>  ;; "true"
  <3 44 ?gt>         ;; "false"

  <true then:"foo">              ;; "foo"
  <false then:"foo">             ;; Nothing
  <false then:"foo" else:"bar">  ;; "bar"

  <0 for:3 do:<10 +>>                   ;; "30"
  <"" i=2 for:<i 5 ?lt> do:<i ~ i ++>>   ;; "234"

*/
package data
