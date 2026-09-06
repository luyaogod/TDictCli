---
title: "Unique and primary key constraint violation"
source: "fgl-topics/c_fgl_Migrate_to_221_pkey_constraint.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.21 upgrade guide > Unique and primary key constraint violation"
type: "concept"
---

# Unique and primary key constraint violation

> Unique and primary key constraint violations mostly return error -268. However, error -269 may be checked too.

When a unique or primary key constraint is violated, the IBM®
Informix® driver returns the error -268
in sqlca.sqlcode if the database uses transaction logging, and error -239 if the
database uses no logging.

Regarding non-Informix drivers, all 2.21 drivers now return -268 when a unique
constraint or primary key constraint is violated. Before 2.21, the Oracle and SQL Server drivers
returned error -239, which is only returned by IBM Informix databases without transaction logging.
Returning error -268 for all drivers is the best choice in a context of
transactional databases.

Check your code for -239 error code usage and replace by -268.
If you still need to test error -239 (for example because you have IBM
Informix databases without transactions), we recommend
that you write a function testing different error codes to check unique constraint
violation:

```
FUNCTION isUniqueConstraintError()
  IF (sqlca.sqlcode==-239 OR sqlca.sqlcode==-268)
  OR (sqlca.sqlcode==-346 AND sqlca.sqlerrd[2]==-100)
  THEN
    RETURN TRUE
  ELSE
    RETURN FALSE
  END IF
END FUNCTION
```
