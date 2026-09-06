---
title: "Using message files at runtime"
source: "fgl-topics/c_fgl_message_files_005.html"
breadcrumb: "User interface > Form definitions > Message files > Using message files at runtime"
type: "concept"
description: "In order to use compiled message files ( .iem ) in programs, specify the current message file with the OPTIONS HELP FILE command: OPTIONS HELP FILE \"mymessages.iem\" The message file provided in the ..."
---

# Using message files at runtime

In order to use compiled message files (`.iem`) in programs, specify the current
message file with the [`OPTIONS HELP FILE`](../09_advanced-features/0931-defining-the-message-file.md "The OPTIONS HELP FILE instruction defines the name of the message file.")
command:

```
OPTIONS HELP FILE "mymessages.iem"
```

The message file provided in the `OPTIONS HELP FILE` command is searched for in
several directories, as described in [the
FGLRESOURCEPATH reference topic](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.").

After the message file is defined, you can start the help viewer by calling the [`SHOWHELP()`](../15_library-reference/2789-showhelp.md "Displays a runtime help text.")
function:

```
CALL showhelp(1242)
```

Use the `HELP` clause in a dialog instruction such
as `INPUT` to define particular message number for that the
dialog:

```
INPUT BY NAME ... HELP 455
```

The help viewer will automatically display the message text corresponding to the number when the
user presses the help key. By default, the help key is Ctrl-W in TUI mode and F1 in GUI mode.

Note that you can implement your own help viewer by overloading
the `SHOWHELP()` function defined in $FGLDIR/src/fglhelp.4gl.
This allows you to customize the help system for your
application.
