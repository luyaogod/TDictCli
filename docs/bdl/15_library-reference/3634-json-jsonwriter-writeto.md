---
title: "json.JSONWriter.writeTo"
source: "fgl-topics/c_gws_jsonJSONWriter_writeTo.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONWriter class > json.JSONWriter methods > json.JSONWriter.writeTo"
type: "concept"
---

# json.JSONWriter.writeTo

> Sets the output stream of the JSONWriter object to a file or an URL, and starts streaming data.

## Syntax

```
writeTo(
   url STRING )
```

1. url defines a valid URL or the name of the file that will contain the
   resulting JSON document.

## Usage

This method sets the output stream of the `JSONWriter` object to a file or a URL
specified in url, and starts the streaming.

Only the following kinds of URLs are supported:

- http://
- https://
- tcp://
- tcps://
- file:///
- alias://

See [fglprofile
Configuration](../16_web-services/4915-web-services-fglprofile-configuration.md "The configuration for the Genero Web Services is defined from entries in the FGLPROFILE file.") for more details about URL mapping with aliases, and for proxy and security
configuration.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Examples

```
writeTo("printerList.json")
```

```
writeTo("http://myserver:1100/documents/printerList.json")
```

```
writeTo("https://myserver:1100/documents/printerList.json")
```

```
writeTo("alias://printerlist")
```

In the example `printerlist`
alias is defined in fglprofile as `ws.printerlist.url =
"http://myserver:1100/documents/printerList.json"`.

## Related links

**Related concepts**  

[Example 1: Writing to a JSON file](3638-example-1-writing-to-a-json-file.md "This example shows two ways to write a JSON file.")
