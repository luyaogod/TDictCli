---
title: "THRU/THROUGH"
source: "fgl-topics/c_fgl_records_THRU.html"
breadcrumb: "Language basics > Records > THRU/THROUGH"
type: "concept"
---

# THRU/THROUGH

> The THRU keyword can be used to specify a set of members of a record.

## Syntax

```
record.first-member [ THRU | THROUGH ] record.last-member
```

1. record defines the record to be used.
2. first-member defines the member of the record starting the group of variables.
3. last-member defines the member of the record ending the group of variables.
4. `THROUGH` is a synonym for `THRU`.

## Usage

The `THRU` keyword can be used in several instructions such as [`DISPLAY`](../11_user-interface/1878-display-to-stdout.md "The DISPLAY instruction displays text in line mode to the standard output channel."), [`PRINT`](../12_reports/2491-print.md "Formats and prints a row of data in a report routine."), [`INITIALIZE`](0698-initialize.md "The INITIALIZE instruction initializes program variables with NULL or default values."), [`VALIDATE`](0702-validate.md "The VALIDATE instructions checks a variable value based on database schema validation rules."), [`LOCATE`](0699-locate-for-text-byte.md "The LOCATE statement specifies where to store data of TEXT and BYTE variables."), to specify a list of record
members by using a starting member and an ending member.

## Example

```
SCHEMA stores
MAIN
  DEFINE cust RECORD LIKE customer.*
  INITIALIZE cust.customer_num THRU cust.phone TO NULL
END MAIN
```

## Related links

**Related concepts**  

[RECORD](0717-record.md "The RECORD keyword defines a structured type or variable.")
