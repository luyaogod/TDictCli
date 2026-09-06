---
title: "JSONSelector"
source: "fgl-topics/c_gws_JSON_attribute_JSONSelector.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > json.Serializer attributes > JSONSelector"
type: "concept"
---

# JSONSelector

> Identifies the variant member of a JSONOneOf record to use during serialization or deserialization.

## Syntax

```
JSONSelector
```

`JSONSelector` is a mandatory attribute used with [JSONOneOf](3681-jsononeof.md "Defines a record that matches exactly one of several possible JSON schemas during serialization or deserialization.").

> **Important:**
>
> This serializer-specific attribute is supported only by [json.Serializer](3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.") and does not work with [util.JSON](3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.").

## Usage

Every [`JSONOneOf`](3681-jsononeof.md "Defines a record that matches exactly one of several possible JSON schemas during serialization or deserialization.") record
must include exactly one field marked with `JSONSelector`. That field must be
`SMALLINT` or `INTEGER`. You set its value to select the variant to
use: the value is a 1-based positional index across the whole record, counting the
`JSONSelector` field itself as position 1. When the `JSONSelector` field is the first member of the record (as in the following example), the first variant is therefore index 2; setting the value to 1 is invalid and causes error
[-15807](4483-genero-bdl-errors.md).

`JSONSelector` represents the `oneOf` keyword in the OpenAPI JSON
schema specification.

## Example using JSONSelector and JSONOneOf

In this example, the `JSONSelector` attribute is set on the
`_selector` member in the input record of the `CreateAccount`
function. To select a variant,
set `_selector` to the positional index of the variant field: 2 for
`id`, 3 for `name`.

For example, in a request from your client application, you
might have a statement like this: `LET in._selector=2`.

```
IMPORT COM

PUBLIC TYPE accountType RECORD
    id INTEGER,
    name STRING,
    date DATETIME YEAR TO SECOND,
    age INTEGER,
    gender BOOLEAN
END RECORD

PRIVATE DEFINE accounts DYNAMIC ARRAY OF accountType

PUBLIC DEFINE accountError RECORD ATTRIBUTE(WSError = 'account service error')
    id INTEGER,
    msg STRING
END RECORD

TYPE createAccountType RECORD ATTRIBUTE(JSONOneOf)
    _selector INTEGER ATTRIBUTE(JSONSelector),
    id INTEGER,
    name STRING
END RECORD

FUNCTION CreateAccount(
    in createAccountType)
    ATTRIBUTE(WSPost, WSPath = "/createaccount")
    RETURNS(createAccountType)

    DEFINE out createAccountType
    DEFINE idx INTEGER
    LET out = in

    CASE
        WHEN in._selector = 2
            LET idx = accounts.search("id", in.id)
            IF idx > 0 THEN
                # raise RESTError with accountError
            ELSE
              CALL accounts.appendElement()
              LET accounts[accounts.getLength()].id = in.id
            END IF
        WHEN in._selector = 3
            LET idx = accounts.search("name", in.name)
            # ... function code ... 
    END CASE

    RETURN out
END FUNCTION
```

## Related links

**Related concepts**  

[JSONOneOf](3681-jsononeof.md "Defines a record that matches exactly one of several possible JSON schemas during serialization or deserialization.")
