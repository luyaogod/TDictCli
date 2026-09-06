---
title: "Need for Informix CSDK to compile C extensions"
source: "fgl-topics/c_fgl_Migrate_to_232_informix_csdk_cext.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.32 upgrade guide > Need for Informix® CSDK to compile C extensions"
type: "concept"
---

# Need for Informix CSDK to compile C extensions

> Compiling C Extensions requires now the Informix CSDK.

> **Note:**
>
> This upgrade note is related to C Extensions or ESQL/C Extensions, and can be
> ignored if your application does not use such extensions.

To compile C or ESQL/C extensions manipulating data types like DECIMAL, you need IBM®
Informix® data type structure definitions such
as `dec_t`, `dtime_t`, `intrvl_t`, as well
as macros like DECLEN() or TU\_ENCODE(). Before version 2.32, these C structure and
macros where provided in the files of the $FGLDIR/include/f2c
directory.

The Genero BDL version 2.32 no longer provides the IBM
Informix ESQL/C structure definitions in
$FGLDIR/include/f2c files, because we have identified that some
of the definitions are platform specific. However, to compile your C extensions, you
need these definitions if your extensions use complex data types such as DECIMAL,
DATETIME/INTERVAL, BYTE/TEXT. The definitions are not required if you use standard C
types such as `int` or `char[]`.

Starting with version 2.32, you need to install an IBM
Informix CSDK on your development machine in
order to get the structure and macro definitions to compile your C extensions. The IBM
Informix CSDK is only required on the
development platform. It is not required to install the CSDK on the
production machines, except of course if you want to connect to an IBM
Informix database server.

## Related links

**Related concepts**  

[C-Extensions](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.")
