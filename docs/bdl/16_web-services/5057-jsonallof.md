---
title: "JSONAllOf"
source: "fgl-topics/c_gws_JSON_attribute_JSONAllOf_2.html"
breadcrumb: "Web services > Reference > JSON serialization rules and customization > json.Serializer attributes > JSONAllOf"
type: "concept"
---

# JSONAllOf

> Combine multiple record schemas into a single type using the JSONAllOf attribute in Genero BDL.

## Syntax

```
JSONAllOf
```

`JSONAllOf` is an optional attribute that allows you to combine multiple record types into a single schema, following the JSON Schema `allOf` keyword.

> **Important:**
>
> This serializer-specific attribute is supported only by [json.Serializer](../15_library-reference/3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.") and does not work with [util.JSON](../15_library-reference/3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.").

## Usage

The `JSONAllOf` attribute lets you define a type that combines the properties of
multiple record types. This is useful when you want to build more complex types from smaller,
reusable components. The resulting type includes all unique properties from the referenced record
types.

If you need to preserve the original structure of `allOf` schemas instead of
merging them, you can use an option in fglrestful. By running [`fglrestful
--no-merge-allof`](../13_programming-tools/2524-fglrestful.md), you disable property merging for schemas that use the
`JSONAllOf` attribute, ensuring the schema’s allOf composition remains intact.

## BDL Example

This sample illustrates how to apply the `JSONAllOf` and
`JSONRequired` attributes to define composite and mandatory properties in JSON
schemas.

```
IMPORT com

DEFINE g_nextId INTEGER = 1
DEFINE Employees DYNAMIC ARRAY OF Employee

TYPE Address RECORD
    street STRING,
    city STRING ATTRIBUTE(JSONRequired)
END RECORD

TYPE Person RECORD
    id INTEGER,
    name STRING ATTRIBUTE(JSONRequired),
    age INTEGER
END RECORD

TYPE Employee RECORD ATTRIBUTE(JSONAllOf)
    employeeInfo Person,
    employeeLocation Address
END RECORD

FUNCTION getEmployee(
    id INTEGER ATTRIBUTE(WSParam))
    ATTRIBUTE(WSGet, WSPath = "/getEmployee/{id}")
    RETURNS(Employee)
    DEFINE emp Employee
    DEFINE ind INTEGER
    DEFINE foundId BOOLEAN

    IF Employees.getLength() = 0 THEN
        LET Employees[1].employeeInfo.id = 1
        LET Employees[1].employeeInfo.name = "Alice"
        LET Employees[1].employeeInfo.age = 30
        LET Employees[1].employeeLocation.street = "Rue Montesquieu"
        LET Employees[1].employeeLocation.city = "Paris"
    END IF

    FOR ind = 1 TO Employees.getLength()
        IF Employees[ind].employeeInfo.id = id THEN
            LET emp = Employees[ind]
            LET foundId = TRUE
            EXIT FOR
        END IF
    END FOR

    IF NOT foundId THEN
        CALL com.WebServiceEngine.SetRestError(404, "Employee not found")
    END IF

    RETURN emp
END FUNCTION

FUNCTION addEmployee(
    empIn Employee)
    ATTRIBUTE(WSPost, WSPath = "/addEmployee", WSRetCode = '2XX')
    RETURNS(Employee)
    DEFINE empOut Employee

    # Generate a dummy ID
    LET g_nextId = g_nextId + 1
    LET empIn.employeeInfo.id = g_nextId
    LET Employees[Employees.getLength() + 1] = empIn

    DISPLAY "getlength: ", Employees.getLength()

    # Log server
    DISPLAY SFMT("Adding employee '%1' with id: %2",
        Employees[Employees.getLength()].employeeInfo.name,
        Employees[Employees.getLength()].employeeInfo.id)

    CALL com.WebServiceEngine.SetRestStatus(201)

    LET empOut = empIn
    RETURN empOut
END FUNCTION
```

Where:

1. The `JSONAllOf` attribute is set on the Employee
   `TYPE` definition. Note: this attribute can only be applied to `TYPE`
   definitions.
2. Two separate records, Person and Address, are defined and
   combined into Employee at runtime. Each record has unique field names, ensuring
   compatibility with JSON Schema `allOf` semantics.
3. `JSONRequired` is used with `name` in `Person` and
   `city` in `Address` to enforce these mandatory fields.

In the OpenAPI documentation, the GWS creates the JSON schema with `allOf`
property.

The output shown in this example is from the Firefox™ browser, which formats JSON for readability. The appearance may vary depending on your browser.

![Image from the OpenAPI document showing the "allof'" and "required" JSON schema properties](../_images/rest_openapi_api_allof.png)

*Sample JSON schema with allof and required properties*

## Compilation and validation rules

1. **Restriction to TYPE definitions**: `JSONAllOf` can only be used in `TYPE` definitions.
2. **Mutual exclusion**: `JSONAllOf` cannot be used with
   `JSONOneOf` or `JSONAdditionalProperties` on the same type. Using
   `JSONAllOf` with `JSONOneOf` or
   `JSONAdditionalProperties` will result in [error-9150](../15_library-reference/4483-genero-bdl-errors.md)
3. **Restriction to RECORD types**: Only `RECORD` types are allowed as subschemas.
   Using `JSONAllOf` with unsupported types or functions will result in [error-9153](../15_library-reference/4483-genero-bdl-errors.md)
4. **Unique property names**: All property names in the combined records must be unique. Using
   `JSONAllOf` with duplicate properties will result in [error-9154](../15_library-reference/4483-genero-bdl-errors.md)
5. **No nested JSONAllOf**: Nesting `JSONAllOf` attributes within another
   `JSONAllOf` structure is not permitted, as it can create ambiguity and will result in
   [error-9155](../15_library-reference/4483-genero-bdl-errors.md).

## Related links

**Related concepts**  

[Using RESTful attributes in functions](4798-using-restful-attributes-in-functions.md "RESTful attributes define functions for your RESTful web service.")

[JSONOneOf](../15_library-reference/3681-jsononeof.md "Defines a record that matches exactly one of several possible JSON schemas during serialization or deserialization.")

[JSONAdditionalProperties](../15_library-reference/3678-jsonadditionalproperties.md "Allows a record to accept JSON properties not explicitly listed in its schema definition.")

**Related reference**  

[json.Serializer options](../15_library-reference/3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.")
