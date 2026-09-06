---
title: "Understanding message files"
source: "fgl-topics/c_fgl_message_files_002.html"
breadcrumb: "User interface > Form definitions > Message files > Understanding message files"
type: "concept"
---

# Understanding message files

> This is an introduction to message files.

Message files define text messages with a unique integer identifier.

Several message files can be created and loaded by the same program.

Message files are typically used to implement application help system, and are especially
designed for the TUI mode.

In order to use a message file, do the following:

1. Create the .msg source message file with a text editor.
2. Compile the source message file with fglmkmsg to create the
   .iem binary format.
3. Copy the binary file to a distribution directory.
4. In programs, specify the message file with the `OPTIONS HELP FILE`
   instruction.
5. Use a specific message with the `HELP` clause of dialogs, or load a
   given message with the `SHOWHELP()` function.

Message files provide a simple way to implement a help system in your application.

For other application messages and texts, consider using localized strings instead of
message files.

## Related links

**Related concepts**  

[Localization](../09_advanced-features/0863-localization.md "Localization support allows you to implement programs that follow specific language and cultural rules.")

[Localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.")
