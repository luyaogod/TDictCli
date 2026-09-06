---
title: "INFORMIXTERM"
source: "fgl-topics/c_fgl_EnvVariables_INFORMIXTERM.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > INFORMIXTERM"
type: "concept"
---

# INFORMIXTERM

> Defines terminal control library to be used.

The INFORMIXTERM environment variable indicates what terminal capabilities database must be
used by the runtime system when running a program in TUI mode on a dumb terminal.

Possible values of INFORMIXTERM are `terminfo` and `termcap`. If
the variable is not set, it defaults to `termcap`.

When set to `termcap` (the default), the runtime system reads terminal
capabilities from the file defined by the TERMCAP environment variable.

When set to `terminfo`, the runtime system reads terminal capabilities from the
terminfo database of the system (curses/ncurses).

TERMCAP is the older implementation of terminal capabilities database. Therefore, it is
recommended to set INFORMIXTERM=`terminfo`.

## Related links

**Related concepts**  

[TERM](0500-term.md "Defines the type of terminal on UNIX platforms.")

[TERMCAP](0501-termcap.md "Defines the termcap terminal capabilities database on UNIX platforms.")
