---
title: "Cut"
draft: false
weight: 3
---

* Assuming you just want to select a column: use awk 
* if you want consecutive delimiters (like spaces) to count as one.  If you have text with multiple spaces between values, awk is your friend.
* use cut if there is only one delimiter (like a comma or pipe).

in cut, fields start from 1.
you can have a space between the option and its argument.
use -f for fields.

Example:
```
alpha,Charlie,85393
beta,Megan,98523
alpha,Billy,43652
omega,Zach,63425
helios,Shaq,82352
zed,Shaniah,153523
faraday,Nicole,89623
helios,Ken,62632
omega,Adam,98765
beta,Sean,62436
alpha,Roger,56147
alpha,Frank,89813
omega,Daniel,120151
beta,Jack,100101
```

```
$ cut -d ',' -f 2- x.csv 
Charlie,85393
Megan,98523
Billy,43652
Zach,63425
Shaq,82352
Shaniah,153523
Nicole,89623
Ken,62632
Adam,98765
Sean,62436
Roger,56147
Frank,89813
Daniel,120151
Jack,100101
```