---
title: "source"
source: "fgl-topics/c_fgl_Debugger_source.html"
breadcrumb: "Programming tools > Integrated debugger > Debugger commands > source"
type: "concept"
---

# source

> The source command executes a file of debugger commands.

## Syntax

```
source cmdfile
```

1. cmdfile is the name of the file containing the debugger commands.

## Usage

The `source` command allows you to execute a command file of lines that
are debugger commands.

The lines in the file are executed sequentially.

The commands are not printed as they are executed, and any messages are not displayed.

Commands are executed without asking for confirmation.

An error in any command terminates execution of the command file.

## Example

Using the text file cmdfile.txt, which contains the single line with a
`break` command:

prog.4gl:

```
MAIN
    DEFINE x INTEGER

    FOR x = 100 TO 200
        DISPLAY x
    END FOR

END MAIN
```

cmdfile.txt:

```
break main
run
step
step
print x
quit
```

```
$ fglcomp -M prog.4gl && fglrun -d prog.42m

(fgldb) source cmdfile.txt
Breakpoint 1 at 0x00000000: file prog.4gl, line 4.
Breakpoint 1, main() at prog.4gl:4
   1     MAIN
   2         DEFINE x INTEGER
   3
-> 4         FOR x = 100 TO 200
   5             DISPLAY x
   6         END FOR
   7
   2         DEFINE x INTEGER
   3
   4         FOR x = 100 TO 200
-> 5             DISPLAY x
   6         END FOR
   7
   8     END MAIN
        100
   1     MAIN
   2         DEFINE x INTEGER
   3
-> 4         FOR x = 100 TO 200
   5             DISPLAY x
   6         END FOR
   7
$1 = 100
```
