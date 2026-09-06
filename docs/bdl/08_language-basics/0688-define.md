---
title: "DEFINE"
source: "fgl-topics/c_fgl_variables_DEFINE.html"
breadcrumb: "Language basics > Variables > DEFINE"
type: "concept"
---

# DEFINE

> The DEFINE instruction declares a program variable with a given type.

## Syntax

```
[PUBLIC|PRIVATE] DEFINE
 { identifier [,...] type-specification
 | identifier type-specification [ = initializer ]
 } [,...]
```

1. identifier is the name of the variable, that must
   follow the convention for [identifiers](0551-identifiers.md "A Genero BDL identifier is a sequence of characters used to identify a program entity.").
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
3. initializer is a [variable initializer](0691-variable-initializers.md "Variables can be initialized in their definition."), that corresponds to the
   type-specification. For example, if the type-specification is
   a [record definition](0717-record.md "The RECORD keyword defines a structured type or variable."), the initializer
   must be a [record initializer](0719-record-initializers.md "Records can be initialized in their definition.").

## Usage

A variable is a named location in memory that can store a single value,
or an ordered set of values. Variables can be global to the program, module-specific, or local
to a function.

Any program variable needs to be declared with a `DEFINE` statement before
it is used.

By default, module-specific variables are private; They cannot be used by an other module of
the program. In order to improve code re-usability by data encapsulation, we recommend you keep
module variables private, except if you want to share large data (like arrays) between modules.
To make a module variable public, add the `PUBLIC` keyword before
`DEFINE`. When a module variable is declared as public, it can be referenced in
another module by using the [`IMPORT FGL`](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.")
instruction.

When defining variables with the `LIKE` clause, the data types are taken from the
database schema file at compile time. Make sure that the schema file of the database schema during
development corresponds to the database schema of the production database; otherwise the variables
defined in the compiled version of your modules will not match the table structures of the
production database. For more details, see [Database column types](0695-database-column-types.md "Simple variables and record structures can be defined from database columns types.").

To write well-structured programs, avoid global variables. If you need persistent data storage
during a program's execution, use variables local to the module and give access to them with
functions, or make the module variables `PUBLIC` to other modules. For more details,
see [Declaration context](0693-declaration-context.md "A variable can be declared in different contexts, which defines its visibility.").

Variables can be defined with the `ATTRIBUTES()` clause, to specify
meta-data information for the variable. For more details, see [Attributes on variable definitions](0692-attributes-on-variable-definitions.md "Variables can be defined with meta-data information.").

Variables can be defined to hold a [function
reference](0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression."). Best practice is to declare a user-defined type with the function type, then
define the variables using the [user type](0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.").

By adding an equal sign followed by a value, you can save an additional [`LET`](0701-let.md "The LET statement assigns values to variables.") instruction to initialize variables.
Initializers are especially useful to set values of structured variables (records, arrays) . For
more details, see [Variable default values](0697-variable-default-values.md "Variables get a default value when defined.").

## Related links

**Related concepts**  

[Records](0715-records.md "Records allow structured program variables definitions.")

[Arrays](0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.")
