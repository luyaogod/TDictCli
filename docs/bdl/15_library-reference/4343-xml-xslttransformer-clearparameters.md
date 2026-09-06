---
title: "xml.XsltTransformer.clearParameters"
source: "fgl-topics/c_gws_XmlXsltTransformer_clearParameter.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML transformation classes > The XsltTransformer class > XsltTransformer methods > xml.XsltTransformer.clearParameters"
type: "concept"
---

# xml.XsltTransformer.clearParameters

> Clear all parameters in an instance of an XSLTTransformer object.

## Syntax

```
clearParameters()
```

## Usage

Clears all parameters corresponding to the XSLT entries named by `param`.

For example, this XSLT sample contains two `param`
definitions:

```
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform" >
  
  <xsl:param name="QADIR" />
  <xsl:param name="NODE_NAME"/>
  ... 
  <MyNode><xsl:value-of select="$QADIR" /></MyNode>
  ...
  <xsl:when test="name()=$NODE_NAME">
    <xsl:copy>
      <xsl:copy-of select="@*"/>
    </xsl:copy>
  </xsl:when>

</xsl:stylesheet>
```

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[xml.XsltTransformer.getParameter](4348-xml-xslttransformer-getparameter.md "Get the value of a given parameter in an instance of an XSLTTransformer corresponding to the XSLT entry named by param.")

[xml.XsltTransformer.setParameter](4349-xml-xslttransformer-setparameter.md "Set parameters in an instance of an XSLTTransformer corresponding to the XSLT entries named by param.")
