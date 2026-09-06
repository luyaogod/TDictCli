---
title: "base.Channel.read"
source: "fgl-topics/c_fgl_ClassChannel_read.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > base.Channel methods > base.Channel.read"
type: "concept"
---

# base.Channel.read

> Reads a list of data delimited by a separator from the channel.

## Syntax

> **Note:**
>
> In the next(s) syntax diagram(s), the `[ ] { } |`
> symbols are part of the syntax.

```
read( variableList )
  RETURNS INTEGER
```

where variableList can be one of:

- ```
  variable
  ```
- ```
  [ variable , ... ]
  ```
- ```
  record
  ```
- ```
  [ record.* ]
  ```

1. variable is a program variable of a [primitive data type](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.") such as `INTEGER`, `VARCHAR(50)`,
   etc.
2. record is a variable defined as a [`RECORD`](../08_language-basics/0715-records.md "Records allow structured program variables definitions.").
3. When there are multiple variables to read, the variable(s) must be specified between `[
   ]` square brackets. These are provided as [a
   variable parameter list](../08_language-basics/0643-variable-parameter-list.md "Variable parameter list delimiters.").
4. If only one variables is to be read, you can specify the variable without the `[
   ]` brackets.
5. To read all values into a `RECORD` variable, the [`record.*`](../08_language-basics/0720-accessing-record-members.md "Record members are accessed with the dot notation.") notation can be
   used when surrounded by `[ ]` brackets. Or, you can also directly specify the record
   name without the `.*` notation and no `[ ]` brackets.

## Usage

After opening the channel object, use the `read()` method to read a record of
data from the channel.

The `read()` method uses the field delimiter defined by [`setDelimiter()`](2996-base-channel-setdelimiter.md "Define the value delimiter for a channel."). The delimiter
also defines deserialization rules for example when using `"CSV"`.

The `read()` method takes a modifiable list of variables as parameter.

A call to `read()` is blocking until the read operation is complete.

If the `read()` method returns less data than expected, then the remaining
variables will be initialized to `NULL`. If the `read()` method
returns more data than expected, the data is silently ignored.

Any target variable must have a [primitive type](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data."), or
be a `RECORD` that contains only members defined with a primitive type.

If data is read, the `read()` method returns [`TRUE`](../08_language-basics/0573-true.md "TRUE is a predefined constant to be used in boolean expressions."). Otherwise, it returns [`FALSE`](../08_language-basics/0574-false.md "FALSE is a predefined constant to be used in boolean expressions."), indicating the end of the
file or stream.

> **Important:**
>
> Files encoded in UTF-8 can start with the UTF-8 Byte Order Mark (BOM), a sequence of `0xEF
> 0xBB 0xBF` bytes, also known as UNICODE `U+FEFF`. When reading files, Genero
> BDL will ignore the UTF-8 BOM, if it is present at the beginning of the file. This applies to
> instructions such as `LOAD`, as well as I/O APIs such as
> `base.Channel.read()` and `readLine()`.

Error [-6346](4483-genero-bdl-errors.md) is thrown, if the channel fails to read data.

## Example

```
DEFINE cust_rec RECORD LIKE customer.*
...
WHILE ch.read(cust_rec) -- equivalent to: ch.read([cust_rec.*])
  ...
END WHILE
```

For a complete example, see [Example 1: Using record-formatted data file](3009-example-1-using-record-formatted-data-file.md).

## Related links

**Related concepts**  

[Read and write record data](3001-read-and-write-record-data.md "Read and write record data")

[base.Channel.dataAvailable](2985-base-channel-dataavailable.md "Tests if some data can be read from the channel.")
