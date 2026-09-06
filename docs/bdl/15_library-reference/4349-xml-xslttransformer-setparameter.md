---
title: "xml.XsltTransformer.setParameter"
source: "fgl-topics/c_gws_XmlXsltTransformer_setParameter.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML transformation classes > The XsltTransformer class > XsltTransformer methods > xml.XsltTransformer.setParameter"
type: "concept"
---

# xml.XsltTransformer.setParameter

> Set parameters in an instance of an XSLTTransformer corresponding to the XSLT entries named by param.

## Syntax

```
setParameter(
   param STRING,
   value STRING )
```

1. param defines the name of the
   parameter.
2. value specifies the value to set the parameter.

## Usage

Set parameters corresponding to the XSLT entries named by `param`.

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

To set the parameter
`QADIR`:

```
CALL xslt.setParameter("QADIR","'"||"/work/tmp"||"'") # Literal
```

> **Note:**
>
> If a parameter is a literal, you must quote it before calling the
> `setParameter()`

To set the parameter
`NODE_NAME`:

```
CALL xslt.setParameter("NODE_NAME","BOOK")
```

If the parameter is already set, the previous value will be replaced by the new one.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[xml.XsltTransformer.getParameter](4348-xml-xslttransformer-getparameter.md "Get the value of a given parameter in an instance of an XSLTTransformer corresponding to the XSLT entry named by param.")

[xml.XsltTransformer.clearParameters](4343-xml-xslttransformer-clearparameters.md "Clear all parameters in an instance of an XSLTTransformer object.")
