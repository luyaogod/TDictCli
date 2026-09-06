---
title: "base.Channel.write"
source: "fgl-topics/c_fgl_ClassChannel_write.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > base.Channel methods > base.Channel.write"
type: "concept"
---

# base.Channel.write

> Writes a list of data delimited by a separator to the channel.

## Syntax

> **Note:**
>
> In the next(s) syntax diagram(s), the `[ ] { } |`
> symbols are part of the syntax.

```
write( valueList )
```

where valueList can be one of:

- ```
  value
  ```
- ```
  [ value , ... ]
  ```
- ```
  record
  ```
- ```
  [ record.* ]
  ```

1. value is an expression of a [primitive
   data type](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.") such as `INTEGER`, `VARCHAR(50)`, etc.
2. record is a variable defined as a [`RECORD`](../08_language-basics/0715-records.md "Records allow structured program variables definitions.").
3. When there are multiple values to write, the value(s) must be specified between `[
   ]` square brackets. These are provided as [a
   variable parameter list](../08_language-basics/0643-variable-parameter-list.md "Variable parameter list delimiters.").
4. If only one value is to be written, you can specify the value without the `[ ]`
   brackets.
5. To write all values from a `RECORD` variable, the [`record.*`](../08_language-basics/0720-accessing-record-members.md "Record members are accessed with the dot notation.") notation can be
   used when surrounded by `[ ]` brackets, to expand all record fields as a list of
   simple values. Or, you can also directly specify the record name without the `.*`
   notation and no `[ ]` brackets.

## Usage

After opening a channel, use the `write()` method to write a
record of data to the channel.

The `write()` method uses the field delimiter defined by [`setDelimiter()`](2996-base-channel-setdelimiter.md "Define the value delimiter for a channel."). The delimiter
also defines serialization rules for example when using `"CSV"`.

The `write()` method takes a modifiable list of variables as the parameter.

Errors [-6344](4483-genero-bdl-errors.md) or [-6345](4483-genero-bdl-errors.md) are
thrown, if the channel fails to write data.

## Example

```
DEFINE cust_rec RECORD LIKE customer.*
...
CALL ch.write(cust_rec)  -- equivalent to ch.write([cust_rec.*])
```

## Related links

**Related concepts**  

[Read and write record data](3001-read-and-write-record-data.md "Read and write record data")
