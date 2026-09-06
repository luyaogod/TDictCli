---
title: "Linking the utility functions library"
source: "fgl-topics/c_fgl_Migrate_to_200_fgl_libraries.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.00 upgrade guide > Linking the utility functions library"
type: "concept"
---

# Linking the utility functions library

> All utility functions are in the libfgl4js.42x library, up until 2.21.

Prior to version 2.00, some utility functions (canvas draw\* and database db\_\* functions) were
linked automatically to the 42r program when using `fglrun -l` or
fgllink. These functions are implemented in the
fgldraw.4gl and fgldbutl.4gl modules, which
were linked in the libfgl.42x library and loaded automatically at
runtime by fglrun.

Starting with version 2.00, all utility functions are now in the
libfgl4js.42x library. So, if you use the draw\* or db\_\* utility
functions, you must now add the libfgl4js.42x library explicitly when
using `fglrun -l` or fgllink, or you can use the [fgl2p](../13_programming-tools/2515-fgl2p.md "The fgl2p tool compiles source files and assembles p-code modules into a .42r program or a .42x library.") tool to link .42r programs. The
fgl2p tool links the program with the
libfgl4js.42x library by default.

Starting with version 2.21, the libfgl.42x library is no longer
provided.

## Related links

**Related concepts**  

[Utility modules](../15_library-reference/2794-utility-modules.md "A utility function is a function provided in a separate library; it is not built in the runtime system.")

[Command reference](../13_programming-tools/2512-command-reference.md "Command line tools provided by FGLGWS packages.")
