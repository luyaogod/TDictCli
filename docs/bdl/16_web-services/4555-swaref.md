---
title: "swaRef (SOAP with attachments using wsi:swaRef)"
source: "fgl-topics/c_gws_swaref.html"
breadcrumb: "Web services > Concepts > swaRef"
type: "concept"
---

# swaRef (SOAP with attachments using wsi:swaRef)

> swaRef is a specific way for sending and receiving attachments in SOAP. It is used when you have to transfer files as attachment and locate on disk.

swaRef refers to "Soap with attachments using the wsi:swaRef XML type from WS-I". It
is a specific case of the SoapWithAttachment specification.

> **Warning:**
>
> Java handles swaRef automatically. Dot Net does not.

With swaRef, you can set an attribute with a dedicated value of "swaRef" on a STRING. When set,
a Web service automatically sends the file in that string as an HTTP part or receives it as an HTTP
part.

In the following example, a `sendAttachment` operation computes the MD5 of a file.
The program sets `XMLOptimizedContent="swaRef"` on the `DataIn`
string and the SOAP engine carries it over the wire as an HTTP
part.

```
# Request Type
TYPE tsendAttachment RECORD ATTRIBUTES( XMLSequence, XSTypeName="sendAttachment",
                                        XSTypeNamespace="http://mywebsite.com/services/swa")
  Name STRING ATTRIBUTES( XMLName="Name", XMLOptional ),
  DataIn STRING ATTRIBUTES( XMLOptimizedContent="swaRef",
                            XMLAttribute, XMLName="DataIn", XMLOptional )
END RECORD

# Response Type
TYPE tsendAttachmentResponse RECORD ATTRIBUTES( XMLSequence, XSTypeName="sendAttachmentResponse",
                                                XSTypeNamespace="http://mywebsite.com/services/swa" )
  return STRING ATTRIBUTES( XMLName="return", XMLOptional )
END RECORD

# VARIABLE : sendAttachment
DEFINE sendAttachment tsendAttachment ATTRIBUTES( XMLName="sendAttachment",
                                                  XMLNamespace="http://mywebsite.com/services/swa" )

# VARIABLE : sendAttachmentResponse
DEFINE sendAttachmentResponse tsendAttachmentResponse ATTRIBUTES( XMLName="sendAttachmentResponse",
                                   XMLNamespace="http://mywebsite.com/services/swa" )
...
    # Publish Operation : sendAttachment
    LET operation = com.WebOperation.CreateDOCStyle( "sendAttachment", "sendAttachment",
                                                     sendAttachment, sendAttachmentResponse )
    CALL service.publishOperation(operation,"")
...
```

## Related links

**Related concepts**  

[XMLOptimizedContent](5045-xmloptimizedcontent.md "Set on STRING or BYTE data type so that such string content represents a file on disk to be transmitted as base64 binary in SOAP via HTTP attachment.")
