---
title: "TERMCAP terminal capabilities"
source: "fgl-topics/c_fgl_DynamicUI_032.html"
breadcrumb: "User interface > User interface basics > Using a text terminal > TERMCAP terminal capabilities"
type: "concept"
description: "When the INFORMIXTERM environment variable is set to termcap or when this variable is undefined, the runtime system will use the termcap terminal capabilities database. The termcap solution is ..."
---

# TERMCAP terminal capabilities

When the INFORMIXTERM environment variable is set to `termcap` or
when this variable is undefined, the runtime system will use the
termcap terminal capabilities database.

The termcap solution is provided for backward compatibility. You
should use terminfo instead, by setting the INFORMIXTERM variable
to `terminfo`.

The default termcap database is in the /etc/termcap file. If this file is
not found, the runtime system will use its default file $FGLDIR/etc/termcap.
Use the TERMCAP environment variable to specify a different termcap file as the
defaults. If you plan to modify the default termcap file, we strongly recommend
that you make a copy of the original file and point to the new file with the TERMCAP variable.

In this section we will briefly describe the syntax of the termcap file. For
a complete definition please refer to your operating system documentation (see man pages describing
the termcap file syntax).

## Related links

**Related concepts**  

[INFORMIXTERM](../07_configuration/0539-informixterm.md "Defines terminal control library to be used.")

[TERM](../07_configuration/0500-term.md "Defines the type of terminal on UNIX platforms.")

[TERMCAP](../07_configuration/0501-termcap.md "Defines the termcap terminal capabilities database on UNIX platforms.")

## Child topics

- [Termcap syntax](1534-termcap-syntax.md)
- [Genero-specific termcap definitions](1535-genero-specific-termcap-definitions.md)
