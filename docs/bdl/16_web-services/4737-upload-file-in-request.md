---
title: "Example: upload a file as attachment in request body"
source: "fgl-topics/c_gws_restful_high_level_upload_file_example.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Handling file attachments and data transfer > Attach files with WSAttachment and WSMedia > Upload file in request"
type: "concept"
---

# Example: upload a file as attachment in request body

> Upload files by adding WSAttachment to an input parameter.

In this sample REST function, there is one input parameter:

- The "fname" parameter is defined as type `STRING` with a [WSAttachment](4839-wsattachment.md "Defines file attachments in the REST message.") attribute. The
  `WSAttachment` pattern "`[A-Za-z0-9_-]+.(jpg|jpeg|png)`" matches
  filenames that begin with one or more letters, digits, underscores, or hyphens, followed by a
  literal dot and a lowercase extension of jpg, jpeg, or
  png (case-sensitive). Example matches: image-01.jpg,
  photo\_2.jpeg. The REST engine treats the parameter value as a path to a file to
  be attached.

  Use a regex to keep filenames predictable, block unsafe input (for example,
  "../"), and limit allowed extensions.
- The `WSMedia` attribute
  handles the data format for images. The wildcard media type (`image/*`) in
  the [WSMedia](4841-wsmedia.md "Defines the supported media (MIME) types for a parameter or return value.") attribute lets the service
  accept or return any image format. If the file can be of any type (not only images), omit the
  `WSMedia` attribute from `WSAttachment`.

The GWS stores the uploaded file in a temporary directory and returns its absolute path in the
input parameter fname. The file will be removed when the call ends unless you
save it to disk.

Compute the destination path and move the file:

- Set the destination path: `LET dest = os.Path.join("/data/images",
  os.Path.baseName(fname))`

  ([os.Path.baseName](../15_library-reference/3700-os-path-basename.md "Returns the last element of a path.") strips any
  directory components)
- Move the uploaded file from its temporary location to dest: `LET ok =
  os.Path.rename(fname, dest)`

```
IMPORT os

PUBLIC FUNCTION uploadImage(
    fname STRING ATTRIBUTES(WSAttachment = "[a-zA-Z0-9_\-]+\.(jpg|jpeg|png)",
        WSMedia = "image/*",
        WSDescription = "Image file to upload (jpg, jpeg or png)"))
    ATTRIBUTES(WSPost,
        WSPath = "/upload/image",
        WSDescription = "Upload an image file with validation of filename")
    RETURNS STRING ATTRIBUTES(WSMedia = "application/json")
 
    DEFINE dest STRING
    DEFINE ok INTEGER
 
    -- Store the image in /data/images with its secure name
    LET dest = os.Path.join("/data/images", os.Path.baseName(fname))
 
    -- Move the temporary file
    LET ok = os.Path.rename(fname, dest)
 
    RETURN SFMT('{"status":"ok","file":"%1"}', dest)
 
END FUNCTION
```

![The image shows the output as the attachment is sent in the message body](../_images/rest_request_attachement.png)

*Output of HTTP request to upload file*

In Figure 1 the
image is sent as an attachment in the message body.

## Related links

**Related concepts**  

[Upload a large object in the request body](4740-upload-large-object-in-request-body.md "This example demonstrates how to upload a large object in the message request body.")

[Example: upload a file in a multipart request](4745-upload-file-and-data-request.md "Shows an example of a method you might use from a client to upload an image along with other data in a function using a form-data type HTTP multipart request.")
