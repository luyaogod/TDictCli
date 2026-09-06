---
title: "Server side"
source: "fgl-topics/c_gws_ssl_security_how_to_003.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to handle WS security > Server side"
type: "concept"
---

# Server side

> Learn the steps the server uses to process messages it sends and receives. Identify the Genero handlers that help implement the security policy on the server side.

We provide 3 handlers to handle WS Security:

- Method [com.WebService.registerWSDLHandler()](../15_library-reference/3766-com-webservice-registerwsdlhandler.md "Registers the function to be executed when a WSDL is generated.") to modify the WSDL to add WS policy.
- Method [com.WebService.registerInputRequestHandler()](../15_library-reference/3763-com-webservice-registerinputrequesthandler.md "Registers the function to be executed on incoming SOAP requests.") to handle WS Security in an
  incoming request
- Method [com.WebService.registerOutputRequestHandler()](../15_library-reference/3765-com-webservice-registeroutputrequesthandler.md "Registers the function to be executed just before the SOAP response is forwarded to the client.") to handle WS Security in an
  outgoing request

In this demo (located in $FGLDIR/demo/WebServices/wssecuritymessage), a
received message is processed:

1. Identify the sender and validate the sender (search in keystore)
2. Decrypt the symmetric key with the server private key
3. Decrypt the body
4. Check the signature with the sender public key
5. Store the message in the box (thanks to the "To" field, "subject" and "message")
6. Create the outgoing message
7. Sign the outgoing message
8. Encrypt the outgoing message with a generated symmetric key. This symmetric key is then
   encrypted with the client public key.
