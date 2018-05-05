/*
Package ciziel provides a parser for the ".czl" configuration format


WORK IN PROGRESS

This is currently only a proof of concept, with no implementation.

Overview

A ciziel document is made of only three types of data: strings, labels, and
tables.

Both strings and labels are arbitrary sequence of characters: they only need to
be quoted if they are to include one the 10 special characters (":,()[]<> and
newline).

  This line is a *single* string.

  "To include the character "" in a quoted string,
  simply double it. All other unicode characters can
  be included as is."

White space is stripped at the beginning and the end of unquoted strings. And if
a quoted string immediately starts or ends with a newline, it is also discarded.

  hop, hop   ,    hop             [three times the same string "hop"]

  "
  A single line.
  "

Tables are made of sections; each section associates a label to a coma-separated
list of data:

  Pink: 255, 33, 140,
  Yellow: 255, 216, 0,
  Blue: 33, 177, 255

Note that the presence or absence of newlines after a coma or a colon is not
meaningful. And when there is a newline, the coma is optional (but not the
colon):

  Lavender: #B57EDC, White: #FFFFFF, Chartreuse Green: #4A8123

The first section of a table is the only one that can be without label; it is
then called a header:

  1, 1, 4                       [Header = first section, if unlabelled]
  1, 2, 3
  1, 1, 3
  1: #CF220E                    [Remaining sections]
  2: #FA9898
  3: #FF6E19
  4: #FCB644
  Size: 3, 3

Tables can be nested, by using parentheses:

  Magenta: (Hex: #D70270, sRGB: 215, 2, 112, Pantone: 226)
  Deep Lavender: (Hex: #734F96, sRGB: 115, 79, 150, Pantone: 258)
  Royal Blue: (Hex: #0038A8, sRGB: 0, 56, 168, Pantone: 168)

Comments are surrounded by square brackets, can be nested and span multiple
lines.

  [
    this is a [single, long] comment,
    spanning multiple lines.
  ]

Finally, small expressions surrounded by angle brackets are used to denote
variables and counters.

  <pi=3.14159>        [Discarded: defines a variable]
  <pi>                [Evaluates to "3.14157"]
  <#>                 [The section counter]
  <@>                 [The item counter]


Ambiguities //TODO

Note that a coma always insert a new item. For example, the following is a list
of *four* items (the second and the fourth are empty strings):

  first,, third,

On the other hand, newlines can only separate non-empty strings; which means
multiple newlines in a row do not create empty items. The following is a list
of two items:

  first


  second

The colon of a label is mandatory, and must be on the same line as the label:

  (this, is (ambiguous))     [Warning: "is" could be a string or a label; assuming string]

It's possible to add one or more newline after the colon (but not before):

  a label:
  (for, this, list)

Labels cannot be empty:

  unlabeled string
  :unlabeled string       [Warning: empty label]
  "": unlabeled string

In a given list, the same label can occur multiple times. The correctness of
this is left to the application:

  (first, second, lang:british, lang:american)  [Correct, depending on the application]

  (first, , third, count:2, count:3)            [Ambiguous, depending on the application]


Variables //TODO

  <pi=3.14159>        [Discarded: defines a variable]
  <pi>                [Evaluates to "3.14157"]


Counters //TODO

  <#>, <#>, <#>, <#>                [1, 2, 3, 4]

  item <#>, item <#>, item <#>      [item 1, item 2, item 3]

  (<#>, <#>, <#>, <#>)              [(1, 1, 1, 1)]

  (<#>, <#>), (<#>, <#>)            [(1, 1), (2, 2)]

  (<#>), (<#>), (<#>), (<#>)        [(1), (2), (3), (4)]

  <#=0>
  <#>, <#>, <#>, <#>                [0, 1, 2, 3]

  <#>, <#=10><#>, <#>, <#=100><#>   [1, 10, 11, 100]

  <#=foo>
  <#>, <#>                          [<value of foo>, <value of foo+1>]


For counting items:

  item<@>, item<@>, item<@>         [item1, item2, item3]

  ~ <#> ~: <@>, <@>, <@>            [~ 1 ~: 1, 2, 3]
  ~ <#> ~: <@>, <@>, <@>            [~ 2 ~: 1, 2, 3]
  ~ <#> ~: <@>, <@>, <@>            [~ 3 ~: 1, 2, 3]

  (<##>, <##>), (<##>, <##>)        [(1, 2), (1, 2)]

  (<##>), (<##>), (<##>), (<##>)    [(1), (1), (1), (1)]

  ((<##>, <##>, <##>, <##>))        [((1, 1, 1, 1))]

  ((<###>, <###>, <###>, <###>))    [((1, 2, 3, 4))]

*/
package ciziel
