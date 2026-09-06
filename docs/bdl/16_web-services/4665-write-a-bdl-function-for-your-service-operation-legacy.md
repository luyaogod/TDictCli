---
title: "Write a BDL function for your service operation (legacy)"
source: "fgl-topics/c_gws_server_tutorial_012_legacy.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Writing a Web server application > Example 2: Writing a server using third-party WSDL (the fglwsdl tool) > Step 2: Write a BDL function for your service operation > Write a BDL function for your service operation (legacy)"
type: "concept"
---

# Write a BDL function for your service operation (legacy)

> Write functions that use the generated legacy code (Genero 3.20 or prior) for the server stub of the WSDL. This allows you to create your own version of the function.

Using the information from the files generated in [Get the WSDL description and generate legacy files](4663-get-the-wsdl-description-and-generate-legacy-files.md "Use the fglwsdl tool legacy option to generate legacy code (Genero 3.20 or prior) for the server stub from a WSDL."), the **Add** operation from [Example 1: Writing the entire server application](4655-example-1-writing-the-entire-server-application.md "Design a simple Web service.") is rewritten to have different
functionality but to still be compatible with the WSDL description of the operation.

This step accomplishes the same thing as [Step 2: Write a BDL function for each service operation](4657-step-2-write-a-bdl-function-for-each-service-operation.md "Each function defines an operation of the service.") in Example 1. In this version of
the add operation, the sum of the two numbers in the input record is increased by 100.

If you have [generated stub
files](4618-generate-the-stub-files-legacy.md "Use the fglwsdl tool to generate legacy client stub files (.inc and .4gl) compatible with apps created with Genero 3.20 or prior.") for compatibility with your GWS client app, you have a reference in the [`GLOBALS`](4601-step-2-import-the-stub-file-legacy.md) statement at the
top of the .4gl module that links the stub file.

```
# my_function.4gl                         -- file containing the function
                                          -- definition
IMPORT com                                -- import the Web Services library
GLOBALS "example1Service.inc"             -- use the generated globals file
#User Public Functions
FUNCTION add()                            -- new version of the add function
  LET AddResponse.r = (Add.a + Add.b)+ 100  -- the global input and output 
                                          -- records are used
END FUNCTION
```

## Related links

**Related concepts**  

[Step 3: Create service, start server and process requests](4666-step-3-create-service-start-server-and-process-requests.md "Code to start the Genero Web Services (GWS) Server.")
