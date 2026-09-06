---
title: "Example: Using xml.XsltTransformer methods"
source: "fgl-topics/c_gws_XmlXsltTransformer_example.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML transformation classes > The XsltTransformer class > Example"
type: "concept"
---

# Example: Using xml.XsltTransformer methods

> This Genero application provides a working example using methods from the XsltTransformer class.

```
IMPORT XML

MAIN
  DEFINE ok BOOLEAN
  IF num_args()!=3 THEN
    DISPLAY "Usage : DoXslp <stylesheet> <source> <result>"
    EXIT PROGRAM 1
  ELSE
    LET ok = RunXSLP(arg_val(1),arg_val(2),arg_val(3))
    IF NOT ok THEN
      DISPLAY "Error: failed"
      EXIT PROGRAM 1
    ELSE
      DISPLAY "Done"
      EXIT PROGRAM
    END IF
  END IF
END MAIN

FUNCTION RunXSLP(style,src,ret)
  DEFINE style,src,ret  STRING
  DEFINE ind            INTEGER
  DEFINE xslt           xml.XSLTTransformer
  DEFINE styleSheet     xml.DomDocument
  DEFINE source         xml.DomDocument
  DEFINE result         xml.DomDocument

  # Load StyleSheet
  TRY
    LET styleSheet = xml.DomDocument.Create()
    CALL styleSheet.load(style)
  CATCH
    DISPLAY "Error: unable to load stylesheet",style
    RETURN FALSE
  END TRY
  
  # Create XSLT transformer
  TRY
    LET xslt = xml.XsltTransformer.CreateFromDocument(styleSheet)
    FOR ind=1 TO xslt.getErrorsCount()
      DISPLAY "StyleSheet error #"||ind||" : ",xslt.getErrorDescription(ind)
    END FOR
  CATCH
    DISPLAY "Error : unable to create XSLT transformer from ",styleSheet
    RETURN FALSE
  END TRY

  # Load Source 
  TRY
    LET source = xml.DomDocument.Create()
    CALL source.load(src)
  CATCH
    DISPLAY "Error : unable to load Source from ",src
    RETURN FALSE
  END TRY
  
  # Execute XSLT 
  TRY
    LET result = xslt.doTransform(source)
    FOR ind=1 TO xslt.getErrorsCount()
      DISPLAY "Error #"||ind||" : ",xslt.getErrorDescription(ind)
    END FOR    
  CATCH
    DISPLAY "Error : unable to apply XSLT stylesheet"
    FOR ind=1 TO xslt.getErrorsCount()
      DISPLAY "Fatal Error #"||ind||" : ",xslt.getErrorDescription(ind)
    END FOR
    RETURN FALSE
  END TRY
  
  # Save resulting 
  TRY
    CALL result.save(ret)
  CATCH
    DISPLAY "Error : unable to save result on disk"
    RETURN FALSE
  END TRY
  
  RETURN TRUE
END FUNCTION
```
