---
title: "standard.composeMail"
source: "fgl-topics/c_fgl_frontcall_standard_composemail.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.composeMail"
type: "concept"
---

# standard.composeMail

> Invokes the user's default mail application for a new mail to send.

## Syntax

```
ui.Interface.frontCall("standard", "composeMail",
   [to, subject, content, cc, bcc],
   [result])
```

1. to - A list of recipients, separated by commas. While the list uses commas to separate the recipients in the list, the list itself is enclosed in a single set of quotes.
2. subject - The subject of the email.
3. content - The body of the email.
4. cc - (optional) A list of recipients for the carbon-copy email field,
   separated by commas. While the list uses commas to separate the recipients in the list, the list itself is enclosed in a single set of quotes.
5. bcc - (optional) A list of recipients for the blind carbon-copy email field,
   separated by commas. While the list uses commas to separate the recipients in the list, the list itself is enclosed in a single set of quotes.
6. result - Holds a status message.

## Usage

The "`composeMail`" front call invokes the default email application and sets
up a new mail to send.

The returned result string takes one of the following values:

- "`ok`": The email app could be started.
- "`failed`": The email app could not be started.

This example opens an email and populates the To field, the Subject line, and the message body.
No CC or BCC recipients and no attachments are specified.

```
DEFINE result STRING
CALL ui.Interface.frontCall("standard","composeMail",
 ["huhu@haha.com","test mail","sent from my device"],[result])
```

## Related links

**Related concepts**  

[mobile.composeMail](3445-mobile-composemail.md "Invokes the user's default mail application for a new mail to send.")
