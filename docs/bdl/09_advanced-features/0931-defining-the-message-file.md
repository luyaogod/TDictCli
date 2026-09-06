---
title: "Defining the message file"
source: "fgl-topics/c_fgl_programs_011.html"
breadcrumb: "Advanced features > Configuration options > OPTIONS (Runtime) > Defining the message file"
type: "concept"
---

# Defining the message file

> The OPTIONS HELP FILE instruction defines the name of the message file.

## Syntax

```
OPTIONS HELP FILE filename
```

## Usage

The `OPTIONS HELP FILE` instruction specifies an expression that returns the
filename of a [help file](../11_user-interface/1593-message-files.md "Message files centralize strings and larger texts identified by a number, that can be used in programs."). This filename can also
include a path name.

Messages in this file can be referenced by number in form-related statements, and are displayed
at runtime when the user presses the Help key.

Message files are found in the directories as described in [the FGLRESOURCEPATH reference topic](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.").
