---
title: "Example 2: Using the Apache POI framework"
source: "fgl-topics/c_fgl_JavaBridge_example_2.html"
breadcrumb: "Extending the language > The Java interface > Examples > Example 2: Using the Apache POI framework"
type: "concept"
description: "This example shows how to create a XLS file, using the Apache POI framework . Note: This demo requires Apache POI version 4.0.0 or higher. Download and install the Apache POI package, set CLASSPATH ..."
---

# Example 2: Using the Apache POI framework

This example shows how to create a XLS file, using the [Apache POI framework](http://poi.apache.org).

> **Note:**
>
> This demo requires Apache POI version 4.0.0 or higher.

Download and install the Apache POI package, set CLASSPATH environment variable to point to the
POI JAR archives.

After execution, a file named "itemlist.xls" is found in the current
directory.

```
-- Needs Apache POI 4.0.0 +
IMPORT JAVA java.io.FileOutputStream
IMPORT JAVA org.apache.poi.hssf.usermodel.HSSFWorkbook
IMPORT JAVA org.apache.poi.hssf.usermodel.HSSFSheet
IMPORT JAVA org.apache.poi.hssf.usermodel.HSSFRow
IMPORT JAVA org.apache.poi.hssf.usermodel.HSSFCell
IMPORT JAVA org.apache.poi.hssf.usermodel.HSSFCellStyle
IMPORT JAVA org.apache.poi.hssf.usermodel.HSSFFont
IMPORT JAVA org.apache.poi.ss.usermodel.IndexedColors
IMPORT JAVA org.apache.poi.ss.usermodel.HorizontalAlignment
IMPORT JAVA org.apache.poi.ss.usermodel.FillPatternType
IMPORT JAVA org.apache.poi.ss.usermodel.CellType

MAIN
    DEFINE fo FileOutputStream
    DEFINE workbook HSSFWorkbook
    DEFINE sheet HSSFSheet
    DEFINE row HSSFRow
    DEFINE cell HSSFCell
    DEFINE style HSSFCellStyle
    DEFINE headerFont HSSFFont
    DEFINE i, id INTEGER, s STRING

    LET workbook = HSSFWorkbook.create()

    LET style = workbook.createCellStyle()
    CALL style.setAlignment(HorizontalAlignment.CENTER)
    CALL style.setFillForegroundColor(
                 IndexedColors.LIGHT_CORNFLOWER_BLUE.getIndex())
    CALL style.setFillPattern(FillPatternType.SOLID_FOREGROUND)
    LET headerFont = workbook.createFont()
    CALL headerFont.setBold(TRUE)
    CALL style.setFont(headerFont)

    LET sheet = workbook.createSheet()

    LET row = sheet.createRow(0)
    LET cell = row.createCell(0)
    CALL cell.setCellValue("Item Id")
    CALL cell.setCellStyle(style)
    LET cell = row.createCell(1)
    CALL cell.setCellValue("Name")
    CALL cell.setCellStyle(style)

    FOR i=1 TO 10
        LET row = sheet.createRow(i)
        LET cell = row.createCell(0)
        CALL cell.setCellType(CellType.NUMERIC)
        LET id = 100 + i
        CALL cell.setCellValue(id)
        LET cell = row.createCell(1)
        LET s = SFMT("Item #%1",i)
        CALL cell.setCellValue(s)
    END FOR

    LET fo = FileOutputStream.create("itemlist.xls")
    CALL workbook.write(fo)
    CALL fo.close()

END MAIN
```
