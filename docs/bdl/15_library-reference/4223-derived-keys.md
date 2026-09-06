---
title: "Derived keys"
source: "fgl-topics/c_gws_XmlCryptoKey_derived_keys.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > Derived keys"
type: "concept"
description: "Key derivation is used on symmetric or HMAC keys to avoid the direct usage of a shared secret password in secured operations. If two parties share a secret password that is successfully hacked by a ..."
---

# Derived keys

Key derivation is used on symmetric or HMAC keys to avoid the direct usage of a shared secret
password in secured operations. If two parties share a secret password that is successfully hacked
by a third party, any future operations become insecure, and the initial two parties do not even
realize that their exchanges are unsafe. However, if a different password based on that shared
secret password is used for each new secured operation, even if one operation is compromised, it
will only be insecure for that operation, but not other operations.

The derivation consists of applying an algorithm with some additional inputs (such as a random
seed value) to a password in order to obtain another password that is then used in one secured
operation. Of course, the algorithm and its additional inputs must also be shared to enable the
computation of the same derived key for the decryption of the message by the person it is intended
for.

Note that passwords are often only composed of alphanumeric characters, which makes the job of a
hacker a little bit easier, whereas a derived key is composed of any binary data produced by the
algorithm used for the derivation.

| Method | Description |
| --- | --- |
| http://schemas.xmlsoap.org/ws/2005/02/sc/dk/p\_sha1 | Only algorithm supported. See [specification](http://docs.oasis-open.org/ws-sx/ws-secureconversation/200512/ws-secureconversation-1.3-os.html#_Toc162064056) for details. |
