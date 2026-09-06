---
title: "fgl_setenv()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_SETENV.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_setenv()"
type: "concept"
---

# fgl_setenv()

> Sets the value of an environment variable.

## Syntax

```
FUNCTION fgl_setenv(
   name STRING,
   value STRING )
```

1. name is the name of the environment variable.
2. value is the value to be set.

## Usage

The `fgl_setenv()` function sets or modifies the value of an environment variable.

> **Important:**
>
> Use the `fgl_setenv()` function with care: [Genero environment variables](../07_configuration/0506-genero-environment-variables.md) such as FGLDIR, DBDATE,
> DBCENTURY or FGL\_LENGTH\_SEMANTICS should be defined before starting the process, and must not be
> changed at runtime, as it can lead to unexpected behavior of the current process. The same rule
> applies to [database client environment variables](../10_sql-support/1062-database-client-environment.md "To connect to a database server, Genero BDL programs use vendor's database client software.") like
> INFORMIXDIR, ORACLE\_HOME, and environment variables of third party software components like
> JAVA\_HOME, when using the [Java interface](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.").

There is a little difference between Windows® and UNIX™ platforms when passing a [NULL](../08_language-basics/0572-null.md "The NULL constant defines a non-value.") as
the value parameter. On Windows
platforms, the environment variable is removed, while on UNIX, the environment variable gets an empty value
(it is not removed from the environment).

## Related links

**Related concepts**  

[fgl\_getenv()](2759-fgl-getenv.md "Returns the value of the environment variable.")
