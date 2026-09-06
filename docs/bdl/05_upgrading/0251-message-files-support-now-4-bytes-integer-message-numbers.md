---
title: "Message files support now 4-bytes integer message numbers"
source: "fgl-topics/c_fgl_Migrate_to_240_message_id.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.40 upgrade guide > Message files support now 4-bytes integer message numbers"
type: "concept"
---

# Message files support now 4-bytes integer message numbers

> 2-byte .msg message number limitation was removed.

Before version 2.40, message files entries were only defined with numbers in the range -32767 to
32767 (SMALLINT). This limitation is no longer true in 2.40; message numbers can now be in the
range -2147483648 to 2147483647 (INTEGER).

## Related links

**Related concepts**  

[Message files](../11_user-interface/1593-message-files.md "Message files centralize strings and larger texts identified by a number, that can be used in programs.")
