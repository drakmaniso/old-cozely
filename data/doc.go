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
    spanning multiple lines.
  ]

Note that the only thing that cannot appear in a comment is a closing square
bracket without an opening one (as this would simply end the comment).


Strings

A string is a sequence characters terminated by a coma or a newline, or both.
The following example contains 6 strings:

  First string, 2nd string, 3.14159
  4/4/04, 50 x 50, 600

Some characters have special meaning, and cannot appear in basic strings:

  " , : ( ) [ ] < >

Outside of these 9 special characters and the newline, every unicode character
is allowed. For example, since the backslash is not a special character, the
following is a valid string:

  C:\Program Files\Cozely

Note also that white space (i.e. space characters and tabs) is discarded at the
beginning and the end of basic strings, but not inside:

  foo,foo  ,   foo   [three times the same string "foo"]
  foo          bar   [a single string of 16 characters "foo          bar"]

A quoted string is a character sequence between double quotes; any character is
allowed, including newlines:

  "This (quoted) string: contains several special characters."
  "This is a
  multi-line string"

To insert a double quote character inside a quoted string, use two in
succession:

  """A lot of "" in this string """    ["A lot of " in this string"]

If a quoted string immediately starts and ends with a newline character, those
are discarded:

  "
  This string has
  only two lines
  "

Lists

A list is a sequence of items separated by comas:

  first item, second item, third item

Lists can be nested:

  first item is a string, (second, is, a, list)

Lists can contain newlines, which are ignored:

  first, second,
  third

Comas are optional when all items are on a separate line:

  first
  second
  third

Which is is equivalent to:

  first,
  second,
  third

It's even possible to mix coma-separated and newline-separated items in the same
list:

  first, second
  third

Note that a coma always insert a new item. For example, the following is a list
of *four* items (the second and the fourth are empty strings):

  first,, third,

On the other hand, newlines can only separate non-empty strings; which means
multiple newlines in a row do not create empty items. The following is a list
of two items:

  first


  second


Labels

A label is a character sequence terminated by a colon ":", and placed in front
of a string or a list.

  a label: in front of a string
  label: (in, front, of, a, list)
  (and, this, is, a labeled: string, inside, a, list)

  (this, is (ambiguous))     [Warning: "is" could be a string or a label; assuming string]

It's possible to add one or more newline after the colon (but not before):

  a label:
  (for, this, list)

Labels follow exactly the same rules as strings; i.e., basic labels can contain
any non-special characters except newlines, while quoted labels can contain
anything:

  this *is*: correct

  "this: is (also)": correct

  "
  and this is a
  multi-line label!
  ":
  (correct, but probably not very useful)

Labels cannot be empty:

  unlabeled string
  :unlabeled string       [Warning: empty label]
  "": unlabeled string

In a given list, the same label can occur multiple times. The correctness of
this is left to the application:

  (first, second, lang:british, lang:american)  [Correct, depending on the application]

  (first, , third, count:2, count:3)            [Ambiguous, depending on the application]


Variables

  <pi=3.14159>        [Discarded: defines a variable]
  <pi>                [Evaluates to "3.14157"]


Counters

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


Sub-List Counter

  (<##>, <##>, <##>, <##>)          [(1, 2, 3, 4)]

  (<##>, <##>), (<##>, <##>)        [(1, 2), (1, 2)]

  (<##>), (<##>), (<##>), (<##>)    [(1), (1), (1), (1)]

  ((<##>, <##>, <##>, <##>))        [((1, 1, 1, 1))]

  ((<###>, <###>, <###>, <###>))    [((1, 2, 3, 4))]

*/
package data
