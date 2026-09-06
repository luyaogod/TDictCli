---
title: "Step 3: Create the service and operations"
source: "fgl-topics/c_gws_server_tutorial_007.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Writing a Web server application > Example 1: Writing the entire server application > Step 3: Create the service and operations"
type: "concept"
---

# Step 3: Create the service and operations

> Describes how you provide your Web service and its operations to users who can access it on the net.

The Genero Web Services library (`com`)
provides classes and methods that allow you to use Genero BDL to configure a Web Service and its
operations.

- [WebService class](../15_library-reference/3754-the-webservice-class.md "The com.WebService class provides an interface to create and manage Genero Web Services.") - this is a container for web
  operations.
- [WebOperation class](../15_library-reference/3770-the-weboperation-class.md "The com.WebOperation class provides an interface to create and manage the operations of a Genero Web Service.") - describes the
  operation.

## Define variables for the WebService and WebOperation objects

```
FUNCTION createservice()
  DEFINE serv  com.WebService    # A WebService
  DEFINE op    com.WebOperation  # Operation of a WebService
```

## Choose a Namespace

[XML](4486-introduction-to-web-services.md "Web services are a standard way of communicating between applications over an intranet or Internet.") uses namespaces to group element
and attribute definitions, and to avoid conflicting names. In practice, a namespace must be a unique
identifier (URI: Uniform Resource Identifier). If you do not know the unique identifier to use, your
company's Web site domain name is guaranteed to be unique (such as "www.mycompany.com"); then,
append any string.

Examples of valid namespaces for the fictional "My Company" company:

- "http://www.mycompany.com/MyServices"
- "http://www.mycompany.com/any\_string"

Another option (for testing only) is to use the temporary namespace "http://tempuri.org/".

## Create the WebService object

Call the constructor method of the `WebService` class. The parameters are:

1. Service name
2. Valid namespace

This example uses the temporary namespace and creates a service named
"MyCalculator".

```
LET serv =
  com.WebService.CreateWebService("MyCalculator", "http://tempuri.org/webservices")
```

## Create the WebOperation object

A `WebService` object can have multiple operations. The operations can be created
in [RPC or Document](4499-web-services-style-options.md "Information on Web services Style options available for SOAP Genero Web services. There is no style concept in REST.") style by calling the corresponding
constructor method of the `WebOperation` class. The parameters are:

1. the name of the BDL function that is executed to process the XML operation
2. the name you wish to assign to the XML operation
3. the [input](4656-step-1-define-input-and-output-records.md "Define records for the input and output messages of the Web function.") record defining the input parameters of the
   operation (or NULL if there is none)
4. the [output](4656-step-1-define-input-and-output-records.md "Define records for the input and output messages of the Web function.") record defining the output
   parameters of the operation (or NULL if there is none)

To create the operation for the previously defined **add** function in RPC style:

```
LET op = com.WebOperation.CreateRPCStyle("add", "Add", add_in, add_out)
```

To create the operation for the previously defined **add** function in Document style:

```
LET op = com.WebOperation.CreateDOCStyle("add", "Add", add_in, add_out)
```

Mixing RPC style and Document style operations in the same service is not recommended, as it is not WS-I compliant. See [Web Services Styles](4673-choosing-a-web-services-style.md "Genero Web Services contains style options for creating SOAP Web services. Your choice is dependent on the type of service, (Document or RPC), and the encoding mechanism (literal or encoded) required.") for additional information about styles.

The rest of the code in your application is the same, regardless of the Web Services style that you have chosen.

## Publish the operation

Once an operation is defined, it must be associated with its corresponding WebService (the
operation must be published). The `publishOperation` method of the [WebService](../15_library-reference/3754-the-webservice-class.md "The com.WebService class provides an interface to create and manage Genero Web Services.") object has the following parameters:

- The `WebOperation` to be published.
- A string to identify the operation if several operations have the same name; if this is NULL,
  the default value is an empty string.

For example, to publish the **Add** operation of the **Calculator** service, which was defined as **op**:

```
CALL serv.publishOperation(op,NULL)
```

## Related links

**Related concepts**  

[Step 4: Register the service](4659-step-4-register-the-service.md "Register the service with the Genero Web Services (GWS) server.")
