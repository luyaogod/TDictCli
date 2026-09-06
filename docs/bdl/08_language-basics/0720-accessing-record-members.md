---
title: "Accessing record members"
source: "fgl-topics/c_fgl_records_008.html"
breadcrumb: "Language basics > Records > Accessing record members"
type: "concept"
---

# Accessing record members

> Record members are accessed with the dot notation.

Record members are accessed by a dot notation.

The notation `record-name.member-name`
refers to an individual member of a
record:

```
DISPLAY rec.cust_name TO FORMONLY.f_name -- Single record member
```

The notation `record-name.*` refers to the entire list of
record members.

```
DISPLAY BY NAME rec.*  -- All record members
```

The notation `record-name.start-member THRU
record-name.end-member` refers to a consecutive set of members.
(`THROUGH` is a synonym for [`THRU`](0724-thru-through.md "The THRU keyword can be used to specify a set of members of a record.")):

```
DISPLAY rec.cust_id THRU rec.cust_address  -- Members from cust_id to cust_address
```

For complex record structures, sub-records can be accessed by chaining the dot-member
notation:

```
DEFINE reader RECORD
            id INTEGER,
            name VARCHAR(100),
            address RECORD
                num VARCHAR(20),
                street VARCHAR(200),
                city_id INTEGER
            END RECORD
       END RECORD

MAIN
    LET reader.name = "Mike Finley"
    LET reader.address.num = "2A"
    LET reader.address.street = "Sunset Bld"
END MAIN
```

## Related links

**Related concepts**  

[LET](0701-let.md "The LET statement assigns values to variables.")
