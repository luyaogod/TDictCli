---
title: "DICTIONARY"
source: "fgl-topics/c_fgl_Dictionary_syntax.html"
breadcrumb: "Language basics > Dictionaries > DICTIONARY"
type: "concept"
---

# DICTIONARY

> A dictionary defines an associative array (hash-map) of elements.

## Syntax

```
DICTIONARY [ attributes-list ] OF type-specification
```

1. type-specification can be one of:
   - A [primitive type](0690-primitive-type-specification.md "Type definitions using a primitive data type define a primitive type.")
   - A [record definition](0717-record.md "The RECORD keyword defines a structured type or variable.")
   - An [array definition](0731-array.md "An array defines a vector variable with a list of elements.")
   - A [dictionary definition](0744-dictionary.md "A dictionary defines an associative array (hash-map) of elements.")
   - A [function type definition](0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.")
   - The name of a [user defined type](0753-type.md "Types define a synonym for a base or structured data type.")
   - The name of a [built-in class](../15_library-reference/2908-built-in-packages.md "These topics cover the built-in classes provided by the Genero Business Development Language.")
   - The name of an [imported extension
     class](../15_library-reference/3480-extension-packages.md "Several utility classes and functions are provided in additional packages.")
   - The name of an [imported Java class](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.")
2. attributes-list is a comma-separated list of name = value
   pairs or name attributes, and defines [attributes for the dictionary type](0746-attributes-on-dictionary-definitions.md "Dictionaries can be defined with attributes, to complete the type description.").

## Usage

A dictionary defines an associative array of unordered elements, accessed by a key.

Dictionary variables can invoke [methods specific to
the dictionary types](../15_library-reference/2956-dictionary-methods.md).

The elements of the dictionary can be of a simple type, or structured records.

For example, to define a dictionary of
strings:

```
DEFINE dict DICTIONARY OF STRING
```

The dictionary subscript syntax consists of a character string expression (the key), specified
between square brackets.

The result of the subscript syntax can be used as l-value (as target variable in
assignments):

```
LET dict["abcdef"] = "the value"
```

or as r-value (in expressions):

```
DISPLAY dict["abcdef"]
```

Dictionary elements are automatically created when needed. For more details, see [Dictionary in action](0747-dictionary-in-action.md).

## Related links

**Related concepts**  

[Primitive Data types](0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.")

[Variables](0686-variables.md "Explains how to define program variables.")
