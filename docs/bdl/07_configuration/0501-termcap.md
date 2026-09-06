---
title: "TERMCAP"
source: "fgl-topics/c_fgl_EnvVariables_TERMCAP.html"
breadcrumb: "Configuration > Environment variables > Operating system environment variables > TERMCAP"
type: "concept"
description: "Defines the termcap terminal capabilities database on UNIX platforms."
---

# TERMCAP

> Defines the termcap terminal capabilities database on UNIX™ platforms.

## Usage

For UNIX platforms, TERMCAP is an environment variable
that defines the terminal capabilities file. This variable must be used in conjunction with TERM,
when INFORMIXTERM is set to *termcap*, or when INFORMIXTERM is not set.

If
the TERMCAP variable is not defined, Genero tries to open /etc/termap.
If no /etc/termcap file exists,
the runtime system uses $FGLDIR/etc/termcap.
You can add more terminal definitions in this file.

TERMCAP is the older implementation of terminal capabilities database. it is recommended that you set
INFORMIXTERM=terminfo.

It is important to define terminal capabilities properly for your text terminal
hardware or the terminal emulation you are using. Especially function keys (F1, F16)
and display attributes (bold, reverse, colors) may not work if the escape sequences
do not correspond to the terminal used.

For more details about the TERMCAP environment variable, please refer to your
UNIX operating system manual.

## Related links

**Related concepts**  

[TERM](0500-term.md "Defines the type of terminal on UNIX platforms.")

[INFORMIXTERM](0539-informixterm.md "Defines terminal control library to be used.")

[Using a text terminal](../11_user-interface/1531-using-a-text-terminal.md "This section covers topics about text terminal configuration when using the TUI mode.")
