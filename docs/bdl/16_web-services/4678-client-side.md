---
title: "Client side"
source: "fgl-topics/c_gws_ssl_security_how_to_004.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to handle WS security > Client side"
type: "concept"
---

# Client side

> From the client side, identify the steps that need to be performed to send and receive secure messages.

The client function consists of sending a message and retrieving messages clients have sent to
it.

Before you begin, create the client stub from the WDSL:

- `fglwsdl -domHandler myservice.wsdl`

The client stub references callback handlers:

- `SecureMessageBox_HandleRequest`
- `SecureMessageBox_HandleResponse`
- `SecureMessageBox_HandleResponseFault`

For more details about client SOAP handlers see [Client
stub and handlers](4616-ws-client-stubs-and-handlers.md "To access a remote Web Service, you first must get the WSDL information from the service provider.").

What to do when a message is sent:

1. Sign and encrypt the request for the server (WS-Security)
   - sign with client private key
   - encrypt with server public key
2. Send key information in the request
   - key to identify the sender/client
   - key to identify the recipient/server
   - key used to encrypt the data (usually a symmetric key encrypted
     by the recipient public key)
3. If the message has to be encrypted for the final recipient
   (XML-Security)
   - sign the message
   - encrypt the message

What to do to retrieve messages:

1. Identify the sender and validate the sender (search in keystore)
2. Identify the recipient (should be the server itself)
3. Decrypt the request
4. Check the signature
5. Retrieve messages for the recipient
