---
title: "ARRAY"
source: "fgl-topics/c_fgl_Arrays_002.html"
breadcrumb: "Language basics > Arrays > ARRAY"
type: "concept"
---

# ARRAY

> An array defines a vector variable with a list of elements.

## Syntax 1: Static array definition

```
ARRAY dim-spec [ attributes-list ]
    OF type-specification
```

where dim-spec is:
> **Note:**
>
> In the next(s) syntax diagram(s), the `[ ] { } |`
> symbols are part of the syntax.

- For a one-dimensional static
  array:

  ```
  [ dim ]
  ```
- For two-dimensional static
  array:

  ```
  [ dim1, dim2 ]
  ```
- For a three-dimensional static
  array:

  ```
  [ dim1, dim2, dim3 ]
  ```

1. dim, dim1, dim2 and
   dim3 can be integer literals or constants. The upper limit is 65535.
2. type-specification can be one of:
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
3. attributes-list is a comma-separated list of
   name = value pairs or name attributes, and defines [attributes for the array type](0736-attributes-on-array-definitions.md "Arrays can be defined with attributes, to complete the type description.").

## Syntax 2: Dynamic array definition

```
DYNAMIC ARRAY [ attributes-list ] [ WITH DIMENSION rank ]
    OF type-specification
```

1. rank an be an integer literal greater than zero. Default is 1.
2. type-specification can be one of:
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
3. attributes-list is a comma-separated list of
   name = value pairs or name attributes, and defines [attributes for the array type](0736-attributes-on-array-definitions.md "Arrays can be defined with attributes, to complete the type description.").

## Syntax 3: Java array definition

(in the next diagram, the square brackets are part of the syntax)

```
ARRAY [ ] OF java-type
```

1. java-type must be a [Java class](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.")
   or a simple data type that has a corresponding primitive type in Java, such as
   `INTEGER` (int), `FLOAT` (double).

## Usage

An array defines an ordered set of elements.

Array variables can invoke [methods specific to the
array types](../15_library-reference/2944-dynamic-array-methods.md).

The type of the array elements can be of a simple type or structured records.

Consider using dynamic arrays instead of static arrays.

Java-style arrays will only be useful to interface with Java calls.

Static and dynamic arrays can be defined with the `ATTRIBUTES()` clause, to
specify meta-data information for the variable. This feature is especially used when defining
variables for XML-based Web Services. For more details about XML attributes, see
[XML serialization rules and customization](../16_web-services/4958-xml-serialization-rules-and-customization.md).

## Related links

**Related concepts**  

[Primitive Data types](0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.")

[Using Java arrays](../14_extending-the-language/2687-using-java-arrays.md "Using Java arrays")

[Variables](0686-variables.md "Explains how to define program variables.")
