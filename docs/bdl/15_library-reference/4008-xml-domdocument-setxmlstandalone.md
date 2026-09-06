---
title: "xml.DomDocument.setXmlStandalone"
source: "fgl-topics/c_gws_XmlDomDocument_setXmlStandalone.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.setXmlStandalone"
type: "concept"
---

# xml.DomDocument.setXmlStandalone

> Sets the XML standalone attribute in the XML declaration to "yes" or "no" in the XML declaration, or removes the standalone attribute.

## Syntax

```
setXmlStandalone(
   alone INTEGER )
```

1. alone defines a boolean flag.
   - `1` sets the standalone attribute to "yes".
   - `0` sets the standalone attribute to "no".
   - `-1` removes the standalone attribute.

## Usage

This sample sets the standalone attribute to
"no":

```
IMPORT XML

MAIN
  DEFINE doc xml.DomDocument

  LET doc = xml.DomDocument.Create()
  CALL doc.setXmlStandalone(0)

  DISPLAY doc.saveToString()
END MAIN
```

The
output displayed by the above example is `<?xml version="1.0" encoding="UTF-8"
standalone="no"?>`.

This sample sets the standalone attribute to
"yes":

```
IMPORT XML

MAIN
  DEFINE doc xml.DomDocument

  LET doc = xml.DomDocument.Create()
  CALL doc.setXmlStandalone(1)

  DISPLAY doc.saveToString()
END MAIN
```

The
output displayed by the above example is `<?xml version="1.0" encoding="UTF-8"
standalone="yes"?>`.

This sample removes the standalone
attribute:

```
IMPORT XML

MAIN
  DEFINE doc xml.DomDocument

  LET doc = xml.DomDocument.Create()
  CALL doc.setXmlStandalone(-1)

  DISPLAY doc.saveToString()
END MAIN
```

The
output displayed by the above example is `<?xml version="1.0"
encoding="UTF-8"?>`.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[xml.DomDocument.isXmlStandalone](3995-xml-domdocument-isxmlstandalone.md "Checks whether the XML standalone attribute is set in the XML declaration.")
