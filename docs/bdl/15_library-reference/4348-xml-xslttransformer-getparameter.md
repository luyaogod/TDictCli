---
title: "xml.XsltTransformer.getParameter"
source: "fgl-topics/c_gws_XmlXsltTransformer_getParameter.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML transformation classes > The XsltTransformer class > XsltTransformer methods > xml.XsltTransformer.getParameter"
type: "concept"
---

# xml.XsltTransformer.getParameter

> Get the value of a given parameter in an instance of an XSLTTransformer corresponding to the XSLT entry named by param.

## Syntax

```
getParameter(
   param STRING)
  RETURNS STRING
```

1. param defines the name of the
   parameter.

## Usage

Get parameters corresponding to the XSLT entries named by `param`.

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

To get the parameter
`QADIR`:

```
DEFINE dir STRING
SET dir=xslt.getParameter("QADIR")
```

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

To see these methods in a working example, see [Example: Using xml.XsltTransformer methods](4350-example.md "This Genero application provides a working example using methods from the XsltTransformer class.").

## Related links

**Related concepts**  

[xml.XsltTransformer.setParameter](4349-xml-xslttransformer-setparameter.md "Set parameters in an instance of an XSLTTransformer corresponding to the XSLT entries named by param.")

[xml.XsltTransformer.clearParameters](4343-xml-xslttransformer-clearparameters.md "Clear all parameters in an instance of an XSLTTransformer object.")
