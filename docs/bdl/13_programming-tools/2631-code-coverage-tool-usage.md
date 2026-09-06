---
title: "Code coverage tool usage"
source: "fgl-topics/c_fgl_coverage_usage.html"
breadcrumb: "Programming tools > Source code coverage > Code coverage tool usage"
type: "concept"
---

# Code coverage tool usage

> This topic describes how to use the code coverage feature.

## Steps to get coverage information

Follow the next steps to collect and interpret coverage data:

1. Delete all existing coverage data files (\*.42m.cov)
2. Enable coverage data file generation by setting the FGLCOV environment variable to 1.
3. Execute your program several times (with fglrun): This will produce
   module.42m.cov files.
4. Use the fglrun --merge-cov command, to merge
   module.42m.cov coverage data files and
   module.4gl source files, to produce readable coverage
   information files (module.4gl.cov)
5. Inspect the module.4gl.cov to find how many times source
   code lines are executed.

## Cleanup coverage data files

Before starting a new coverage data collection session, remove existing
\*.42m.cov data files.

```
$ rm -f *.42m.cov
```

## Enable coverage data generation

Before starting a program, set the [FGLCOV](../07_configuration/0521-fglcov.md "Enables code coverage data collection.")
environment variable, to instruct fglrun to produce coverage data
files:

```
$ export FGLCOV=1
```

In this topic, we will use the following program sample, to show how coverage data is produced
(pay attention to the line containing the `IF FALSE THEN` instruction: This will
illustrate how unreachable lines are detected):

```
MAIN
   DEFINE i INTEGER
   FOR i=1 TO 10
       IF i MOD 2 == 0 THEN
          CALL func1()
       END IF
   END FOR
END MAIN

FUNCTION func1()
   DISPLAY "Hello"
   IF FALSE THEN
      DISPLAY "... world!"
   END IF
END FUNCTION
```

A coverage data file (module.42m.cov) will be created for
each module (module.42m), when running a
program:

```
$ fglcomp prog.4gl

$ fglrun prog
Hello
Hello
Hello
Hello
Hello

$ ls prog.*
prog.42m  prog.42m.cov  prog.4gl
```

The coverage data file is created in the directory where the module is found.

> **Important:**
>
> Coverage information is collected for several program executions: If a
> module.42m.cov file exist for a module, the execution count
> of source code lines will be summed up.

## Merge coverage data with source files

The .42m.cov coverage data files contain only numbers. In order to interpret
coverage data properly, you need to merge the coverage data files with the source file, by using the
fglrun --merge-cov command:

```
$ fglrun --merge-cov prog.4gl
fglcov: prog.4gl.cov created
```

If the .4gl source file and the .42m.cov coverage data
file are not located in the same directory as the .42m module, specify the full
path to the coverage data file as second parameter for the --merge-cov
option:

```
fglrun --merge-cov mymodule.4gl ../dist/lib/mymodule.42m.cov
```

After a first time program execution, the resulting prog.4gl.cov file
contains:

```
        -:    1:MAIN
        -:    2:   DEFINE i INTEGER
       11:    3:   FOR i=1 TO 10
       10:    4:       IF i MOD 2 == 0 THEN
        5:    5:          CALL func1()
        -:    6:       END IF
        -:    7:   END FOR
        1:    8:END MAIN
        -:    9:
        -:   10:FUNCTION func1()
        5:   11:   DISPLAY "Hello"
        5:   12:   IF FALSE THEN
    =====:   13:      DISPLAY "... world!"
        -:   14:   END IF
        5:   15:END FUNCTION
```

In the above output, a couple of code coverage indicators can be interpreted as follows:

- Line #2 shows a `-` (hyphen-minus) indicator: The `DEFINE` statement
  is not reachable, because it is a declaration statement.
- Line #3 shows the number 11: The `FOR` (head) instruction has been executed 11
  times (10 + 1 = 11, to jump out of the loop)
- Line #4 shows the number 10: The `IF` statement inside the `FOR`
  loop has been executed 10 times.
- Line #13 shows `=====`: This line is reachable, but is never executed, since the
  parent `IF` condition always evaluates to `FALSE`.

By executing the sample program a second time, the existing coverage data files are completed
with additional information:

```
$ fglrun prog
```

Merge again the data files with source
files:

```
$ fglrun --merge-cov prog.4gl
fglcov: prog.4gl.cov created
```

And now check the new result
file:

```
        -:    1:MAIN
        -:    2:   DEFINE i INTEGER
       22:    3:   FOR i=1 TO 10
       20:    4:       IF i MOD 2 == 0 THEN
       10:    5:          CALL func1()
        -:    6:       END IF
        -:    7:   END FOR
        2:    8:END MAIN
        -:    9:
        -:   10:FUNCTION func1()
       10:   11:   DISPLAY "Hello"
       10:   12:   IF FALSE THEN
    =====:   13:      DISPLAY "... world!"
        -:   14:   END IF
       10:   15:END FUNCTION
```

You can see that the execution counts of source code lines are summed up across program
execution.
