---
title: "Prerequisites to produce documentation"
source: "fgl-topics/c_fgl_AutoDoc_003.html"
breadcrumb: "Programming tools > Source documentation > Prerequisites to produce documentation"
type: "concept"
---

# Prerequisites to produce documentation

> This topic lists the requirements to generate source documentation.

To generate the HTML pages, fglcomp first generates .xa
files which must be converted to .html files. The conversion from
.xa to .html is done with an XSLT processor using the
.xsl style sheet files provided in $FGLDIR/lib/fgldoc/

You must have an XSLT processor installed on the machine where
the documentation is generated.

- On UNIX™, fglcomp runs the
  $FGLDIR/lib/fgldoc/Transform.sh script to convert .xa files
  to .html files. Therefore you need the xsltproc command line
  XSLT processor (from the libxml package).
- On Windows®, fglcomp runs the
  $FGLDIR\lib\fgldoc\Transform.js script to convert .xa files
  to .html files. To run the Transform.js script, you must
  have cscript.exe installed with the Microsoft™.XMLDOM class (this is the case on recent Windows versions).

If the default result of the transformation does not fit your needs, the style sheets provided in
$FGLDIR/lib/fgldoc can be adapted to generate different HTML files.
