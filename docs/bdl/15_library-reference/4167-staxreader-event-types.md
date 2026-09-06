---
title: "StaxReader Event Types"
source: "fgl-topics/r_gws_XmlStaxReader_event_types.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > StaxReader Event Types"
type: "reference"
---

# StaxReader Event Types

> Event types of the xml.StaxReader class.

| Type | Description | XML sample |
| --- | --- | --- |
| START\_DOCUMENT | StaxReader cursor points to the beginning of the XML document. | `<?xml version="1.0" standalone="no"?>` |
| END\_DOCUMENT | StaxReader cursor has reached the end of the XML document.No additional parsing operation will succeed. |  |
| START\_ELEMENT | StaxReader cursor points to a XML start element or empty element node. | `<p:elt attr="val">` or `<p:elt attr="val"/>` |
| END\_ELEMENT | StaxReader cursor points to a XML end element node. | `</p:elt>` |
| CHARACTERS | StaxReader cursor points to a XML text node. | `... eltA/>This is text<eltB ...` |
| CDATA | StaxReader cursor points to a XML CData node. | `<![CDATA[<Hello, world!>]]>` |
| SPACE | StaxReader cursor points to a XML text node containing only whitespaces. | `... eltA/> <eltB ...` |
| COMMENT | StaxReader cursor points to a XML comment node. | `<!-- a comment -->` |
| DTD | StaxReader cursor points to a DTD string. | `<!DOCTYPE A [ <!ELEMENT B (C+)> ]>` |
| ENTITY\_REFERENCE | StaxReader cursor points to a XML entity reference node. | `&ref;` |
| PROCESSING\_INSTRUCTION | StaxReader cursor points to a XML processing instruction node. | `<?target data?>` |
| ERROR | StaxReader cursor points to an unexpected XML node. |  |
