---
title: "break"
source: "fgl-topics/c_fgl_Debugger_break.html"
breadcrumb: "Programming tools > Integrated debugger > Debugger commands > break"
type: "concept"
---

# break

> The break command defines a break point to stop the program execution at a given line or function.

## Syntax

```
break [ location ] [ if condition ]
```

where location
is:

```
{ [module.]function
| [module:]line
}
```

1. module is the name of a source file, without extension.
2. function is a function name.
3. line is a source code line.
4. condition is an expression evaluated dynamically.

## Usage

The `break` command sets a break point at a given source code line of the
program.

When the program is running, the debugger stops automatically at breakpoints defined by this
command.

If a condition is specified, the program stops at the breakpoint only if the
condition evaluates to true. For example, with `break 10 if var =
1`, the debugger adds a conditional breakpoint at line 10, and stops the program execution
at line 10 if the `var` variable is equal to 1.

If you do not specify any location, the breakpoint is created for the current
line.

Breakpoints are defined on source code lines with executable instructions. When the
`break` command specifies a function name or a line number that does not contain an
executable instruction, the debugger automatically selects the next executable line. For example, if
you enter `break main`, the breakpoint will be defined on the first executable line
in the `MAIN` block.

Use the info breakpoints command to list existing breakpoints.

Use the [enable](2598-enable.md "The enable command enables breakpoints that have previously been disabled.") / [disable](2594-disable.md "The disable command disables the specified breakpoint.")
commands to activate / deactivate breakpoints.

Use the [delete](2592-delete.md "The delete command allows you to remove breakpoints that you have specified in your debugger session.") or [clear](2590-clear.md "The clear command clears the breakpoint at a specified line or function.")
command to remove a breakpoint.

## Example

```
MAIN
    DEFINE pkey INTEGER

    LET pkey = 983
    DISPLAY pkey

END MAIN
```

```
$ fglcomp -M prog.4gl && fglrun -d prog.42m

(fgldb) break main
Breakpoint 1 at 0x00000000: file prog.4gl, line 4.

(fgldb) break prog:5 if pkey > 100
Breakpoint 2 at 0x00000000: file prog.4gl, line 5.

(fgldb) break prog:5 if pkey > 200
Note Breakpoint 2 also set at pc 0x00000000.
Breakpoint 3 at 0x00000000: file prog.4gl, line 5.

(fgldb) info breakpoints
Num Type           Disp Enb Address    What
1   breakpoint     keep y   0x00000000 in main at prog.4gl:4
2   breakpoint     keep y   0x00000000 in main at prog.4gl:5
	stop only if pkey > 100
3   breakpoint     keep y   0x00000000 in main at prog.4gl:5
	stop only if pkey > 200

(fgldb) disable 2
(fgldb) delete 3
(fgldb) info breakpoints
Num Type           Disp Enb Address    What
1   breakpoint     keep y   0x00000000 in main at prog.4gl:4
2   breakpoint     keep n   0x00000000 in main at prog.4gl:5
	stop only if pkey > 100
```

## Related links

**Related concepts**  

[BREAKPOINT](../09_advanced-features/0832-breakpoint.md "The BREAKPOINT instruction sets a program breakpoint when running in debug mode.")

[Defining a breakpoint in the code](2582-defining-a-breakpoint-in-the-code.md "Set a breakpoint in the program source code with the BREAKPOINT instruction.")
