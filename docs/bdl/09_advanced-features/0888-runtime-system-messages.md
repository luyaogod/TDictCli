---
title: "Runtime system messages"
source: "fgl-topics/c_fgl_localization_026.html"
breadcrumb: "Advanced features > Localization > Application locale > Runtime system messages"
type: "concept"
---

# Runtime system messages

> This section describes how to translate default English runtime system message files in a different language.

## Understanding system message handling

Runtime system error messages are provided in .msg and
.iem message files. The system message files use the same technique as user
defined message files. The default message files (.msg) are located in the
$FGLDIR/msg/en\_US directory.

For backward compatibility with IBM® Informix® 4GL, some of these system error messages are used by the runtime
system to display messages during a dialog instruction. For example, end users may get the error
message [-1309](../15_library-reference/4483-genero-bdl-errors.md) "*There are no
more rows in the direction you are going*" when scrolling an a `DISPLAY ARRAY`
list in TUI mode.

If your application language is not English, you will need to translate some of the system
messages to a specific locale and language. If your application language is English, you might just
want to customize the default messages.

Here are some examples of system messages that can appear at runtime:

| Number | Message text |
| --- | --- |
| -1204 | Invalid year in date. |
| -1205 | Invalid month in date. |
| -1206 | Invalid day in date. |
| -1213 | A character to numeric conversion process failed. |
| -1218 | String to date conversion error. |
| -1301 | This value is not among the valid possibilities. |
| -1302 | The two entries were not the same -- please try again. |
| -1303 | You cannot use this editing feature because a picture exists. |
| -1304 | Error in field. |
| -1305 | This field requires an entered value. |
| -1306 | Please type again for verification. |
| -1307 | Cannot insert another row - the input array is full. |
| -1309 | There are no more rows in the direction you are going. |
| -8105 | Not found. |
| -8133 | Input lost for field '%s' |
| -8138 | Too many input characters. |

In order to customize runtime system messages displayed during a dialog, you have two options:

1. Define localized strings with a specific key name using the error number.
2. Patch the .msg files, recompile to .iem and distribute
   in your own folder.

The option #1 is the recommended solution: Since .msg/.iem error messages
are shared by the compilers, runtime system error and dialog interaction messages for the end user,
it is better to keep .msg/.iem error messages in English for programming, and
have a specific localization solution for messages displayed to the end user during a dialog
instruction.

## Defining dialog messages with localized strings

First read the [Localized strings](0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.") chapter to understand how
localized strings work.

When a dialog instruction needs to display an error message such as `-1304 "Error in
field."`, the runtime system first reads .42s localized strings resource files with a
specific string key using the error number. If this string is not found, it will fallback to the
legacy .msg/.iem error message.

In your .str localized string resource file, define entries with a key name
using the `fgl.dialog.error.` prefix, followed by a dot a the error number with the
minus
sign:

```
fgl.dialog.error.<error-number-without-minus-sign>
```

or, for informational (non-error) messages identified with a positive
number:

```
fgl.dialog.message.<message-number>
```

The string text can hold a %s sprintf placeholder for error messages with parameter.

For example, to define the text for the errors -1213, -1218 and
-1304:

```
"fgl.dialog.error.1213" = "Invalid number."
"fgl.dialog.error.1218" = "Invalid date."
"fgl.dialog.error.1304" = "Input error."
"fgl.dialog.error.8133" = "Input was lost for field '%s'"
```

Compile the .str source into .42s, distribute the
.42s in your program resource directory, reference the file in your fglprofile
and you are done.

## Customizing .msg/.iem message files

First read the [Message files](../11_user-interface/1593-message-files.md "Message files centralize strings and larger texts identified by a number, that can be used in programs.") chapter to understand how message
files work.

> **Tip:**
>
> Customizing .msg/.item message files is the legacy solution to translate
> runtime system error messages. Consider using solution #1 based on localized strings.

To customize .msg/.iem system messages, do this:

1. Create a new directory under $FGLDIR/msg, using the same name as your
   current locale. For example, if `LANG=fr_FR.ISO8859-1`, you must create
   `$FGLDIR/msg/fr_FR.ISO8859-1`.
2. Copy the original system message source files (.msg) from
   `$FGLDIR/msg/en_US` to the locale-specific directory.
3. Edit the source files with the .msg suffix and translate the messages.
4. Recompile the message files with the [fglmkmsg](../13_programming-tools/2518-fglmkmsg.md "The fglmkmsg tool compiles .msg message files into a binary version used by programs.") tool to produce .iem files. Make sure you
   have set the correct locale!
5. Run a program to check if the new messages are used.

With this technique, you can deploy multiple message files in different languages and locales in
the same FGLDIR/msg directory.

You can use the fglmkmsg tool with the `-r` option to revert a
.iem file to a source .msg file.

There is no need to translate all messages of the .msg files: Most of the
error messages are unexpected during a program execution and therefore can stay in English. The
messages subject of translation can be found in the 4glusr.msg and
rds.msg files.

The locale can be set with different environment variables (see setlocale manual pages for more
details). To identify the locale name, the runtime system first looks for the LC\_ALL value, then
LC\_CTYPE and finally LANG.

Pay attention to locale settings when editing message files and compiling with
fglmkmsg: The current locale must match the locale used in the
.msg files.

The .iem files used at runtime must match the current locale used by
programs. This is automatic, as long as you put the correct files in the corresponding
$FGLDIR/msg/$LANG directory.

## Related links

**Related concepts**  

[Localized strings](0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.")

[Message files](../11_user-interface/1593-message-files.md "Message files centralize strings and larger texts identified by a number, that can be used in programs.")
