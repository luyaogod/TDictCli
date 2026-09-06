---
title: "signal"
source: "fgl-topics/c_fgl_Debugger_signal.html"
breadcrumb: "Programming tools > Integrated debugger > Debugger commands > signal"
type: "concept"
---

# signal

> The signal command sends an interruption signal to the program.

## Syntax

```
signal signal
```

## Usage

The `signal` comment resumes execution where your program stopped, but immediately
gives it the signal signal.

signal can be the name or the number of a signal.

For example, on many systems signal 2 and signal SIGINT are both ways of sending an interrupt
signal. The `signal SIGINT` command resumes execution of your program where it
has stopped, but immediately sends an interrupt signal. The source line that was current when
the signal was received is displayed.

> **Note:**
>
> The current version only allows the `SIGINT` signal.

## Example

```
(fgldb) signal SIGINT 
Program exited normally.
16      for i = 1 to 10
(fgldb)
```
