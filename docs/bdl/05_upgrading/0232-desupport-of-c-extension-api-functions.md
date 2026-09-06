---
title: "Desupport of C-Extension API functions"
source: "fgl-topics/c_fgl_Migrate_to_251_cext_api.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.51 upgrade guide > Desupport of C-Extension API functions"
type: "concept"
---

# Desupport of C-Extension API functions

> BIGINT and BOOLEAN stack functions and C API functions for C-Extensions are no longer supported.

> **Note:**
>
> Starting with version 3.10 (also backported in 3.00.10), the `popbigint()` and
> `pushbigint()` function are again available.

Since version 2.51:

The C-Extension stack functions to handle `BIGINT` and `BOOLEAN`
types have been removed:

| popboolean() |
| --- |
| popbigint() |
| pushboolean() |
| pushbigint() |

The C API functions such as `decadd()`, `risnull()`,
`rsetnull()`, have been removed. These functions are part of the IBM®
Informix® ESQL/C product and cannot be part
of the Genero BDL product. The Genero runtime system provides only the C functions to
[push and pop](../14_extending-the-language/2711-runtime-stack-functions.md "To pass values between a C function and a program, the C function and the runtime system use the runtime stack.") data on the Genero BDL stack.

Below is the list of C API functions that have been removed, check your C extension code for the
usage of these functions. If such functions are required, link your C-Extensions with the IBM
Informix ESQL/C libraries.

| bycmpr() |
| --- |
| byleng() |
| bycopy() |
| byfill() |
| risnull() |
| rsetnull() |
| rgetmsg() |
| rgetlmsg() |
| rtypalign() |
| rtypmsize() |
| rtypname() |
| rtypwidth() |
| rdatestr() |
| rdayofweek() |
| rdefmtdate() |
| ifx\_defmtdate() |
| rfmtdate() |
| rjulmdy() |
| rleapyear() |
| rmdyjul() |
| rstrdate() |
| ifx\_strdate() |
| rtoday() |
| ldchar() |
| rdownshift() |
| rfmtdouble() |
| rfmtint4() |
| rstod() |
| rstoi() |
| rstol() |
| rupshift() |
| stcat() |
| stchar() |
| stcmpr() |
| stcopy() |
| stleng() |
| decadd() |
| deccmp() |
| deccopy() |
| deccvasc() |
| deccvdbl() |
| deccvflt() |
| deccvint() |
| deccvlong() |
| decdiv() |
| dececvt() |
| decfcvt() |
| decmul() |
| decround() |
| decsub() |
| dectoasc() |
| dectodbl() |
| dectoflt() |
| dectoint() |
| dectolong() |
| dectrunc() |
| rfmtdec() |
| dtaddinv dtaddinv() |
| dtcurrent() |
| dtcvasc() |
| ifx\_dtcvasc() |
| dtcvfmtasc() |
| ifx\_dtcvfmtasc() |
| dtextend() |
| dtsub() |
| dtsubinv() |
| dttoasc() |
| dttofmtasc() |
| ifx\_dttofmtasc() |
| incvasc() |
| incvfmtasc() |
| intoasc() |
| intofmtasc() |
| invdivdbl() |
| invdivinv() |
| invextend() |
| invmuldbl() |

## Related links

**Related concepts**  

[C-Extensions](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.")
