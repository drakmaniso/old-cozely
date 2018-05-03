/*
Package data provides a parser for the ".czl" configuration format


WORK IN PROGRESS

This is currently only a proof of concept, with no implementation.

Comments

Comments are surrounded by square brackets:

  [This whole line is a comment]
  (1, 2, 3) [Comment on the same line as meaningful data]

They can contain any character, including newlines and special characters; they
can even contain nested square brackets:

  [
    this is a [single, long] comment,
    spanning multiple lines!
  ]

Note that the only character that cannot appear in a comment is a closing square
bracket without an opening one (as this would simply end the comment). Note also
that forgetting to close a square bracket is an error, as this would will extend
the comment until the end of the file.


Strings

A basic string is a sequence of any non-special characters; they always ends
with the first newline or special character. For example, the following example
contain 8 strings:

  First string, Second string;
  3.14159
  4/04/2004
  50 x 50
  6, 7, 8th

The three correct way to terminate basic strings is the newline, the coma and
the semi-colon. All other special characters will also act as separators, but
will either have special meaning or issue a warning.

Note that white spaces (e.g. space characters, tabs) are stripped at the
beginning and the end of basic strings, but not inside:

  foo,foo  ,   foo   [three times the same string "foo"]
  foo          bar   [a single string of 16 characters "foo          bar"]

There is 13 special characters:

  " ; ,
  : =
  ( ) { }
  [ ] < >

Note that backslash is not a special character. The following is a valid string:

  C:\Program Files\Cozely

A quoted string is a character sequence between double quotes; any character is
allowed including newlines:

  "This (quoted) string: contains several special characters."
  "This is a
  multi-line string"

To insert a double quote character inside a quoted string, use two in
succession:

  "Here is a ""quote"" inside a string."  [= Here is a "quote" inside a string.]

If a quoted string immediately starts and ends with a newline character, those
are stripped:

  "
  This is a
  two-line string
  "

Lists

A list is a sequence of items surrounded by round or curly brackets, and
separated by comas or semi-colon. All the following examples are equivalent:

  (first, second, third)
  {first, second, third}
  (first; second; third)
  {first, second; third}      [Warning: mixing comas and semi-colon in the same list]
  (first, second, third}      [Warning: mismatched closing bracket]

Lists can be nested:

  (first is a string, (second, is, a, list))
  (first is a string; {second, is, a, list})
  {first is a string; ({s;e;c;o;n;d}, is, a, list, of, {l;i;s;t;s})}

Lists can contain newlines, which are ignored:

  (first, second,
  third)
  (
    first,
    second,
    third
  )

Comas and semi-colons are optional at end of lines:

  (first, second
  third)
  (
    first
    second
    third
  )

Note that a coma always insert a new item:

  (,,)        [a list of three empty strings: ("", "", "")]
  (
    ,
    second,
  )           [a list of three strings: ("", "second", "")]

On the other hand, multiple newlines in a row do not create multiple items (in
other words, a newline only creates a new item if it also terminates the
previous one):

  (
    first


    second
    third
  )            [a list of three strings: ("first", "second", "third")]


Labels

Both strings and lists can, optionally, be prefixed with a label. The label and
its value are separated by either a colon ":" or an equal sign "=":

  label: in front of a string
  (and, this, is, also, a, labeled=string, inside, a, list)
  label: (in, front, of, a, list)

  {this, is (ambiguous)}     [Warning: "is" could be a string or a label; assuming string]

  this is: correct

Labels follow exactly the same rules as strings; i.e., basic labels can contain
any non-special characters except newlines, while quoted labels can contain
anything:

  "this is": correct

  "this: is (also)" = correct

  "
  and this is a
  multi-line label!
  " : unusual

The empty (quoted) string can also be used as a label, but is equivalent to an
unlabeled value:

  value
  ""=value    [Same as above]

If a colon or an equal sign is present, the label is mandatory:

  =incorrect  [Warning: no label; assuming empty label ""]

Basic labels are case-insensitive.

In a given list, the same label can occur multiple times. The correctness of
this is left to the application:

  (first, second, lang=british, lang=american) [Correct, depending on the application]
  (first, _, third, count=2, count=3)          [Ambiguous, depending on the application]
  (first, first, first)                        [Correct: items don't need to be unique]


Shortcuts

  <pi=3.14159>        [Discarded: defines a string constant]
  <pi>                [Evaluates to "3.14157"]

  <foo:>33            [Discarded: defines a string anchor]
  <foo>               ["33"]
  <vowels:>(1, 2, 3)  [Discarded: defines a list anchor]
  <vowels>            [(1, 2, 3)]

  <#> <#> <#> <#>                    ;; "0 1 2 3"
  (<#>, <#>, <#>, <#>)               ;; (0, 1, 2, 3)
  (<#>, <#>, {<#>, <#>})             ;; (0, 1, {2, 3})

  (<#>, <##>, <#>, <##>, <#>, <##>)  ;; (0, 0, 1, 1, 2, 2)
  (<#>, <##>, <##>, <##>, <#>)       ;; (0, 0, 1, 2, 0, 1)

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
  <#+3> <#> <#> <#> <#> <#> <#>     ;; "0 3 6 9 12 15 18"
  <#=3-1> <#> <#> <#>               ;; "3 2 1 0"
  <#=3-1> <#> <#> <#> <#> <#>       ;; Ambiguous! Interpreted as: "3 2 1 0 0 0"
  <#=$> <#> <#> <#>                 ;; "3 2 1"
  <#=$> <#> <#> <#> <#> <#>         ;; "5 4 3 2 1 0"
  <#*2> <#> <#> <#> <#> <#> <#>     ;; "0 1 2 4 8 16 32"
  <#=8/2> <#> <#> <#> <#>           ;; "8 4 2 1 0"
  <#=5/2> <#> <#> <#>               ;; "5 2 1 0"
  <#%3> <#> <#> <#> <#> <#> <#>     ;; "0 1 2 0 1 2 0"
  <#=2%3> <#> <#> <#> <#> <#> <#>   ;; "2 0 1 2 0 1 2"
  <#%3+2> <#> <#> <#> <#> <#> <#>   ;; "2 3 4 2 3 4 2"

  <foo?>33                          ;; "33" if foo is defined, string is discarded othewise
  (1, <foo?>33, 2)                  ;; (1, 33, 2) if foo is defined, (1, 2) otherwise
  <foo?>(1, 33, 2)                  ;; (1, 33, 2) if foo is defined, list is discarded otherwise
  <foo!>33                          ;; "44" if foo is *not* defined
  <!>33                             ;; Always discarded

  <*4>hop                           ;; "hophophophop"
  <*3>(1, 2)                        ;; (1, 2), (1, 2), (1, 2)
  <*8><#>                           ;; "01234567"
  <*8><#=1>                         ;; "12345678"


*/
package data
