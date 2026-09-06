---
title: "Upload a large object in the request body"
source: "fgl-topics/c_gws_restful_high_level_upload_data.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Handling file attachments and data transfer > Transfer data in large objects > Upload large object in request body"
type: "concept"
---

# Upload a large object in the request body

> This example demonstrates how to upload a large object in the message request body.

To upload a large object (LOB) to a Web service server, you must specify an input parameter using
a `TEXT` or `BYTE` data type. Any binary file contents (image,
PDF, zip, and so on) may be uploaded using this method.

## Example: Uploading a BYTE object and writing to a file

In this sample REST function there is an example of sending a large object. The function has one
input paramater "b" defined as type `BYTE`, which will
transfer a large object in the message body.

The `writeFile()` method writes the contents of the variable
b to a file ("help.pdf"). If the write file
operation fails, this is handled in the `TRY/CATCH` block.

```
IMPORT os

PUBLIC FUNCTION UploadByteToFile(b BYTE)
  ATTRIBUTES (WSPost,
              WSDescription="Upload BYTE and write to file on server",
              WSPath="/help/pdf")
  RETURNS STRING
    CONSTANT filename = "help.pdf"
    DEFINE ret STRING
    TRY   
       CALL b.writeFile(filename)
       LET ret = SFMT("Data received, saved to file: %1",filename)
    CATCH
       LET ret = SFMT(" Error getting data, not saved to file: %1",filename)
    END TRY 
    RETURN ret
END FUNCTION
```

## Example: Client app calling the function

In this sample there is an example of the client app. The client must import the stub file
(clientStub) in order to call the service:

- A variable (`b`) is defined as type `BYTE` and is
  located in memory.
- The `readFile()` method reads the contents of a PDF file
  ("help.pdf") into variable `b`.

In the call to the stub function `UploadByteToFile`
`b` is set as the input parameter. In the `CASE` statement,
the variable (`wsstatus`) is set to handle errors in the return of the
function call to the service. For more information on calling a Web service, see [Create the client application](4785-create-the-client-application.md "Prepare the client application that uses the generated stub to call the REST service.").

```
# clientApp.4gl

IMPORT util
-- import the service stub file
IMPORT FGL clientStub

MAIN
  DEFINE wsstatus INTEGER
  DEFINE retS STRING
 
  DEFINE b BYTE
  LOCATE b IN MEMORY
  CALL b.readFile("help.pdf")

  CALL clientStub.UploadByteToFile(b) RETURNING wsstatus, retS
  CASE wsstatus
    WHEN clientStub.C_SUCCESS
      DISPLAY SFMT("Got %1 ", retS)
    OTHERWISE
      DISPLAY "wsstatus is : ", wsstatus
      DISPLAY "Error in getting file uploaded"
  END CASE

END MAIN
```

## Related links

**Related concepts**  

[Example: upload a file in a multipart request](4745-upload-file-and-data-request.md "Shows an example of a method you might use from a client to upload an image along with other data in a function using a form-data type HTTP multipart request.")

[Example: upload a file as attachment in request body](4737-upload-file-in-request.md "Upload files by adding WSAttachment to an input parameter.")
