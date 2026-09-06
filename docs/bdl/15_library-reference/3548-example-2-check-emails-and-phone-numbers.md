---
title: "Example 2: Check emails and phone numbers"
source: "fgl-topics/c_fgl_ext_util_Regexp_example_2.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Regexp class > Examples > Example 2: Check emails and phone numbers"
type: "concept"
description: "The following program uses regular expressions to verify that a string follows a given format such as email addresses and phone numbers: IMPORT util MAIN DISPLAY ..."
---

# Example 2: Check emails and phone numbers

The following program uses regular expressions to verify that a string follows
a given format such as email addresses and phone numbers:

```
IMPORT util
MAIN

    DISPLAY isValidEmailAddressASCII("mike.torn@4js.com")
    DISPLAY isValidEmailAddressASCII("MIKE.TORN@4JS.COM")
    DISPLAY isValidEmailAddressASCII("mike.torn@4js")
    DISPLAY isValidEmailAddressASCII("@4js.com")

    DISPLAY isValidPhoneNumberFrance("+33 6 99 88 77 66")
    DISPLAY isValidPhoneNumberFrance("+33 6 99 88 77")
    DISPLAY isValidPhoneNumberFrance("33 6 99 88 77 66")

    DISPLAY isValidPhoneNumberUSA("+1 (213) 887-7676")
    DISPLAY isValidPhoneNumberUSA("+1 (213)")

END MAIN

FUNCTION isValidEmailAddressASCII(s STRING) RETURNS BOOLEAN
    RETURN util.Regexp.compile(`^[a-zA-Z0-9._%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`).matches(s)
END FUNCTION

FUNCTION isValidPhoneNumberFrance(s STRING) RETURNS BOOLEAN
    RETURN util.Regexp.compile(`^\+33 [\d] [\d]{2} [\d]{2} [\d]{2} [\d]{2}$`).matches(s)
END FUNCTION

FUNCTION isValidPhoneNumberUSA(s STRING) RETURNS BOOLEAN
    RETURN util.Regexp.compile(`^\+1 \([2-9][0-9][0-9]\) [0-9]{3}-[0-9]{4}$`).matches(s)
END FUNCTION
```
