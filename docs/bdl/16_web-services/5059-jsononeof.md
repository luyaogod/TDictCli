---
title: "JSONOneOf"
source: "fgl-topics/c_gws_JSON_attribute_JSONOneOf_2.html"
breadcrumb: "Web services > Reference > JSON serialization rules and customization > json.Serializer attributes > JSONOneOf"
type: "concept"
---

# JSONOneOf

> Defines a record that matches exactly one of several possible JSON schemas during serialization or deserialization.

## Syntax

```
JSONOneOf
```

`JSONOneOf` is an optional attribute.

> **Important:**
>
> This serializer-specific attribute is supported only by [json.Serializer](../15_library-reference/3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.") and does not work with [util.JSON](../15_library-reference/3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.").

## Usage

Set `JSONOneOf` on a record, elements of a record, or a user-defined type to define a polymorphic JSON
structure where exactly one variant member is used during serialization or deserialization.

Every `JSONOneOf` record must include one field marked with [`JSONSelector`](../15_library-reference/3684-jsonselector.md "Identifies the variant member of a JSONOneOf record to use during serialization or deserialization."). At runtime, you
set the value of that field to select which variant to use. The value is a 1-based positional index
across the whole record, counting the `JSONSelector` field itself as position 1. When the `JSONSelector` field is the first member of the record (as in these examples), the first variant is therefore index 2; setting the selector to 1 is invalid and causes error
[-15807](../15_library-reference/4483-genero-bdl-errors.md).

For example:

```
TYPE createAccountType RECORD ATTRIBUTE(JSONOneOf)
    _selector INTEGER ATTRIBUTE(JSONSelector),  # position 1
    id INTEGER,                                  # position 2: LET _selector = 2
    name STRING                                  # position 3: LET _selector = 3
END RECORD
```

When defining schemas, keep in mind how Genero BDL types map to JSON types; see [OpenAPI types mapping: BDL](4847-openapi-types-mapping-bdl.md "Conversion mapping for Genero BDL data types in OpenAPI documentation.").

During deserialization, error [-15807](../15_library-reference/4483-genero-bdl-errors.md) is
returned if an incoming value is valid for two or more variants that map to the same JSON type. For
example, this record causes -15807 because both `street_address` and
`po_box` map to JSON `STRING`:

```
TYPE deliveryAddressType RECORD ATTRIBUTE(JSONOneOf)
    _selector SMALLINT ATTRIBUTE(JSONSelector),
    street_address STRING,   # maps to JSON STRING
    po_box STRING            # maps to JSON STRING: error -15807
END RECORD
```

The following JSON types can conflict:

- `STRING`
- `STRING format: DATE/DATETIME`
- `BOOLEAN`
- `INTEGER`
- `NUMBER`

Note that `TINYINT`, `SMALLINT`, and `BIGINT` all
map to `INTEGER`, and `DECIMAL` maps to `NUMBER`; schemas that appear distinct in BDL may conflict once converted.

When a value could match more than one type, special priority rules apply; see [Special rules](../15_library-reference/3681-jsononeof.md).

`JSONOneOf` represents the `oneOf` keyword in the OpenAPI JSON
schema specification.

## Special rules

When a `oneOf` value is valid for more than one type, the GWS engine resolves
the conflict with a priority rule:

- `DATE` or `DATETIME` takes priority over
  `STRING`.
- `INTEGER` takes priority over `NUMBER`.

Priority applies only when deserialization succeeds for the higher-priority type. The
Deserialization, Priority, and
Result columns in the tables below show how the rule is applied for each
type combination.

| Type | Deserialization | Priority | Result |
| --- | --- | --- | --- |
| `DATETIME` | YES | YES | `DATETIME` |
| `STRING` | YES | NO | `DATETIME` |

| Type | Deserialization | Priority | Result |
| --- | --- | --- | --- |
| `DATE` | NO | YES | `STRING` |
| `STRING` | YES | NO | `STRING` |

| Type | Deserialization | Priority | Result |
| --- | --- | --- | --- |
| `DATETIME` | NO | YES | `STRING` |
| `STRING` | YES | NO | `STRING` |

| Type | Deserialization | Priority | Result |
| --- | --- | --- | --- |
| `INTEGER` | YES | YES | `INTEGER` |
| `NUMBER` | YES | NO | `INTEGER` |

| Type | Deserialization | Priority | Result |
| --- | --- | --- | --- |
| `INTEGER` | YES | YES | `INTEGER` |
| `NUMBER` | YES | NO | `INTEGER` |

| Type | Deserialization | Priority | Result |
| --- | --- | --- | --- |
| `INTEGER` | NO | YES | `NUMBER` |
| `NUMBER` | YES | NO | `NUMBER` |

## Example 1: using JSONOneOf

In this example, `JSONOneOf` is set on the input record of the
`CreateAccount` function. To select a variant,
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

The GWS engine exposes the `oneOf` property in the OpenAPI schema for records
defined with `JSONOneOf`. The following schema shows both variants:

The output shown in this example is from the Firefox™ browser, which formats JSON for readability. The appearance may vary depending on your browser.

![Image from the OpenAPI document showing the oneOf JSON schema property](../_images/rest_openapi_api_jsononeof_json_schema.png)

*Sample JSON schema with oneOf property*

## Example 2: using JSONOneOf with JSONRequired

By default, JSON schema allows additional properties, so all variant schemas validate against
any submitted document and `oneOf` cannot resolve. Adding
`JSONRequired` to one distinctive field per variant causes the GWS engine to reject
any schema that is missing that field, leaving exactly one valid schema. For an alternative
approach using `JSONAdditionalProperties`, see [Example 3: using JSONOneOf with JSONAdditionalProperties](../15_library-reference/3678-jsonadditionalproperties.md).

```
# In this example, a property with attribute JSONRequired is mandatory 
# Otherwise, all schemas will be valid and oneOf can only accept one valid schema
 
TYPE complexType RECORD ATTRIBUTE(JSONOneOf)
    _selector INTEGER ATTRIBUTE(JSONSelector),
    accountById RECORD
        id INTEGER ATTRIBUTE(JSONRequired),
        birthday DATE,
        gender BOOLEAN
    END RECORD,
    accountByname RECORD
        name STRING ATTRIBUTE(JSONRequired),
        birthday DATE,
        gender BOOLEAN
    END RECORD,
    accountBydate RECORD
        date DATETIME YEAR TO SECOND ATTRIBUTE(JSONRequired),
        birthday DATE,
        gender BOOLEAN
    END RECORD
END RECORD

FUNCTION echoComplexType(
    in complexType)
    ATTRIBUTE(WSPost, WSPath = "/echoComplexType")
    RETURNS(complexType)

    DEFINE out complexType

    LET out = in
    DISPLAY out._selector

    CASE
        WHEN out._selector = 2
            DISPLAY "schema 1, id: ", out.accountById.id

        WHEN out._selector = 3
            DISPLAY "schema 2, name: ", out.accountByname.name

        WHEN out._selector = 4
            DISPLAY "schema 3, date: ", out.accountBydate.date
    END CASE

    RETURN out
END FUNCTION
```

In the OpenAPI documentation, the GWS creates the JSON schema with `oneOf` and
`required` properties.

The output shown in this example is from the Firefox browser, which formats JSON for readability. The appearance may vary depending on your browser.

![Image from the OpenAPI document showing the oneOf and required JSON schema properties](../_images/rest_openapi_api_jsononeof_jsonrequired_json_schema.png)

*Sample JSON schema with oneOf and required properties*

## Related links

**Related concepts**  

[JSONSelector](../15_library-reference/3684-jsonselector.md "Identifies the variant member of a JSONOneOf record to use during serialization or deserialization.")
