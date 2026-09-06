---
title: "xml.StaxReader.getBytesRead"
source: "fgl-topics/c_gws_XmlStaxReader_getBytesRead.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > xml.StaxReader methods > xml.StaxReader.getBytesRead"
type: "concept"
---

# xml.StaxReader.getBytesRead

> Returns the number of bytes read by the StaxReader from the current XML document.

## Syntax

```
getBytesRead()
  RETURNS INTEGER
```

## Usage

This method returns the number of bytes read by the StaxReader since the start of the current XML
document. It returns `-1` if no bytes have been read.

For example, this method can be used to give an indication of progress when processing a large
file.
