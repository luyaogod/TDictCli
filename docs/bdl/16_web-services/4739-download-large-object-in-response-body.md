---
title: "Download a large object in the response body"
source: "fgl-topics/c_gws_restful_high_level_download_lob.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Handling file attachments and data transfer > Transfer data in large objects > Download large object in response body"
type: "concept"
---

# Download a large object in the response body

> This example demonstrates how to send a large object in the message response body.

To transfer a large object (LOB) to a client, specify a return parameter using a `TEXT` or `BYTE` data type. The large object is returned in the response body.

## Example: downloading a TEXT object

In this sample REST function there is
an example of returning a large object in the message body. The function has two return
parameters:

- A return paramater defined as `INTEGER` with a
  WSHeader attribute will form a response header.
- A return parameter defined as `TEXT` type will have a value transferred
  in the message body.

A variable (`t`) is defined as `TEXT` and it is located
in memory. The `readFile()` method reads the content of a text file
("file3.txt") into the `TEXT` variable
(`t`). The function's `RETURNS` clause returns the
values.

```
PUBLIC FUNCTION help3()
  ATTRIBUTES (WSGet,
              WSDescription="Download text to the client in a TEXT object",
              WSPath="/help/file3")
  RETURNS ( INTEGER ATTRIBUTES(WSHeader), TEXT )
    DEFINE t TEXT

    LOCATE t IN MEMORY
    CALL t.readFile("file3.txt")

    RETURN 3, t
END FUNCTION
```

![Sample output of HTTP response sending a file in the body](../_images/rest_response_non_multipart.png)

*Output of HTTP response*

The header is given a default name ("`rv0`") at runtime. Use the
`WSName` attribute
to rename the
header.

```
RETURNS (INTEGER ATTRIBUTES(WSHeader, WSName="MyHeader"), TEXT)
```

## Related links

**Related concepts**  

[Download a file in a multipart response](4744-download-file-and-data-response.md "This example demonstrates how to return a file in a form-data type HTTP multipart response while also transferring data of other types in the same message response.")

[Download a file as an attachment to a response](4736-download-file-in-response.md "This example demonstrates how to return a file as an attachment using the WSAttachment attribute")
