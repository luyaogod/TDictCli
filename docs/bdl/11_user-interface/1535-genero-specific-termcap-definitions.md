---
title: "Genero-specific termcap definitions"
source: "fgl-topics/c_fgl_DynamicUI_034.html"
breadcrumb: "User interface > User interface basics > Using a text terminal > TERMCAP terminal capabilities > Genero-specific termcap definitions"
type: "concept"
description: "Controlling the delay for ESC key When using termcap, the ESCDELAY environment variable can be set to mimic the feature of the terminfo implementation: ESCDELAY specifies the delay (in milliseconds) ..."
---

# Genero-specific termcap definitions

## Controlling the delay for ESC key

When using termcap, the ESCDELAY environment variable can be set to mimic the feature
of the terminfo implementation: ESCDELAY specifies the delay (in milliseconds)
before returning the ESC as a single key. Otherwise, ESC is the leading key of a
control character sequence. The granularity of the delay value is 100 milliseconds.
If ESCDELAY is not set (or set to 0), a default value of 500 milliseconds is
used.

```
$ export ESCDELAY=600
```

## Extending Function Key Definitions

In TUI mode, the runtime system recognizes function keys F1 through F36. These
keys correspond to the termcap capabilities k0 through k9, followed
by kA through kZ.

The termcap entry for these capabilities is the sequence of ASCII characters your
terminal sends when you press the function keys (or any other keys you choose to use
as function keys).

This example shows some function key definitions for
the xterm terminal:

```
k0=\E[11~:k1=\E[12~:k2=\E[13~:k3=\E[14~:\
... 
k9=\E[21~:kA=\E[23~:kB=\E[24~:\
```

## Defining dialog action keys

Dialog action keys for insert, delete and list navigation can be defined
with the following capabilities:

- `ki` : Insert line (default is CTRL-J)
- `kj` : Delete line (default is CTRL-K)
- `kf` : Next page (default is CTRL-M)
- `kg` : Previous page (default is CTRL-N)

Use the [OPTIONS](../09_advanced-features/0932-defining-control-keys.md "The OPTIONS logical-key KEY physical-key instruction defines physical keys for logical keys (TUI mode).") statement to name other function keys
or CTRL keys for these operations.

## Specifying Characters for Window Borders

The runtime system uses the graphics characters in the termcap file when you
specify a window border in an `OPEN WINDOW` statement.

The runtime system uses characters defined in the termcap file to draw the
border of a window. If no characters are defined in this file, the runtime system uses the hyphen (
- ) for horizontal lines, the vertical bar ( | ) for vertical lines, and the plus sign ( + ) for
corners.

Steps to define the graphical
characters for window borders for your terminal type:

1. Determine the escape sequences for turning the terminal graphics
   mode ON and OFF (Refer to the manual of your terminal). For example,
   on Wyse 50 terminals, the escape sequence for entering graphics mode
   is ESC H^B, and the escape sequence for leaving graphics mode is ESC
   H^C.
2. Identify the ASCII equivalents for the six graphics characters
   that Genero requires to draw the window borders. The ASCII equivalent
   of a graphics character is the key you would press in graphics mode
   to obtain the indicated character. The six graphical characters needed
   by Genero are:
   1. The upper left corner
   2. The lower left corner
   3. The upper right corner
   4. The lower right corner
   5. The horizontal line
   6. The vertical line
3. Edit the termcap entry for your terminal, and define the following
   string capabilities:
   - `gs` : The escape sequence for entering graphics mode. In the
     termcap file, ESCAPE is represented as a backslash ( \ ) followed by the letter
     E; CTRL is represented as a caret ( ^ ). For example, the Wyse 50 escape sequence ESC-H CTRL-B is
     represented as \EH^B.
   - `ge` : The escape sequence for leaving graphics mode. For example, the Wyse 50
     escape sequence ESC-H CTRL-C is represented as `\EH^C`.
   - `gb` : The concatenated, ordered list of ASCII equivalents for the six graphics
     characters used to draw the border. Using the order as listed in (2).

     For example, if you are
     using a Wyse 50 terminal, you would add the following, in a linear
     sequence:

     `:gs=\EH^B:ge=\EH^C:gb=2135z6:\`

For terminals without graphics capabilities, you must enter a blank value for the
`gs` and `ge` capabilities. For `gb`, enter the
characters you want Genero to use for the window border. The following example shows possible values
for `gs`, `ge`, and `gb` in an entry for a terminal
without graphics capabilities:

```
:gs=:ge=:gb=.|.|_|:
```

With these settings, window borders would be drawn using underscores ( `_` ) for
horizontal lines, vertical bars ( `|` ) for vertical lines, periods (
`.` ) for the top corners, and vertical bars ( `|` ) for the lower
corners.

## Adding Color and Intensity

In TUI mode, a Genero program can be written either for a monochrome or a color terminal, and
then you can run the program on either type of terminal. If you set up the
termcap files as described, the color attributes and the intensity attributes
are related.

| Number | Color | Intensity | Note |
| --- | --- | --- | --- |
| 0 | WHITE | NORMAL |  |
| 1 | YELLOW | BOLD |  |
| 2 | MAGENTA | BOLD |  |
| 3 | RED | BOLD (\*) | If the keyword BOLD is indicated as the attribute, the field will be RED on a color terminal |
| 4 | CYAN | DIM |  |
| 5 | GREEN | DIM |  |
| 6 | BLUE | DIM (\*) | If the keyword DIM is indicated as the attribute, the field will be BLUE on a color terminal |
| 7 | BLACK | INVISIBLE |  |

The background for colors is BLACK in all cases. In either color or monochrome mode,
you can add the REVERSE, BLINK, or UNDERLINE attributes if your terminal supports them.

## The ZA String Capability

Genero uses a parameterized string capability named `ZA` in the
termcap file to determine color assignments. Unlike other termcap string
capabilities that you set to a literal sequence of ASCII characters, ZA is a function string that
depends on the following four parameters:

| Parameter | Name | Description |
| --- | --- | --- |
| 1 | p1 | Color number between 0 and 7 (see Table 1). |
| 2 | p2 | 0 = Normal; 1 = Reverse. |
| 3 | p3 | 0 = No-Blink; 1 = Blink. |
| 4 | p3 | 0 = No-underscore; 1 = Underscore. |

ZA uses the values of these four parameters and a stack
machine to determine which characters to send to the terminal. The
ZA function is called, and these parameters are evaluated, when a
color or intensity attribute is encountered in a Genero program. Use
the information in your terminal manual to set the ZA parameters to
the correct values for your terminal.

The ZA string uses stack
operations to push values onto the stack or to pop values off the
stack. Typically, the instructions in the ZA string push a parameter
onto the stack, compare it to one or more constants, and then send
an appropriate sequence of characters to the terminal. More complex
operations are often necessary; by storing the display attributes
in static stack machine registers (named a through z), you can have
terminal-specific optimizations.

The different stack operators
that you can use to write the descriptions are summarized here. For
a complete discussion of stack operators, see your operating system
documentation.

## Operators that Send Characters to the Terminal

- `%d` pops a numeric value from the stack and sends a maximum of three digits to
  the terminal. For example, if the value 145 is at the top of the stack, `%d` pops the
  value off the stack and sends the ASCII representations of 1, 4, and 5 to the terminal. If the value
  2005 is at the top of the stack, `%d` pops the value off the stack and sends the
  ASCII representation of 5 to the terminal.
- `%2d` pops a numeric value from the stack and sends a maximum of two digits to
  the terminal, padding to two places. For example, if the value 145 is at the top of the stack,
  `%2d` pops the value off the stack and sends the ASCII representations of 4 and 5 to
  the terminal. If the value 5 is at the top of the stack, `%2d` pops the value off the
  stack and sends the ASCII representations of 0 and 5 to the terminal.
- `%3d` pops a numeric value from the stack and sends a maximum of three digits to
  the terminal, padding to three places. For example, if the value 7 is at the top of the stack,
  `%3d` pops the value off the stack and sends the ASCII representations of 0, 0, and 7
  to the terminal.
- `%c` pops a single character from the stack and sends it to the terminal.

## Operators that Manipulate the Stack

- `%p[1-9]` pushes the value of the specified parameter on the stack. The notation
  for parameters is p1, p2, ... p9. For example, if the value of p1 is 3, `%p1` pushes
  3 on the stack.
- `%P[a-z]` pops a value from the stack and stores it in the specified variable.
  The notation for variables is Pa, Pb, ... Pz. For example, if the value 45 is on the top of the
  stack, `%Pb` pops 45 from the stack and stores it in the variable Pb.
- `%g[a-z]` gets the value stored in the corresponding variable
  (`P[a-z]`) and pushes it on the stack. For example, if the value 45 is stored in the
  variable `Pb`, `%gb` gets 45 from Pb and pushes it on the stack.
- `%'c'` pushes a single character on the stack. For example, `%'k'`
  pushes k on the stack.
- `%{n}` pushes an integer constant on the stack. The integer can be any length and
  can be either positive or negative. For example, `%{0}` pushes the value 0 on the
  stack.
- `%S[a-z]` pops a value from the stack and stores it in the specified static
  variable. (Static storage is nonvolatile since the stored value remains from one attribute
  evaluation to the next.) The notation for static variables is `Sa`,
  `Sb`, ... `Sz`. For example, if the value 45 is on the top of the
  stack, `%Sb` pops 45 from the stack and stores it in the static variable Sb. This
  value is accessible for the duration of the Genero program.
- `%G[a-z]` gets the value stored in the corresponding static variable
  (`S[a-z]`) and pushes it on the stack. For example, if the value 45 is stored in the
  variable Sb, `%Gb` gets 45 from `Sb` and pushes it on the stack.

## Arithmetic Operators

Each arithmetic operator pops the top two values from the stack, performs an operation,
and pushes the result on the stack.

- `%+` Addition.

  For example, `%{2}%{3}%+` is equivalent to
  2+3.
- `%-` Subtraction.

  For example, `%{7}%{3}%-` is equivalent to
  7-3.
- `%*` Multiplication.

  For example, `%{6}%{3}%*` is equivalent to
  6\*3.
- `%/` Integer division.

  For example, `%{7}%{3}%/` is equivalent
  to 7/3 and produces a result of 2.
- `%m` Modulus (or remainder).

  For example, `%{7}%{3}%m` is
  equivalent to (7 mod 3) and produces a result of 1.

## Bit Operators

The following bit operators pop the top two values from the stack, perform
an operation, and push the result on the stack:

- `%&` Bit-and.

  For example, `%{12}%{21}%&` is equivalent
  to (12 and 21) and produces a result of 4.
- `%|` Bit-or.

  For example, `%{12}%{21}%|` is equivalent to (12 or
  21) and produces a result of 29.
- `%^` Exclusive-or.

  For example, `%{12}%{21}%^` is equivalent to
  (12 exclusive-or 21) and produces a result of 25.

The following unary operator pops the top value from the stack,
performs an operation, and pushes the result on the stack:

- `%~` Bitwise complement.

  For example, `%{25}%~` results in a
  value of -26.

## Logical Operators

The following logical operators pop the top two values from the stack,
perform an operation, and push the logical result (0 for false or
1 for true) on the stack:

- `%=` Equal to.

  For example, if the parameter `p1` has the value
  3, the expression `%p1%{2}%=` is equivalent to 3=2 and produces a result of 0
  (false).
- `%>` Greater than.

  For example, if the parameter p1 has the value 3, the
  expression `%p1%{0}%>` is equivalent to 3>0 and produces a result of 1
  (true).
- `%<` Less than.

  For example, if the parameter p1 has the value 3, the
  expression `%p1%{4}%<` is equivalent to 3<4 and produces a result of 1
  (true).

The following unary operator pops the top value from the stack,
performs an operation, and pushes the logical result (0 or 1) on the
stack.

- `%!` Logical negation.

  This operator produces a value of zero for all nonzero
  numbers and a value of 1 for zero. For example, `%{2}%!` results in a value of 0, and
  `%{0}%!` results in a value of 1.

## Conditional Statements

The conditional statement has the following format:

```
%? expr %t thenpart %e elsepart %;
```

The `%e elsepart` is optional. You can nest conditional
statements in the thenpart or the elsepart.

When Genero evaluates a conditional statement, it pops the top value from the stack and evaluates
it as either true or false. If the value is true, the runtime performs the operations after the
`%t;` otherwise it performs the operations after the `%e` (if
any).

For example, the expression:

```
%?%p1%{3}%=%t;31%;
```

is
equivalent to:

```
if p1 = 3 then print ";31"
```

Assuming that p1 in the example has the value 3, Genero would perform the following steps:

- `%?` does not perform an operation but is included to make the conditional
  statement easier to read.
- `%p1` pushes the value of p1 on the stack.
- `%{3}` pushes the value 3 on the stack.
- `%=` pops the value of p1 and the value 3 from the stack, evaluates the boolean
  expression p1=3, and pushes the resulting value 1 (true) on the stack.
- `%t` pops the value from the stack, evaluates 1 as true, and executes the
  operations after `%t`. (Since "`;31`" is not a stack machine
  operation, Genero prints ";31" to the terminal.)
- `%;` terminates the conditional statement.

## ZA example

The ZA sequence for the ID Systems Corporation ID231 (color terminal) is:

```
ZA =
\E[0;                # Print lead-in
%?%p1%{0}%=%t%{7}    # Encode color number (translate color number
                     # to number for the ID231 term)
%e%p1%{1}%=%t%{3}    # 
%e%p1%{2}%=%t%{5}    # 
%e%p1%{3}%=%t%{1}    #
%e%p1%{4}%=%t%{6}    #
%e%p1%{5}%=%t%{2}    #
%e%p1%{6}%=%t%{4}    #
%e%p1%{7}%=%t%{0}%;  #
%?%p2%t30;%{40}%+%2d # if p2 is set, print '30' and '40' + color number (reverse)
%e40;%{30}%+%2d%;    #  else print '40' and '30' + color number (normal)
%?%p3%t;5%;          # if p3 is set, print 5 (blink)
%?%p4%t;4%;          # if p4 is set, print 4 (underline)
m                    # print 'm' to end character sequence
```
