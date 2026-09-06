---
title: "The BIGSERIAL / SERIAL8 data types"
source: "fgl-topics/c_fgl_odiagifx_026.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Partially supported IBM® Informix® SQL features > The BIGSERIAL / SERIAL8 data types"
type: "concept"
description: "IBM® Informix® supports the BIGSERIAL and SERIAL8 data types for auto-generated 64 bit integer sequences. The BIGINT data type can be used to store data from BIGSERIAL SERIAL8 values. Note that ..."
---

# The BIGSERIAL / SERIAL8 data types

IBM® Informix®
supports the BIGSERIAL and SERIAL8 data types for auto-generated 64 bit integer sequences.

The [BIGINT](../08_language-basics/0554-bigint.md "The BIGINT data type is used for storing very large whole numbers.") data type can be used to store
data from BIGSERIAL SERIAL8 values.

Note that sqlca.sqlerrd[2] is defined as an INTEGER and therefore cannot be used to get the
last generated serial. To retrieve the last generated BIGSERIAL or SERIAL8, you must use the
`dbinfo()` SQL function as in the following code example:

```
MAIN
  DEFINE new_val BIGINT
  INSERT INTO mytable VALUES ( 0, 'aaaa' )
  SELECT dbinfo('bigserial') INTO new_val
    FROM systables WHERE tabid=1
  DISPLAY new_val
END MAIN
```

## Related links

**Related concepts**  

[The sqlca diagnostic record](0988-the-sqlca-diagnostic-record.md "The sqlca variable is a predefined record containing SQL statement execution information.")

[SERIAL and BIGSERIAL data types](1377-serial-and-bigserial-data-types.md "SERIAL and BIGSERIAL data types")
