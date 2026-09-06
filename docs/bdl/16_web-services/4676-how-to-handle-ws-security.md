---
title: "How to handle WS security"
source: "fgl-topics/c_gws_ssl_security_how_to_001.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to handle WS security"
type: "concept"
---

# How to handle WS security

> The Genero Web Services engine does not entirely manage WS-Security; however, Genero BDL provides XML APIs to help the development of Web Services with security. Exploring the demo Web Service included with the FGLGWS installation will help you handle security in your own SOAP Web service.

These topics describe how to handle Web Services security using the
wssecuritymessage demo, located in
$FGLDIR/demo/WebServices. See the readme.txt file provided
with the demo for information and instructions on how to run the service.

You are encouraged to treat the demo as an example that you can adapt to your needs. It is based
on the WS-Security (WSS) standard. For more information, refer to the WS-Security Policy
documentation.

The demo demonstrates a secure messaging service using a WS-Security policy to exchange messages.
It involves three clients exchanging secured messages. Those clients post and retrieve messages on a
secured server. Each client is identified by a certificate that signs its messages.

The demo assumes that all the clients have sent their public keys to the other clients and to the
server. Those keys are kept in each host's (server or clients) keystore.

As a prerequisite, we recommend that you become familiar with security concepts described in the
[Encryption and Authentication Concepts](4566-encryption-and-authentication.md "A scenario involving a person (Georges) and his bank guides you through the concepts of secured communication, certificates, and certificate authorities.")
page.

> **Important:**
>
> The certificates included in this package are provided for demonstration
> purposes only. As they are distributed with this package, anybody using this product can decrypt the
> messages exchanged. Do NOT use them in production.

## Child topics

- [Server side](4677-server-side.md): Learn the steps the server uses to process messages it sends and receives. Identify the Genero handlers that help implement the security policy on the server side.
- [Client side](4678-client-side.md): From the client side, identify the steps that need to be performed to send and receive secure messages.
- [SOAP security standards](4679-soap-security-standards.md): In this section read about the security policy standards, such as bindings, and the options for signing and encryption.
- [Useful links](4684-useful-links.md): For more information on SOAP message security, follow the links to the policy standards.
