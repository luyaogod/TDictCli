---
title: "Logging options"
source: "fgl-topics/c_fgl_logging.html"
breadcrumb: "Programming tools > Logging options"
type: "concept"
---

# Logging options

> Logging solutions allow you to display exchanges between components when a program executes.

Genero provides several build-in options to get debug information, as well as logging features:

- Get the stack trace with [`base.Application.getStackTrace()`](../15_library-reference/2976-base-application-getstacktrace.md "Returns the function call stack trace.")
- Display the GUI protocol exchange in stderr with [FGLGUIDEBUG](../11_user-interface/1527-debugging-the-front-end-protocol.md).
- Display the SQL statements execution in stderr with [FGLSQLDEBUG](../07_configuration/0534-fglsqldebug.md "Defines the debug level for tracing SQL instructions.").
- Display Web Services API calls in stderr with [FGLWSDEBUG](../07_configuration/0537-fglwsdebug.md "The FGLWSDEBUG environment variable enables web services library debugging.").
- Produce function call stack trace, with [`fglrun
  --trace`](2632-execution-trace.md "Print a function call stack of your program.") option.
- Produce source code coverage info with [`fglrun
  --merge-cov` / FGLCOV](2628-source-code-coverage.md "Collect information about used source lines").
- Log front-end protocol exchange with [`fglrun
  --start-guilog`](../11_user-interface/1528-front-end-protocol-logging.md) option.
- Front-ends can add specific markers in the front-end protocol log with [GUI log events](../11_user-interface/1529-gui-log-events.md).
- Produce application log files in case of runtime error, with [STARTLOG()](../15_library-reference/2790-startlog.md "Initializes error logging and opens the error log file passed as the parameter.").

> **Important:**
>
> Sensitive and personal data may be written to the output. Make sure that the log output is
> written to files that can only be read by application administrators.
