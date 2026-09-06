---
title: "Loading presentation styles"
source: "fgl-topics/c_fgl_presentation_styles_011.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Using presentation styles > Loading presentation styles"
type: "concept"
---

# Loading presentation styles

> Presentation styles are defined in an XML file with a 4st extension. In order to load the presentation styles, the runtime system needs to locate the appropriate style file.

By default, the runtime system searches for default.4st in several
directories, as described in [the
FGLRESOURCEPATH reference topic](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files."). If the file is not found using the FGLRESOURCEPATH/DBPATH
environment variables, default presentation styles are loaded from the
$FGLDIR/lib/default.4st file.

Overwrite the default search by loading a specific presentation style file with the [`ui.Interface.loadStyles()`](../15_library-reference/3114-ui-interface-loadstyles.md "Load the presentation styles file.")
method:

```
MAIN
  CALL ui.Interface.loadStyles("mystyles")
  ...
END MAIN
```

This method accepts an absolute path with the .4st extension, or
a simple filename without the .4st extension. If you give a simple file
name, for example "`mystyles`", the runtime system searches for the
mystyles.4st file in the current directory or in the application directory. If
the file does not exist, it searches in the directories defined by the FGLRESOURCEPATH environment
variable. If FGLRESOURCEPATH is not defined, it searches in the directories defined by the DBPATH
environment variable.

The presentation styles must be defined in a unique .4st file. When loading
a styles file with the `ui.Interface.loadStyles()` method, current styles created
from the default file or from a prior load will be replaced. The styles will not be combined when
loading several files.

The default styles file located in $FGLDIR/lib should not be modified
directly: your changes would be lost if you upgrade the product. Make a copy of the original file
into the program directory of your application, then modify the copied file.
