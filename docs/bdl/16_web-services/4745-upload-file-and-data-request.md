---
title: "Example: upload a file in a multipart request"
source: "fgl-topics/c_gws_restful_high_level_multipart_request.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Multipart requests or responses > Upload file and data (request)"
type: "concept"
---

# Example: upload a file in a multipart request

> Shows an example of a method you might use from a client to upload an image along with other data in a function using a form-data type HTTP multipart request.

## Example multipart request

In this sample REST function there is an example of a multipart request. The function
has three input parameters:

- "id" is defined as type `INTEGER`.
- "img" is defined as type `STRING` with a
  `WSAttachment` attribute. The REST engine treats the parameter value as a path to a
  file to be attached. The `WSMedia` attribute handles the data format for images. The wildcard media type (`image/*`) in
  the [WSMedia](4841-wsmedia.md "Defines the supported media (MIME) types for a parameter or return value.") attribute lets the service
  accept or return any image format. If the file can be of any type (not only images), omit the
  `WSMedia` attribute from `WSAttachment`.
- An input parameter "submit" is defined as type
  `STRING`.

The GWS stores the file in a temporary directory on the
server, and returns the absolute path to the file in the input parameter
("img"). When the call to the function ends, the file will be
removed unless it is saved to a location on disk.

The dest variable is set to create a path to
the current directory, to which the filename is added with the [os.Path.baseName](../15_library-reference/3700-os-path-basename.md "Returns the last element of a path.") instruction. The [os.Path.rename](../15_library-reference/3737-os-path-rename.md "Renames a file or a directory.") instruction moves the file from the temporary directory to the current directory.

[WSThrows](4829-wsthrows.md "Defines the list of error codes the REST function may return.") – is set to handle
errors. If the `id` is missing or there is a problem renaming and
moving the file, the `message` field of the
`userError` variable is set, and a call to `SetRestError()` returns the message defined in
`WSThrows` for the error.

```
IMPORT os
IMPORT com

PUBLIC DEFINE userError RECORD ATTRIBUTES(WSError = "User error")
  message STRING
END RECORD

PUBLIC FUNCTION fetchFiles(
    id INTEGER,
    img STRING ATTRIBUTES(WSAttachment,WSMedia = "image/*"),
    submit STRING
    )
  ATTRIBUTES (WSPost,
              WSPath = "/files/fetch",
              WSDescription = "upload image file to the server from browser",
              WSThrows = "400:@userError")
  RETURNS STRING ATTRIBUTES(WSMedia = "text/html")
    DEFINE ret, dest STRING
    DEFINE ok INTEGER
    IF id == 0 THEN
        LET userError.message = "Missing id"
        LET ret = SFMT("<HTML><body><h1>Must have an ID %1</h1></body></HTML>",id)
        CALL com.WebServiceEngine.SetRestError(400,userError)
    ELSE   
        TRY 
            LET dest = os.Path.baseName(img)
            LET dest = os.Path.pwd()||"/"||dest
            LET ok = os.Path.rename(img,dest)
            LET ret = SFMT("Got image with ID %1",id)
        CATCH   
            LET userError.message = "Error uploading file"
            LET ret = SFMT("Error uploading file %1",id)
            CALL com.WebServiceEngine.SetRestError(400,userError)  
        END TRY
    END IF
    RETURN ret
END FUNCTION
```

## Create a client side HTML page to upload file

The "fetchFiles" function needs to be called by the client to upload a file. This example shows
you how to create a simple HTML page to select the file to upload. The HTML page
contains a form with two input fields and one submit button, according to the function's input parameters.

1. In a text editor create the form from the code sample. Remember to change
   the `action` tag value to the URL of your server.
2. Save the file to a .html file, for example,
   upload.html, and place it on the client side.
3. Open the file locally in a browser and select a file to upload. Ensure the
   service is running.

```
<!DOCTYPE html>
<HTML>
   <form name = "Upload" method = "post" 
    action = "http://localhost:8090/Myservice/files/fetch" 
    enctype = "multipart/form-data">
    <div>
        <label for = "name">ID:</label>
        <input type = "text" name = "id"/>
    </div>
    <div>
        <label>Picture:</label>
        <input type = "file" name = "img" accept = "image/png, image/jpeg" />
    </div>
    
    <input type = "submit" name = "submit" value="Send"/>
    
    </form>
</HTML>
```

![Sample output of multipart HTTP request](../_images/rest_request_multipart.png)

*Sample output of multipart HTTP request*

## Related links

**Related concepts**  

[Example: upload a file as attachment in request body](4737-upload-file-in-request.md "Upload files by adding WSAttachment to an input parameter.")

[Upload a large object in the request body](4740-upload-large-object-in-request-body.md "This example demonstrates how to upload a large object in the message request body.")

[Set data format with WSMedia](4741-set-data-format-with-wsmedia.md "It is important to set the correct MIME type for a Web service request or response. You can specify the data format via the WSMedia attribute.")
