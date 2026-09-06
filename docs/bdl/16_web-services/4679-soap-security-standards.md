---
title: "SOAP security standards"
source: "fgl-topics/c_gws_ssl_security_how_to_005.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to handle WS security > SOAP security standards"
type: "concept"
---

# SOAP security standards

> In this section read about the security policy standards, such as bindings, and the options for signing and encryption.

A Web service security policy requirements are included in the Web Services Description Language
(WSDL) file. The policy is divided into these sections:

- Security bindings
- Security options
- Signed parts
- Encrypted parts

The following defines a
policy:

```
<wsp:Policy xmlns:wsp="http://schemas.xmlsoap.org/ws/2004/09/policy ... />
```

The following section includes policy rules.

```
<wsp:ExactlyOne>
```

Only one assertion is fulfilled.

```
<wsp:All>
```

All the assertions included within this tag are fulfilled.

## Demo WSDL

For a complete WS security policy document, see the WSDL file included with the demo
"wssecuritymessage", located in $FGLDIR/demo/WebServices. For more information
on the standard, refer to [WS-SecurityPolicy](http://docs.oasis-open.org/ws-sx/ws-securitypolicy/200702/ws-securitypolicy-1.2-spec-os.pdf).

## Child topics

- [Security bindings](4680-security-bindings.md): Understand the mechanism of bindings that allows the secure exchange of SOAP messages over HTTP. The demo application security policy is referenced to illustrate this.
- [SOAP message security options](4681-soap-message-security-options.md): Describes the Wss10 SOAP Message Security 1.0 options that are supported.
- [SignedParts](4682-signedparts.md): The SignedParts section of the policy specifies which part of the message should be signed.
- [EncryptedParts](4683-encryptedparts.md): The EncryptedParts section of the policy specifies which part of the message should be encrypted.
