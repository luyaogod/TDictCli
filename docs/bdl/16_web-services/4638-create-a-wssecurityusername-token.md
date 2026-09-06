---
title: "How to create a SOAP WSSecurityUserName token"
source: "fgl-topics/c_gws_soap_wssecurityusername_token.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Create a WSSecurityUserName token"
type: "concept"
---

# How to create a SOAP WSSecurityUserName token

> Creating a WSSecurityUserName token requires you to generate a client stub using the --domHandler option and using methods from the securityHelper module to create a WSSecurityUserName token.

## Creating the client stub

Use the `-domHandler` option when creating a client stub with
fglwsdl.

## Code your application

You will use two functions contained in the securityHelper module. In your
module source, add

```
IMPORT FGL securityHelper
```

to access these functions.

The source module used to create the p-code module for securityHelper
can be found at:
$FGLDIR/demo/WebServices/wssecuritymessage/common/securityHelper.4gl.
Reference this file to get full specifics regarding these two functions.

In the request callback, create the WSSecurity header using
`BuildSOAP11WSSecurity()`. You pass the header node coming from the callback; it
returns the Web Service Security node that is automatically added as a child of the SOAP header.

Create the WSUserName token using `BuildWSSUsernameToken()`. You pass the node
returned by `BuildSOAP11WSSecurity()` in the previous step, along with a unique
identifier to reference that token, a username and password, a unique nonce value, and the date the
timestamp was created; it returns the WS-UsernameToken node. The token node is automatically added
as child of the WSSecurity header.

Return `TRUE` from the callback function to inform GWS to continue the standard
process. GWS will sent the SOAP request with the WSSecurity header inside.

## Example

This example is a simple example of how to create a SOAP WSSecurityUserName token. Values would
be required to replace the variable placeholders
my-username,my-password,unique-nonce-value,
date-timestamp-created in this example.

For a more complex example, examine the wssecuritymessage demo located at
$FGLDIR/demo/WebServices/wssecuritymessage/.

```
IMPORT FGL securityHelper

....

FUNCTION MyService_HandleRequest(operation,doc,header,body)
  DEFINE operation STRING          -- Operation name of the request to be modified.
  DEFINE doc       xml.DomDocument -- Entire XML document of the request
  DEFINE header    xml.DomNode     -- XML node of the SOAP header of the request
  DEFINE body      xml.DomNode     -- XML node of the SOAP body of the request
  DEFINE wss       xml.DomNode
  DEFINE node      xml.DomNode
  DEFINE id        STRING

  LET wss = BuildSOAP11WSSecurity(header,TRUE)
  LET id = SFMT("UsernameToken-%1",security.RandomGenerator.CreateRandomString(8))
  LET node = BuildWSSUsernameToken(wss,id,"<my-username>","<my-password>",
                                   "<unique-nonce-value>","<date-timestamp-created>")
  RETURN TRUE
END FUNCTION
```
