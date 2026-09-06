---
title: "Security bindings"
source: "fgl-topics/c_gws_ssl_security_how_to_006.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to handle WS security > SOAP security standards > Security bindings"
type: "concept"
---

# Security bindings

> Understand the mechanism of bindings that allows the secure exchange of SOAP messages over HTTP. The demo application security policy is referenced to illustrate this.

There are 3 types of security bindings:

- TransportBinding
- SymmetricBinding
- AsymmetricBinding

The demo (located in $FGLDIR/demo/WebServices/wssecuritymessage) uses
Asymmetric binding.

## Asymmetric Binding

This section of the security policy is divided in sub sections:

- InitiatorToken
- RecipientToken
- AlgorithmSuite
- Layout
- Additional assertions

`AsymmetricBinding` is the root node of the security policy for protection
description.

```
<sp:AsymmetricBindingxmlns:sp=
"http://schemas.xmlsoap.org/ws/2005/07/securitypolicy">
```

## InitiatorToken

`InitiatorToken` is the message sender (client)

For example:

```
<sp:InitiatorToken>
  <wsp:Policy>
    <sp:X509Token sp:IncludeToken="http://schemas.xmlsoap.org/ws/
     2005/07/securitypolicy/IncludeToken/AlwaysToRecipient">
      <wsp:Policy>
        <sp:RequireThumbprintReference />
        <sp:WssX509V1Token10 />
      </wsp:Policy>
    </sp:X509Token>
  </wsp:Policy>
</sp:InitiatorToken>
```

The value for the `sp:IncludeToken` attribute is one contiguous string with
no spaces. For this document, it is shown covering two lines.

The token is used for the message signature from initiator to recipient and encryption from
recipient to initiator.

The initiator key is a X509 certificate that is always sent to the recipient.

`sp:IncludeToken` attribute indicates if the token must be included.

`IncludeToken/AlwaysToRecipient` means each requests sent to the recipient must
include the initiator token. But the token is not to be included in messages from recipient
to initiator.

The token must send its Thumbprint Reference.

The token must be of type X509 version 1 as defined in "X509 token profile 1.0".

What needs to be done in BDL is described in [Client Side](4678-client-side.md "From the client side, identify the steps that need to be performed to send and receive secure messages.")
section.

To retrieve the thumbprint reference you can use the API function [xml.CryptoX509.getThumbprintSHA1](../15_library-reference/4234-the-cryptox509-class.md "The xml.CryptoX509 class provides methods to manipulate X509 certificates needed for identification of individual persons, groups or any entities during XML encryption or signature process.")

To create the x509 certificate, use an appropriate tool like openssl.

## RecipientToken

`RecipientToken` is the message receiver (server)

```
<sp:RecipientToken>
  <wsp:Policy>
    <sp:X509Token sp:IncludeToken="http://schemas.xmlsoap.org/
     ws/2005/07/securitypolicy/IncludeToken/Never">
      <wsp:Policy>
        <sp:RequireThumbprintReference />
        <sp:WssX509V3Token10 />
      </wsp:Policy>
    </sp:X509Token>
  </wsp:Policy>
</sp:RecipientToken>
```

The value for the `sp:IncludeToken` attribute is one contiguous string with
no spaces. For this document, it is shown covering two lines.

The token is used for encryption from initiator to recipient, and for the message signature from
recipient to initiator.

The recipient key is a X509 certificate that is never sent to the initiator.

`sp:IncludeToken` attribute indicates if the token must be included.

Use of the `IncludeToken/Never` means the token is never to be included in any
requests between the initiator and the recipient. This is the required and recommended
setting.

Instead the recipient ThumbprintReference is sent.

The token must be of type X509 version 3 as defined in "X509 token profile 1.0"

What needs to be done in BDL is described in [Server Side](4677-server-side.md "Learn the steps the server uses to process messages it sends and receives. Identify the Genero handlers that help implement the security policy on the server side.")
section. To retrieve the thumbprint reference you can use the API function [xml.CryptoX509.getThumbprintSHA1](../15_library-reference/4234-the-cryptox509-class.md "The xml.CryptoX509 class provides methods to manipulate X509 certificates needed for identification of individual persons, groups or any entities during XML encryption or signature process.").
To create the appropriate certificate use an appropriate tool like openssl.

## AlgorithmSuite

`AlgorithmSuite` specifies which algorithm is used to encrypt
the data.

```
<sp:AlgorithmSuite>
  <wsp:Policy>
    <sp:TripleDesRsa15 />
  </wsp:Policy>
</sp:AlgorithmSuite>
```

TripleDesRsa15 refers to key [http://www.w3.org/2001/04/xmlenc#tripledes-cbc](../15_library-reference/4234-the-cryptox509-class.md "The xml.CryptoX509 class provides methods to manipulate X509 certificates needed for identification of individual persons, groups or any entities during XML encryption or signature process.").

## Layout

`Layout` describes the way information is added to the message header.

```
<sp:Layout>
  <wsp:Policy>
    <sp:Strict />
  </wsp:Policy>
</sp:Layout>
```

For example, with `Strict` layout, tokens that are included in the message must be
declared before use. For more details on the rules to follow see the [WS-SecurityPolicy](http://docs.oasis-open.org/ws-sx/ws-securitypolicy/200702/ws-securitypolicy-1.2-spec-os.pdf) specifications section 7.7.

## Additional Assertions

`PartsToSign`

```
<sp:OnlySignEntireHeadersAndBody />
```

The assertion means if there is any signature on the header or the body it applies to the entire
header and the entire body not to their child element.
