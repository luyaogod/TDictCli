---
title: "JSONReader Event Types"
source: "fgl-topics/r_gws_jsonJSONReader_event_types.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONReader class > JSONReader Event Types"
type: "reference"
---

# JSONReader Event Types

> Event types of the json.JSONReader class.

| Type | Description | JSON sample |
| --- | --- | --- |
| START\_DOCUMENT | JSONReader cursor points to the beginning of the JSON document. | An open curly brace `{` or an open square bracket `[` |
| END\_DOCUMENT | JSONReader cursor has reached the end of the JSON document.No additional parsing operation will succeed. | A close curly brace `}` or a close square bracket `]` |
| START\_OBJECT | JSONReader cursor points to a JSON start object node. | `{` An open curly braceFor example, `{"firstName":"John", "lastName":"Doe"}` |
| END\_OBJECT | JSONReader cursor points to a JSON end object node. | `}` A close curly braceFor example, `{"firstName":"John", "lastName":"Doe"}` |
| START\_ARRAY | JSONReader cursor points to a JSON start array node. | `[` an open square bracketFor example, `["Ford", "BMW", "Fiat"]` |
| END\_ARRAY | JSONReader cursor points to a JSON end array node. | `]` a close square bracketFor example, `["Ford", "BMW", "Fiat"]` |
| PROPERTY | JSONReader cursor points to the field name of a JSON data name/value pair. | `{"firstName":"John"}` |
| STRING | JSONReader cursor points to a string value of a JSON data name/value pair. | `{"lastname":"Doe"}` |
| NUMBER | JSONReader cursor points to a number value of a JSON data name/value pair. | `{"age":30}` |
| BOOLEAN | JSONReader cursor points to a boolean value of a JSON data name/value pair. | `{"sale":true}` |
| NULL | JSONReader cursor points to a null value of a JSON data name/value pair. | `{"middlename":null}` |
