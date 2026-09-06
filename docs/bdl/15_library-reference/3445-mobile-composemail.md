---
title: "mobile.composeMail"
source: "fgl-topics/c_fgl_frontcall_mobile_composemail.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile common front calls > mobile.composeMail"
type: "concept"
---

# mobile.composeMail

> Invokes the user's default mail application for a new mail to send.

## Syntax

```
ui.Interface.frontCall("mobile", "composeMail",
   [to, subject, content, cc, bcc, attachments ...],
   [result])
```

1. to - A list of recipients, separated by commas. While the list uses commas to separate the recipients in the list, the list itself is enclosed in a
   single set of quotes.
2. subject - The subject of the email.
3. content - The body of the email.
4. cc - (optional) A list of recipients for the carbon-copy email field,
   separated by commas. While the list uses commas to separate the recipients in the list, the list itself is enclosed in a single set of quotes.
5. bcc - (optional) A list of recipients for the blind carbon-copy email field,
   separated by commas. While the list uses commas to separate the recipients in the list, the list itself is enclosed in a single set of quotes.
6. attachments ... - (optional) All remaining arguments are treated as paths to
   attachment files. Each attachment filename is enclosed in its own set of quotes. The comma is used
   to separate the attachments in the list. Only supported with GMA/GMI and may be desupported in a
   future version.
7. result - Holds a status message.

## Usage

The "`mobile.composeMail`" front call is a synomym for [standard.composeMail](3395-standard-composemail.md "Invokes the user's default mail application for a new mail to send."), when executed in browsers or GDC (without email
attachment support). On GMA and GMI, this front call additionally supports the
attachment parameter. This GMA/GMI specific behavior may be desupported in a
future version.

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

The next example opens an email and populates the To, CC, and BCC fields, the Subject line, the
message body, and it specifies two attachments..

```
DEFINE result STRING
CALL ui.Interface.frontCall("standard", "composeMail",
 ["john.doe@4js.com,jane.doe@4js.com", "Hello world",
  "This is the hello world text", "john.doe@4js.com,jane.doe@4js.com", "hidden@4js.com",
  "/sdcard/Pictures/photo1.jpg", "/sdcard/Pictures/photo2.jpg" ], [result])
```

Support for email attachments is only available with GMA/GMI applications and may be desupported
in a future version.
