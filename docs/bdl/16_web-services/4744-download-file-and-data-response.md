---
title: "Download a file in a multipart response"
source: "fgl-topics/c_gws_restful_high_level_multipart_response_body.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Multipart requests or responses > Download file and data (response)"
type: "concept"
---

# Download a file in a multipart response

> This example demonstrates how to return a file in a form-data type HTTP multipart response while also transferring data of other types in the same message response.

## Example multipart response

In this sample REST function there is an example of a multipart response. The
`RETURNS` clause of the function has three return values: an
`INTEGER,``STRING`, and `TEXT`. The
`INTEGER` is defined with the attribute [WSHeader](4838-wsheader.md "Defines a custom HTTP header for a parameter or return value."), so its value is sent in a
HTTP header. The `STRING` and `TEXT` type have no
`WSHeader` attributes, so therefore values are sent in the response body.

The `TEXT` return value is defined with the `WSName` attribute to identify
it by the name "txt\_file" instead of the DVM default name.

- A variable (t) is defined as type `TEXT` and located
  in memory.
- The `readFile()` method reads the contents of a text file
  ("README.txt") into variable t.

```
PUBLIC FUNCTION help2()
  ATTRIBUTES (WSGet,
              WSPath = "/help2")
  RETURNS (INTEGER ATTRIBUTES(WSHeader),
           STRING,
           TEXT ATTRIBUTES(WSName = "txt_file") )
    DEFINE t TEXT
    LOCATE t IN MEMORY
    CALL t.readFile("README.txt")
    RETURN 3, "hello world", t
END FUNCTION
```

## Output of multipart HTTP response

In the output the header is given the default name, "rv0", at runtime.

In the multipart response parts are combined into one or more sets of data in the body.
Parts are separated by boundaries, "`--MIME_boundary`" in the output example.
The text file part is identified by the name ("txt\_file")

![Sample output of multipart HTTP response](../_images/rest_response_multipart.png)

## Related links

**Related concepts**  

[Download a file as an attachment to a response](4736-download-file-in-response.md "This example demonstrates how to return a file as an attachment using the WSAttachment attribute")

[Download a large object in the response body](4739-download-large-object-in-response-body.md "This example demonstrates how to send a large object in the message response body.")
