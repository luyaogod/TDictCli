---
title: "TERMINFO"
source: "fgl-topics/c_fgl_EnvVariables_TERMINFO.html"
breadcrumb: "Configuration > Environment variables > Operating system environment variables > TERMINFO"
type: "concept"
---

# TERMINFO

> Defines the terminal capabilities database.

On UNIX™ platforms, the TERMINFO environment variable
points to the terminal capabilities database. This variable must be used along with TERM, when
INFORMIXTERM is set to `terminfo`.

Setting this environment variable is generally not necessary. The default is defined by the UNIX system, it can be for example
/etc/terminfo, /usr/lib/terminfo, or
/lib/terminfo.

It is important to define terminal capabilities properly for your text terminal hardware,
or the terminal emulation you are using. In particular, function keys (F1, F16) and display
attributes (bold, reverse, colors) may not work if the escape sequences do not correspond
to the terminal used.

For more details about the TERMINFO environment variable, please refer to your UNIX operating system manual.

## Related links

**Related concepts**  

[TERM](0500-term.md "Defines the type of terminal on UNIX platforms.")

[INFORMIXTERM](0539-informixterm.md "Defines terminal control library to be used.")

[Using a text terminal](../11_user-interface/1531-using-a-text-terminal.md "This section covers topics about text terminal configuration when using the TUI mode.")
