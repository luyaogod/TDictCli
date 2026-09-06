---
title: "OAuthAPI library"
source: "fgl-topics/c_gws_oauthapi.html"
breadcrumb: "Web services > Reference > OAuthAPI library"
type: "concept"
---

# OAuthAPI library

> The OAuthAPI library provides a set of functions for working with access to web services using the OAuth protocol.

## The OAuthAPI file

The OAuthAPI library is delivered in the FGLGWS package. It is located in
$FGLDIR/lib.

To use an OAuthAPI function, declare the module with the [`IMPORT FGL`](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.") instruction:

```
IMPORT FGL OAuthAPI
#...
  CALL OAuthAPI.InitService(5, access_token)
```

It is recommended that the
OAuthAPI.42m library file is linked into every Genero Web Services client
program that requires access to a web service secured by OpenID Connect.

## Child topics

- [OAuthAPI overview](4931-overview.md): The OAuthAPI library supports the OAuth protocol that authenticates user access and issues access tokens.
