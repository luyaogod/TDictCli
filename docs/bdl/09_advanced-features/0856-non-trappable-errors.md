---
title: "Non-trappable errors"
source: "fgl-topics/c_fgl_Exceptions_013.html"
breadcrumb: "Advanced features > Exceptions > Non-trappable errors"
type: "concept"
---

# Non-trappable errors

> Non-trappable errors are fatal errors that generally prevent further program execution.

If a non-trappable error occurs, neither [WHENEVER](0850-whenever-directive.md "Use the WHENEVER directive to define how exceptions must be handled for the rest of the module.")
instructions, nor [TRY/CATCH](0853-try-catch-block.md "Use TRY / CATCH blocks to trap runtime exceptions in a delimited code block.") blocks can trap the
error.

In case of non-trappable error, the runtime system will do the following:

1. In GUI mode, display a pop-up window with the error message,
2. If [`STARTLOG()`](../15_library-reference/2790-startlog.md "Initializes error logging and opens the error log file passed as the parameter.") was
   previously called, write an error record to the log file,
3. Print the error message to stderr,
4. Stop the program.

Some non-trappable errors such as [-1212](../15_library-reference/4483-genero-bdl-errors.md) can occur at runtime initialization. In such case, the GUI error message box and
`STARTLOG()` output cannot be done.

List of non-trappable FGL errors: [-1212](../15_library-reference/4483-genero-bdl-errors.md), [-1260](../15_library-reference/4483-genero-bdl-errors.md), [-1320](../15_library-reference/4483-genero-bdl-errors.md), [-1326](../15_library-reference/4483-genero-bdl-errors.md), [-1328](../15_library-reference/4483-genero-bdl-errors.md), [-1332](../15_library-reference/4483-genero-bdl-errors.md), [-1337](../15_library-reference/4483-genero-bdl-errors.md), [-1338](../15_library-reference/4483-genero-bdl-errors.md), [-1340](../15_library-reference/4483-genero-bdl-errors.md), [-4448](../15_library-reference/4483-genero-bdl-errors.md), [-6000](../15_library-reference/4483-genero-bdl-errors.md), [-6200](../15_library-reference/4483-genero-bdl-errors.md), [-6203](../15_library-reference/4483-genero-bdl-errors.md), [-6204](../15_library-reference/4483-genero-bdl-errors.md), [-6205](../15_library-reference/4483-genero-bdl-errors.md), [-6206](../15_library-reference/4483-genero-bdl-errors.md), [-6207](../15_library-reference/4483-genero-bdl-errors.md), [-6208](../15_library-reference/4483-genero-bdl-errors.md), [-6216](../15_library-reference/4483-genero-bdl-errors.md), [-6217](../15_library-reference/4483-genero-bdl-errors.md), [-6220](../15_library-reference/4483-genero-bdl-errors.md), [-6221](../15_library-reference/4483-genero-bdl-errors.md), [-6222](../15_library-reference/4483-genero-bdl-errors.md), [-6223](../15_library-reference/4483-genero-bdl-errors.md), [-6301](../15_library-reference/4483-genero-bdl-errors.md), [-6302](../15_library-reference/4483-genero-bdl-errors.md), [-6312](../15_library-reference/4483-genero-bdl-errors.md), [-6313](../15_library-reference/4483-genero-bdl-errors.md), [-6316](../15_library-reference/4483-genero-bdl-errors.md), [-6327](../15_library-reference/4483-genero-bdl-errors.md), [-6330](../15_library-reference/4483-genero-bdl-errors.md), [-6366](../15_library-reference/4483-genero-bdl-errors.md), [-6367](../15_library-reference/4483-genero-bdl-errors.md), [-6368](../15_library-reference/4483-genero-bdl-errors.md), [-6606](../15_library-reference/4483-genero-bdl-errors.md), [-8003](../15_library-reference/4483-genero-bdl-errors.md), [-8006](../15_library-reference/4483-genero-bdl-errors.md), [-8007](../15_library-reference/4483-genero-bdl-errors.md), [-8024](../15_library-reference/4483-genero-bdl-errors.md), [-8026](../15_library-reference/4483-genero-bdl-errors.md), [-8070](../15_library-reference/4483-genero-bdl-errors.md), [-8088](../15_library-reference/4483-genero-bdl-errors.md), [-8092](../15_library-reference/4483-genero-bdl-errors.md), [-8107](../15_library-reference/4483-genero-bdl-errors.md), [-8113](../15_library-reference/4483-genero-bdl-errors.md), [-8114](../15_library-reference/4483-genero-bdl-errors.md), [-8126](../15_library-reference/4483-genero-bdl-errors.md), [-8127](../15_library-reference/4483-genero-bdl-errors.md), [-8128](../15_library-reference/4483-genero-bdl-errors.md), [-8130](../15_library-reference/4483-genero-bdl-errors.md), [-8132](../15_library-reference/4483-genero-bdl-errors.md), [-8134](../15_library-reference/4483-genero-bdl-errors.md), [-8135](../15_library-reference/4483-genero-bdl-errors.md), [-8300](../15_library-reference/4483-genero-bdl-errors.md), [-8301](../15_library-reference/4483-genero-bdl-errors.md), [-8303](../15_library-reference/4483-genero-bdl-errors.md), [-8400](../15_library-reference/4483-genero-bdl-errors.md), [-8401](../15_library-reference/4483-genero-bdl-errors.md), [-8500](../15_library-reference/4483-genero-bdl-errors.md), [-8501](../15_library-reference/4483-genero-bdl-errors.md).

Some non-trappable errors may be trappable in a specific context. But as a general rule, consider
the above errors always non-trappable.
