---
title: "Utility modules"
source: "fgl-topics/c_fgl_utility_functions.html"
breadcrumb: "Library reference > Utility modules"
type: "concept"
---

# Utility modules

> A utility function is a function provided in a separate library; it is not built in the runtime system.

To use a utility function, declare the module where the function is defined with
the [`IMPORT FGL`](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.")
instruction:

```
IMPORT FGL fgldialog
...
   CALL fgl_winmessage( ... )
```

For backward compatibility, utility functions are also grouped in a 42x library named
`libfgl4js.42x`, which can be linked to your programs.

The 42x library file, 42m modules and 42f forms are located in $FGLDIR/lib.
The sources of the utility functions and form files are provided in the
$FGLDIR/src directory.

## Child topics

- [fgldialog: Common dialog functions](2795-fgldialog-common-dialog-functions.md)
- [fgldbutl: Database utility module](2801-fgldbutl-database-utility-module.md)
- [fglwinexec: Front-end dialogs module](2808-fglwinexec-front-end-dialogs-module.md)
- [VCard: VCF file format module](2815-vcard-vcf-file-format-module.md)
- [fglgallery: Image gallery module](2825-fglgallery-image-gallery-module.md)
- [fglsvgcanvas: SVG drawing module](2841-fglsvgcanvas-svg-drawing-module.md)
- [getopt: Command line options module](2893-getopt-command-line-options-module.md)
