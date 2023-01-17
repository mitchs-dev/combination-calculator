# Combination Calculator

## What is it?

This simple calculator provides all possible combinations, avoiding comparisons to duplicate values as well as priority on the first available match (I.e 1 + 4 would take priority over 4 + 1 matching.)

Here is an example of valid values if you are looking for all possible combinations of "4" values:

| <span style="font-weight:normal"></span> | <span style="font-weight:normal">1</span> | <span style="font-weight:normal">2</span> | <span style="font-weight:normal">3</span> | <span style="font-weight:normal">4</span> |
| -------------- | -------------- | -------------- | -------------- | -------------- |
| 1 | ❌ | ✅ | ✅ | ✅ |
| 2 | ❌ | ❌ | ✅ | ✅ |
| 3 | ❌ | ❌ | ❌ | ✅ |
| 4 | ❌ | ❌ | ❌ | ❌ |

This provides you with **6** possible combinations.

## Weight

An additional feature of this tool is the `--weight` flag which allows you to specify a percentage to allow weight to these values.

For example, if you were to set `--weight 50` and `--value 100` this means that 50% of the values would be discarded. It's important to note that which combinations are discarded are determined after the initial prioritization and combinations are generated and the discarded values are, in theory, "random" as they are not _actually_ calculated.

A few additional details:

* The percentage can be any positive standard percentage integer value between. (0-100)
* The final weight is rounded to the nearest whole number.
* You can use a decimal value or a whole number value.

Here is an example of valid values if you are looking for combinations for a value of "4" with a 50% weight:

| <span style="font-weight:normal"></span> | <span style="font-weight:normal">1</span> | <span style="font-weight:normal">2</span> | <span style="font-weight:normal">3</span> | <span style="font-weight:normal">4</span> |
| -------------- | -------------- | -------------- | -------------- | -------------- |
| 1 | ❌ | ✅ | ❌ | ❌ |
| 2 | ❌ | ❌ | ✅ | ❌ |
| 3 | ❌ | ❌ | ❌ | ✅ |
| 4 | ❌ | ❌ | ❌ | ❌ |

This provides you with **3** possible combinations.

## Flags

There are a few flags available that do various things 

| Name | Description | Required | Usage |
| :-----: | :-----: | :-----: | :-----: |
| value |  Value to check combinations for | ✅ | `--value <NUMBER>` |
| weight |  Specify a percentage to adjust the outcome to the highest rounded whole number | ❌ | `--weight <NUMBER>` |
| silent | Only return count | ❌ | `--silent` |
| debug |  Enable debug mode | ❌ | `--debug` |
| version |   Shows version of this tool | ❌ | `--version` |

## Example

```golang
$ ./combination-calculator --value 10 --weight 25
INFO[0000] Completed Cycle (1/10)                       
INFO[0000] Completed Cycle (2/10)                       
INFO[0000] Completed Cycle (3/10)                       
INFO[0000] Completed Cycle (4/10)                       
INFO[0000] Completed Cycle (5/10)                       
INFO[0000] Completed Cycle (6/10)                       
INFO[0000] Completed Cycle (7/10)                       
INFO[0000] Completed Cycle (8/10)                       
INFO[0000] Completed Cycle (9/10)                       
INFO[0000] Completed Cycle (10/10)                      
----------------------------------
Total Combinations:  34
```

## Disclaimer

This was created as a personal "hackathon" project and it wasn't ever intended to be maintained or updated and will be provided as is.