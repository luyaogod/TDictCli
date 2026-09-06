---
title: "Type checking with fglcomp compiler"
source: "fgl-topics/c_fgl_Migrate_to_310_fglcomp_type_check.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > Type checking with fglcomp compiler"
type: "concept"
---

# Type checking with fglcomp compiler

> The fglcomp compiler is more strict regarding type checking.

Starting with Genero 3.10, the fglcomp compiler is now more strict when
assigning complex types.

Any assignment potentially throwing runtime error [-1260](../15_library-reference/4483-genero-bdl-errors.md) is subject
to the new compiler error [-6631](../15_library-reference/4483-genero-bdl-errors.md).

> **Note:**
>
> Assignments using primitive types such as `CHAR` to `INTEGER`
> are not checked, except for `DATETIME`, `INTERVAL`,
> `TEXT` and `BYTE` assignments, which are candidates for runtime error
> -1260.

Type checking errors are raised on:

- incompatible assignments in `LET variable = value`.
- passing incompatible values to functions (user functions, built-in functions, C extensions).
- returned types of functions when the `FUNCTION` is defined with the
  `RETURNS` clause.

Better type checks include:

- better predictions (for autocompletion in vim and Studio)
- being able to call methods on return values of methods.

For example, when assigning invalid object references to variables defined with a different
class:

```
$ cat tc1.4gl
MAIN
    DEFINE sb base.StringBuffer
    LET sb =  "foo" -- illegal: assigns a  String to a  StringBuffer
END MAIN
```

With versions prior to 3.10, you get no compiler error, but an error at
runtime:

```
$ fglcomp -V
fglcomp 3.10.14
...
$ fglcomp tc1

$ fglrun tc1
Program stopped at 'tc1.4gl', line number 3.
FORMS statement error number -1260.
It is not possible to convert between the specified types.
```

With version 3.10, you get now a compiler
error:

```
$ fglcomp -V
fglcomp 3.10.03
...
$ fglcomp -M tc1
tc1.4gl:3:5:3:18:error:(-6631) incompatible types, found: CHAR, required: base. StringBuffer.
```

In the following example, the type returned by `base.Channel.create()` does not
match the variable
definition:

```
$ cat tc2.4gl
MAIN
    DEFINE sb base.StringBuffer
    LET sb = base.Channel.create() -- what's the return type of base.Channel.create()?
END MAIN

$ fglcomp -M tc2
tc1.4gl:3:5:3:18:error:(-6631) incompatible types, found: base.Channel, required: base. StringBuffer.
```

This example tries to pass a `STRING` to a method whereas an
`om.DomNode` is
expected:

```
$ cat tc3.4gl
MAIN
    DEFINE doc om.DomDocument
    DEFINE n1, n2 om.DomNode
    LET doc = om.DomDocument.createFromString("<Foo><Bar/></Foo>")
    LET n1 = doc.getDocumentElement()
    LET n2 = doc.getDocumentElement().getFirstChild()
    LET n2 = n1.getFirstChild()
    CALL n1.removeChild(n2) -- legal: n2 is an om.DomNode
    CALL n1.removeChild("Bar") -- illegal: n2 is a STRING
END MAIN

$ fglcomp -M tc3.4gl
tc3.4gl:9:26:9:30:error:(-6631) incompatible types, found: CHAR, required: om.DomNode.
```

This example uses the return value of FGL-functions:

```
$ cat tc4.4gl
MAIN
    DEFINE s STRING
    LET s = function1() -- unchecked: return type unknown
    LET s = function2() -- illegal: returns a base.Channel
    LET s = function3() -- legal: returns a STRING
END MAIN

FUNCTION function1()
    RETURN base.Channel.create()
END FUNCTION

FUNCTION function2() RETURNS base.Channel
    RETURN base.Channel.create()
END FUNCTION

FUNCTION function3() RETURNS STRING
    RETURN  "foo"
END FUNCTION

$ fglcomp tc4.4gl
The compilation was not successful.  Errors found: 1.
 The file 'tc4.err' has been written.

$ cat tc4.err
MAIN
    DEFINE s STRING
    LET s = function1() -- unchecked: return type unknown
    LET s = function2() -- illegal: returns a base.Channel
| incompatible types, found: base.Channel, required: STRING.
| See error number -6631.
    LET s = function3() -- legal: returns a STRING
END MAIN
...
```

## Related links

**Related concepts**  

[Primitive Data types](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.")

[fglcomp](../13_programming-tools/2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.")

[Calling functions](../08_language-basics/0766-calling-functions.md "Functions can be invoked, to execute the code they define.")

[Function references](../08_language-basics/0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.")
