---
title: "json.Serializer.getLastErrorDescription"
source: "fgl-topics/c_gws_serializer_getLastErrorDescription.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > json.Serializer methods > json.Serializer.getLastErrorDescription"
type: "concept"
---

# json.Serializer.getLastErrorDescription

> Retrieves the most recent error generated during serialization.

## Syntax

```
getLastErrorDescription() RETURN STRING
```

The error message is returned or `NULL` if no errors were encountered.

## Usage

Use this method to retrieve the last error message generated. The last error message is stored
and will remain accessible even after multiple successful calls to the serializer. The error message
is prefixed with the class name to identify the source of the error:

- JSONReader
- JSONWriter
- Serializer
- Deserializer

The `getLastErrorDescription()` method can be used in preference to
`sqlca.sqlerrm` because it is not restricted to a message length of 72 characters.

Where messages in `sqlca.sqlerrm` can include ellipsis ("`...`") to
signify a message is incomplete, the complete message can be viewed via the
`getLastErrorDescription()` method. In addition,
`getLastErrorDescription()` returns helpful information with the name of the class,
the name of the fields, when known, along with a detailed description of the error.

## Example:

In this example, a deserializer error is thrown
because the required fields are missing. In the `CATCH` block, an error message is
displayed by both `sqlca.sqlerrm` and
`getLastErrorDescription()`

```
IMPORT json

MAIN
  CALL deserialization()
END MAIN

FUNCTION deserialization()
  DEFINE reader json.JSONReader
  DEFINE rec RECORD
    firstname   STRING ATTRIBUTE(JSONRequired),
    lastname    STRING ATTRIBUTE(JSONRequired),
    address     STRING ATTRIBUTE(JSONRequired),
    phoneNumber STRING ATTRIBUTE(JSONRequired)
  END RECORD
  DEFINE t TEXT

  LOCATE t IN MEMORY

  TRY
    LET t = "{}"
    LET reader = json.JSONReader.Create()
    CALL reader.readFromText(t)
    CALL reader.next()
    CALL json.Serializer.JSONToVariable(reader, rec)
  CATCH
    DISPLAY sqlca.sqlerrm
    DISPLAY json.Serializer.getLastErrorDescription()
  END TRY

END FUNCTION
```

The error message output by `sqlca.sqlerrm` is
truncated:

```
Missing required properties 'firstname, lastname, address, p... at (1,2)
```

The
message from `getLastErrorDescription()` is complete:

```
Deserializer: Missing required properties 'firstname, lastname, address, phoneNumber' at (1,2)
```
