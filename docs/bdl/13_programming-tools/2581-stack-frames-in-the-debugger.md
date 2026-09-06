---
title: "Stack frames in the debugger"
source: "fgl-topics/c_fgl_Debugger_006.html"
breadcrumb: "Programming tools > Integrated debugger > Stack frames in the debugger"
type: "concept"
---

# Stack frames in the debugger

> The stack frame contains information about function call stack.

Each time your program performs a function call, information about
the call is saved in a block of data called a *stack frame*.
Each frame contains the data associated with one call to one function.

The stack frames are allocated in a region of memory called the *call stack*. When your
program is started, the stack has only one frame, that of the function
`main`. This is the initial frame, also known as the *outermost
frame*. As the debugger executes your program, a new frame is made each time a
function is called. When the function returns, the frame for that function call is
eliminated.

The debugger assigns numbers to all existing stack frames, starting
with zero for the innermost frame, one for the frame that called it,
and so on upward. These numbers do not really exist in your program;
they are assigned by the debugger to allow you to designate stack
frames in commands.

Each time your program stops, the debugger automatically selects
the currently executing frame and describes it briefly. You can use
the [frame](2600-frame.md "The frame command selects and prints a stack frame.") command
to select a different frame from the current call stack.

## Related links

**Related concepts**  

[Runtime stack](../09_advanced-features/0833-runtime-stack.md "The runtime stack is used to pass/return values to/from functions.")
