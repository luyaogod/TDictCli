---
title: "Code beautifier tool usage"
source: "fgl-topics/c_fgl_beautifier_usage.html"
breadcrumb: "Programming tools > Source code beautifier > Code beautifier tool usage"
type: "concept"
---

# Code beautifier tool usage

> This topic describes how to use the code beautifier tool.

## Using the code beautifier of fglcomp

The code beautifier is a feature of the fglcomp compiler which is enabled with
the `--format` option:

```
$ fglcomp --format module.4gl
```

Several formatting `--fo-*` options are available and explained below.

The code beautifier supports uppercase or lowercase keyword coding conventions. See the `--fo-lowercase-keywords` option usage for more details.

## Using a configuration file

All `--fo-*` formatting options (except `--fo-inplace`) can be
specified in a configuration file named .fgl-format in the souce directory or
in one of the parent directories of the source.

If no .fgl-format is found, fglcomp --format will read
the file specified with the `--fo-fallback-style=filename`
option.

This configuration file is typically placed in the top directory of your project, to apply the
same formatting rules to all sources located under this root
directory:

```
$ tree -a .
.
`-- top
    |-- .fgl-format
    |-- app1
    |   `-- main.4gl
    `-- lib
        `-- myutils.4gl
3 directories, 3 files

$ cd top

$ cat .fgl-format 
--fo-indent-width=5
--fo-pack=1

$ cd lib

$ fglcomp --format myutils.4gl 
PRIVATE DEFINE debug_level SMALLINT
FUNCTION set_debug_level(l SMALLINT)
     LET debug_level = l
END FUNCTION
```

## Controlling the beautifier result output

By default, the reformatted source is written to the standard output stream (stdout).

To overwrite the original source file, use the `--fo-inplace`
option:

```
$ fglcomp --format --fo-inplace module.4gl
```

Before reformatting the source file with the `--fo-inplace` option,
fglcomp creates a copy of the original source file, by adding a tilde character
after the .4gl file extension:
filename.4gl~.

The `--fo-inplace` option cannot be used in the .fgl-format
configuration file.

## Disabling formatting for a group of lines

The formatter understands special control comments that disable formatting for a block of source
code lines.

The code between single-line comments having `"fgl-format off/on"` keywords will
not be formatted.

The formatting control comments can use the `--`, `#` or `{
}` [comment syntax](../08_language-basics/0550-source-comments.md "The --, # and { } characters can be used to add source comments."), and must appear
on a line by itself.

The control comment lines themselves will be formatted and aligned, depending on the surrounding
source code lines.

For example:

```
MAIN
    DEFINE formatted INT = 0
    -- fgl-format off
       define  unformatted  int=1
    -- fgl-format on
    DEFINE formatted_again INT = 2
END MAIN
```

> **Tip:**
>
> Make sure that the formatting control comments are placed in the same source code
> block and scope. For example, if you start with an `fgl-format off` control comment
> just before a `FOR` loop, and you put the closing `fgl-format on`
> control comment inside the `FOR` loop, the formatter will produce unexpected
> alignments.

## Formatting a range of lines

By default, the complete source file is reformatted. Use the
`--fo-lines=start-line:end-line` option to
specify a range of lines to be formatted.

The `--fo-lines` option is especially useful with version control systems, to
limit the commit change to the piece of code that is updated and then reformatted.

To format only a range of lines of a source file, use the `--fo-lines` option by
specifying the starting and ending
line:

```
$ fglcomp --format --fo-lines=51:57 module.4gl
```

The `--fo-lines` option can be used multiple times to specify several pieces of
code to be reformatted.

## Defining the maximum width for a source code line

Use the `--fo-column-limit=cols` option to specify the width
limit for a source code line:

```
$ cat module.4gl
PUBLIC FUNCTION update_customer( id INTEGER, name VARCHAR(50) )
    UPDATE customer SET cust_name = name WHERE cust_id = id AND cust_valid = 'Y'
END FUNCTION

$ fglcomp --format --fo-column-limit=40 module.4gl
PUBLIC FUNCTION update_customer(
    id INTEGER, name VARCHAR(50))
    UPDATE customer
        SET cust_name = name
        WHERE cust_id = id
            AND cust_valid = 'Y'
END FUNCTION
```

The default is 80.

## Defining the base indentation width

The `--fo-indent-width=cols` option can be used to set width
for base indentations:

```
$ cat module.4gl
PUBLIC FUNCTION cleanup( mode INTEGER )
CASE
 WHEN mode==1 DISPLAY "Mode 1"
 WHEN mode==2 DISPLAY "Mode 2"
 END CASE
END FUNCTION

$ fglcomp --format --fo-indent-width=6 module.4gl
PUBLIC FUNCTION cleanup(mode INTEGER)
      CASE
^^^^^^ -- 6 spaces from --fo-indent-width
            WHEN mode == 1
      ^^^^^^ -- 6 spaces from --fo-indent-width
                  DISPLAY "Mode 1"
            WHEN mode == 2
                  DISPLAY "Mode 2"
      END CASE
END FUNCTION
```

The default is 4.

## Defining the continuation indentation width

The `--fo-continuation-indent-width=cols` option defines the
indentation width for line-breaks. When breaking long lines, this number of spaces inserted for each
extra
level:

```
$ cat module.4gl
PUBLIC FUNCTION update_customer( id INTEGER, name VARCHAR(50) )
    UPDATE customer SET cust_name = name WHERE cust_id = id AND cust_valid = 'Y'
END FUNCTION

$ fglcomp --format --fo-column-limit=40 --fo-continuation-indent-width=2 module.4gl
PUBLIC FUNCTION update_customer(
  id INTEGER, name VARCHAR(50))
^^ -- 2 spaces from --fo-continuation-indent-width
    UPDATE customer
^^^^ -- 4 spaces from --fo-indent-width default
      SET cust_name = name
    ^^ -- 2 spaces from --fo-continuation-indent-width
      WHERE cust_id = id
    ^^ -- 2 spaces from --fo-continuation-indent-width
        AND cust_valid = 'Y'
      ^^ -- 2 spaces from --fo-continuation-indent-width
END FUNCTION
```

Default is 4.

## Using TAB as indentation character

By default,
the indentation character is a blank space: It is used all over to produce
indentation.

Specify the `--fo-use-tab=1` formatter option, to produce tabs
whenever it is needed to fill whitespace that spans at least from one tab stop to the next one,
based on `--fo-tab-width` and `--fo_indent-width` options.

Use
the `--fo-tab-width=cols` option to define the number columns used
for tab stops.

For example, with the default `--fo-indent-width` of 4 columns and
the default `--fo-tab-width` or 8 columns, `--fo-use-tab=1` will put
tabs whenever 8 columns tab stops are found, and put 4 spaces to complete the
indentation:

```
$ fglcomp --format --fo-use-tab=1 module.4gl 
PUBLIC FUNCTION cleanup(mode INTEGER)
    CASE
^^^^ -- 4 spaces
        WHEN mode == 1
< tab  > -- 1 tab
            DISPLAY "Mode 1"
< tab  >^^^^ -- 1 tab + 4 spaces
        WHEN mode == 2
            DISPLAY "Mode 2"
    END CASE
END FUNCTION
```

The
next example produces only tabs, because `--fo-tab-with` is set to 4 and matches the
default
`--fo-indent-width`:

```
$ fglcomp --format --fo-use-tab=1 --fo-tab-width=4 module.4gl 
PUBLIC FUNCTION cleanup(mode INTEGER)
        CASE
< tab  > -- 1 tab
                WHEN mode == 1
< tab  >< tab  > -- 2 tabs
                        DISPLAY "Mode 1"
< tab  >< tab  >< tab  > -- 3 tabs
                WHEN mode == 2
                        DISPLAY "Mode 2"
        END CASE
END FUNCTION
```

## Controlling instruction clause indentation

The `--fo-label-indent=1` option indicates that indentation should apply to
instruction sub-clauses, such as the `WHEN` clause of a `CASE` /
`END CASE` block. When zero, no sub-clause indentation is
done:

```
$ cat module.4gl
PUBLIC FUNCTION cleanup( mode INTEGER )
CASE
 WHEN mode==1 DISPLAY "Mode 1"
 WHEN mode==2 DISPLAY "Mode 2"
 END CASE
END FUNCTION

$ fglcomp --format --fo-label-indent=0 module.4gl
PUBLIC FUNCTION cleanup(mode INTEGER)
    CASE
    WHEN mode == 1
        DISPLAY "Mode 1"
    WHEN mode == 2
        DISPLAY "Mode 2"
    END CASE
END FUNCTION
```

When 0, the instruction clauses are aligned with the head of the statement
(`CASE`, INPUT etc). If 1, those "labels" are indented. This results in indenting
the statements following those labels 2 levels relatively to the head of the statement.

The default is 1 (clause indentation is enabled).

## Controlling how source lines are packed

The `--fo-pack=1` option can be used to pack items as much as possible together on
the same
line:

```
$ cat module.4gl
PUBLIC FUNCTION func1( )
CALL func2( "aaa", "bbb", "ccc", "ddd", "eee" )
END FUNCTION

$ fglcomp --format --fo-column-limit=30 module.4gl
PUBLIC FUNCTION func1()
    CALL func2(
        "aaa",
        "bbb",
        "ccc",
        "ddd",
        "eee")
END FUNCTION

$ fglcomp --format --fo-column-limit=30 --fo-pack=1 module.4gl
PUBLIC FUNCTION func1()
    CALL func2(
        "aaa", "bbb", "ccc",
        "ddd", "eee")
END FUNCTION
```

## Using uppercase or lowercase keywords

Genero BDL uses traditionally uppercase keywords and therefore the code formatter produces
uppercase keywords by default.

When the coding conventions require lowercase keywords, use the
`--fo-lowercase-keywords=1` option of the formatter, to produce lower case
keywords:

```
$ cat module.4gl
function myfunc(x INT) retuRNS int
  return x + 1
end funcTION

$ fglcomp --format --fo-lowercase-keywords=1 module.4gl 
function myfunc(x int) returns int
    return x + 1
end function

$ fglcomp --format --fo-lowercase-keywords=0 module.4gl 
FUNCTION myfunc(x INT) RETURNS INT
    RETURN x + 1
END FUNCTION
```

## Aligning consecutive items

Several formatting options are provided to align consecutive language elements.

The `--fo-align-consecutive-assignments=1` option aligns consecutive
assignements:

```
    LET name = "Phil"
    LET da   = get_last_date()
    LET id   = 1032

    LET rec.cust_id      = 1243
    LET rec.cust_name    = "McTough"
    LET rec.cust_address = "5 Sunset Av."
```

Alignment never produces additional linebreaks. The formatter produces groups of aligned lines,
for example (with `--fo-align-consecutive-assignments=1` and
`--fo-column-limit=40`):

```
LET long_var_name_1      = 1
LET very_long_var_name_2 = 1
LET short_name = "aaaaaaaaaaaaaaa"
LET tn         = "aaaaaaaaaaaaaaa"
```

The `--fo-align-consecutive-types=1` option aligns types in consecutive variable
definitions:

```
    DEFINE v1           INTEGER
    DEFINE current_date DATE
    DEFINE last_name    VARCHAR(50)

    DEFINE v2   INTEGER
    DEFINE rate DECIMAL(10, 2)
```

The `--fo-align-trailing-comments=1` option aligns consecutive
comments:

```
    DISPLAY a    -- first comment
    DISPLAY foo, -- second comment
        bar      -- third comment
```

## The fglformatdiff command

The fglformatdiff command reads a unified diff from the
(stdin) standard input stream, calculates what lines have changed and calls the
formatter to beautify the modified lines.

The fglformatdiff command is provided to integrate with version control
systems such as CVS abd SVN. If GIT is your version control system, use
fglgitformat instead.

The diff output must be unified and without any context lines (diff -U0).

The following example uses a diff output produced from the svn tool (note the
options to get a unified diff
format):

```
svn diff --diff-cmd=diff -x-U0 | fglformatdiff --inplace
```

See also [fglformatdiff](2528-fglformatdiff.md "The fglformatdiff tool reformats source lines of a unified diff text.")
command reference.

## The fglgitformat command

The fglgitformat command can reformat code changes that are tracked in a GIT
repository.

For example, the next command reformats the lines of code changes registered in the 2 last
commits (`HEAD~2`), modifies directly the related sources
(`--inplace`), and prints the changes to the stdout stream
(`--verbose`):

```
fglgitformat --verbose --inplace HEAD~2
```

See also [fglgitformat](2527-fglgitformat.md "The fglgitformat tool reformats the source lines changed in a GIT history.")
command reference.
