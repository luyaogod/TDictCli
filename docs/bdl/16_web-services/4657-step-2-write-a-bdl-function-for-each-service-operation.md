---
title: "Step 2: Write a BDL function for each service operation"
source: "fgl-topics/c_gws_server_tutorial_006.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Writing a Web server application > Example 1: Writing the entire server application > Step 2: Write a BDL function for each service operation"
type: "concept"
---

# Step 2: Write a BDL function for each service operation

> Each function defines an operation of the service.

You will need to write a function to implement each operation, using the [input and output](4656-step-1-define-input-and-output-records.md "Define records for the input and output messages of the Web function.") records.

If you code the operations of the service in a separate module to the service creation, you
will need to import the service module to reference the input and output records.

To implement your **Add** operation:

```
#User public function
PUBLIC FUNCTION add()
   LET add_out.r = add_in.a + add_in.b
END FUNCTION
```

## Related links

**Related concepts**  

[Step 3: Create the service and operations](4658-step-3-create-the-service-and-operations.md "Describes how you provide your Web service and its operations to users who can access it on the net.")
