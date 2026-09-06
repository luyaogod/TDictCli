---
title: "4GL programming topics"
source: "fgl-topics/c_fgl_MigI4GL_022.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics"
type: "concept"
---

# 4GL programming topics

> When migrating from I4GL to Genero BDL, review the programming differences between the two products. Reviewing the differences allows you to plan and prepare for a smooth migration.


## Child topics

- [Form specification file syntax](0374-form-specification-file-syntax.md): This topic describes syntax differences between I4GL and FGL in .per form specification file.
- [Dynamic arrays](0375-dynamic-arrays.md): Support for dynamic arrays differs between I4GL and Genero BDL.
- [Debugger command syntax](0376-debugger-command-syntax.md): While I4GL and Genero BDL both provide a program debugger, the commands used and how it is used can differ.
- [GLOBALS variables usage](0377-globals-variables-usage.md): This topic describes differences between I4GL and FGL regading global variables definitions and usage.
- [Strict function signature checking](0378-strict-function-signature-checking.md): With Genero BDL, a function's signature is detected at link time.
- [STRING versus CHAR/VARCHAR](0379-string-versus-char-varchar.md): Genero BDL supports the STRING data type in addition to CHAR and VARCHAR. While the STRING data type is useful in certain situations, there are times when you should use CHAR or VARCHAR instead.
- [Review user-made C routines](0380-review-user-made-c-routines.md): Genero BDL provides libraries which may replace some of the C routines required by I4GL applications.
- [Web Services support](0381-web-services-support.md): Both I4GL and Genero BDL support Web services, albeit with different implementations.
- [File I/O statements and APIs](0382-file-i-o-statements-and-apis.md): Both I4GL and Genero BDL support accessing files on operating systems. The base.Channel built-in class supported by Genero BDL can also open streams to subprocesses and sockets.
- [SQL cursor management](0383-sql-cursor-management.md): This topic describes differences between I4GL and Genero BDL in SQL cursor management.
- [arg_val() returns NULL if no argument](0384-arg-val-returns-null-if-no-argument.md): If the index passed to the function references an argument that does not exist, the arg_val() function must be handled differently between I4GL and Genero BDL.
- [Checking function parameters/returns](0385-checking-function-parameters-returns.md): With Genero BDL, a parameter count mismatch is detected at link time.
- [Using variable subscripts in SQL](0386-using-variable-subscripts-in-sql.md): Use of subscript operations on host variables in static SQL statements is not supported by Genero BDL.
- [MONEY type and CLIENT_LOCALE](0387-money-type-and-client-locale.md): I4GL adapts the scale of MONEY / MONEY(p) types from the client locale setting, Genero BDL always uses the same default scale.
- [Variable assignment in function calls](0388-variable-assignment-in-function-calls.md): Genero BDL behaves differently to I4GL when variables are changed in expressions evaluated in function parameters.
- [REPORT syntax and behavior](0389-report-syntax-and-behavior.md): This topic describes differences between I4GL and Genero BDL in the report engine.
- [String to numeric type conversions](0390-string-to-numeric-type-conversions.md): This topic describes differences between I4GL and FGL in string to numeric type conversions.
- [Status variable handling](0391-status-variable-handling.md): I4GL and Genero BDL support the status variable differently.
- [Creating multi-byte chars with ASCII](0392-creating-multi-byte-chars-with-ascii.md): I4GL allows to create multi-byte characters with the ASCII operator. This is not supported by Genero BDL.
- [CONSTRUCT and Informix ANSI database](0393-construct-and-informix-ansi-database.md): I4GL CONSTRUCT generates the ANSI-mode SQL table owner in the WHERE clause, while Genero BDL does not.
- [Writing expression errors in log file](0394-writing-expression-errors-in-log-file.md): I4GL writes expression errors in the log file defined by startlog(), while FGL does not.
- [LOAD and UNLOAD instrutions](0395-load-and-unload-instrutions.md): This topic describes differences between I4GL and FGL with the LOAD and UNLOAD instructions.
- [Command line tools](0396-command-line-tools.md): I4GL and FGL provide different compiler commands with different options.
- [Overlaping error numbers in ranges](0397-overlaping-error-numbers-in-ranges.md): Genero BDL uses error numbers in a range of Informix SQL errors.
- [Fetching numbers into CHAR/VARCHAR](0398-fetching-numbers-into-char-varchar.md): I4GL and FGL format numbers in a different way when fetching direcly MONEY, DECIMAL column values into CHAR/VARCHAR variables.
