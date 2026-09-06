---
title: "Program stop error message box"
source: "fgl-topics/c_fgl_Migrate_to_310_fatal_errors.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > Program stop error message box"
type: "concept"
---

# Program stop error message box

> In GUI mode, runtime errors stopping the program are now displayed to the end user.

Starting with Genero 3.10.11, when using the GUI mode, runtime errors that stop the program
execution are displayed to the end user in a pop-up message box.

> **Tip:**
>
> Runtime errors stopping the program execution are either non-trappable errors, or
> trappable errors that occur in code pieces not protected by `WHENEVER` or
> `TRY/CATCH`.

It is now much easier to identify the reason of the error stopping the program. Before this
feature, a error stopping the program could not be seen by the end user, and the program windows
just disappear with any further feedback. The only way to identify the problem was to inspect the
fglrun output or errorlog on the application server.

The end user can now take a screenshot and report the issue to the application provider.

> **Note:**
>
> This new behavior is also available with older front-end versions.

Feature extension: Starting with version 3.20.11, an FGLPROFILE entry can define a generic/static
message to be displayed to the end user, to overwrite the original error message and hide issue
details. See [Default exception handling](../09_advanced-features/0855-default-exception-handling.md "Default exception handling must be adapted to your programming pattern.").

## Related links

**Related concepts**  

[Non-trappable errors](../09_advanced-features/0856-non-trappable-errors.md "Non-trappable errors are fatal errors that generally prevent further program execution.")

[WHENEVER directive](../09_advanced-features/0850-whenever-directive.md "Use the WHENEVER directive to define how exceptions must be handled for the rest of the module.")

[TRY - CATCH block](../09_advanced-features/0853-try-catch-block.md "Use TRY / CATCH blocks to trap runtime exceptions in a delimited code block.")

[startlog()](../15_library-reference/2790-startlog.md "Initializes error logging and opens the error log file passed as the parameter.")
