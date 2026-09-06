---
title: "finish"
source: "fgl-topics/c_fgl_Debugger_finish.html"
breadcrumb: "Programming tools > Integrated debugger > Debugger commands > finish"
type: "concept"
---

# finish

> The finish command continues the execution of a program until the current function returns normally.

## Syntax

```
finish
```

## Usage

The `finish` command instructs the program to continue running until
just after the function in the selected stack frame returns, and then stop.

The returned value, if any, is printed.

## Example

```
MAIN
    CALL func1()
END MAIN

FUNCTION func1() RETURNS ()
    DEFINE x INT, s STRING
    CALL func2() RETURNING x, s
    DISPLAY "end of func1"
END FUNCTION

FUNCTION func2() RETURNS (INT,STRING)
    RETURN 999, "xxxxxx"
END FUNCTION
```

```
$ fglcomp -M prog.4gl && fglrun -d prog.42m

(fgldb) break func1
Breakpoint 1 at 0x00000000: file prog.4gl, line 7.

(fgldb) run
Breakpoint 1, func1() at prog.4gl:7
   4
   5     FUNCTION func1() RETURNS ()
   6         DEFINE x INT, s STRING
-> 7         CALL func2() RETURNING x, s
   8         DISPLAY "end of func1"
   9     END FUNCTION
   10

(fgldb) finish
Run till exit func1() at prog.4gl:7
end of func1
Value returned is void
main() at prog.4gl:3
   1     MAIN
   2         CALL func1()
-> 3     END MAIN
   4
   5     FUNCTION func1() RETURNS ()
   6         DEFINE x INT, s STRING
```
