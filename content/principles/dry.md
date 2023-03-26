---
title: Don't Repeat Yourself
weight: 2
draft: false
summary: Minimize the code duplication.
lastmod: 2023-03-24
---

Software engineers often feel rushed to deliver, and when it comes time to add a new
block of code, they might just copy something that works and make a small tweak to it.
The code works so everybody is happy.

The problem comes later when it is discovered that a bug was in that code.  They promptly
fix it in the original, but they forgot that it was copied, or the person fixing the 
bug is not the same person who copied it before.  So now instead of suffering from a 
bug once, you are suffering from it multiple times.

This also goes against the KISS principle, duplicating code increases complexity, and the 
chance of bugs being present.

As a principle, "Don't Repeat Yourself" (DRY) is self explanatory.  You should remove as
much code duplication as you can.

# Modularity

If you find a block of code that is highly similar to others, except for one small aspect, 
then make a new function for it.  Use function arguments to control the aspect that
differentiated the code before.

Detecting duplicated code should not be done by manual review, there are several command
line programs that are able to detect code duplication, such as CPD, dupl, and Sonar.
I recommend running these tools whenever you review a pull request.