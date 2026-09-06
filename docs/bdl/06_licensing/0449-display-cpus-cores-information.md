---
title: "Display CPUs (cores) information"
source: "genero-install-topics/t_license_controller_display_cpu_info.html"
breadcrumb: "Licensing > Manage Genero BDL local license > Displaying environment and statistics > Display CPUs (cores) information"
type: "task"
---

# Display CPUs (cores) information

> Use this procedure on the application server where the Genero runtime is installed, to determine the exact number of CPUs (cores/threads) for CPU licensing.

To display the number of logical CPUs in your machine type the command:

```
fglWrt -a cpu
```

The number of CPUs (cores/threads) you have are displayed. An example of a
machine which has one physical CPU with four cores is shown.

![Image shows the number of CPUs output from the command fglWrt -a cpu](../_images/fglwrt_cpu_number.png)

*Display the Number of CPUs using License Controller (fglWrt)*

## Related links

**Related concepts**  

[Displaying environment and statistics](0445-displaying-environment-and-statistics.md "When the Four Js License Manager is used for licensing, it works with the license controller locally to access information. The topics in this section describe license controller options that apply to the use of the FLM.")

**Related reference**  

[License details reference](0451-license-details-reference.md "A reference to license details.")
