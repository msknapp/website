---
title: Don't Reinvent the Wheel
draft: false
weight: 59
lastmod: 2023-03-29
summary: Avoid duplication of efforts by leveraging open source.
nextlink: /reliability/
---

There is a lot of code already written and public, often with very permissive licenses.
If you are writing something similar to an existing open source project, then you are
probably wasting your time.  The open source project will have more contributors, 
support, maturity, testing, and features than you can possibly keep up with.  It's better
to make contributions to the open source project as you need them.  Include contributions
that make the open source project more extensible too, rather than directly 
contributing code to its core business logic.  This is in line with the open/closed
principle.

Eventually, as the open source project matures, you might be forced to drop your own
library anyways in favor of the open source one.  Your efforts will have been for nothing.

**It is very important for professional software engineers to be aware of major open
source projects and libraries that are relevant to their efforts.  They should search
for these libraries at least once every quarter year.**

Likewise, you should search for company wide common projects or libraries.  The
collaboration here can result in improved performance of all teams.  If your company
doesn't have common tools or libraries, an organization for this should be 
recommended or suggested to management.  It can make a bid difference.  Starting
one on your own won't make much of a difference unless you can get it broadly
known and supported across the company.  There really needs to be a dedicated
team to support common shared libraries.

# Evaluating Open Source

Of course, you also need to be careful evaluating these projects.  Consider the following:
* Is it easily extended without direct contributions?
* What is the license it uses?  Does your company allow that?
* How much unit test coverage does it have?
* How many contributors does it have?
* How many forks or stars does it have?
* How many dependencies does it have?  Fewer tends to be better.
* How old is the most recent commit?  If it's quite old, it might not be supported in the future.
* How thorough is its documentation?  Nobody wants to use a library that is poorly documented.
* Are there vulnerability scans on the library?
* Is it in Sonar?  Or otherwise evaluated for code quality?
* Does it build on your local machine?
* Does it have tools built around it?
* Does it have its own website?
* Is there a company or foundation that is officially supporting it?

Lastly, you need to perform a spike and proof of concept with the project before committing to it.
If you commit to a project that ultimately does not work for the team, that can seriously
impact your ability to deliver future results, and damage your team's reputation.

# Builder Bias

Software engineers tend to feel very proud of the work they have done.  This can lead them to 
insist on keeping it, in spite of better options existing.  Engineers should not feel so attached
to the code that they have written.  Code is often built with a short shelf like, like
half of code is deleted within a year.  If you leave the team, there is a fair chance your
work will just be deleted in the near future anyways.  This is not to say engineers were unproductive, it 
might have been needed in that time.

My advice is to not take it personally if your own code gets deleted one day, it doesn't 
imply you did bad work.  Always embrace the simpler and more mature or robust option,
even if it means abandoning some of your own good work.  You should always be willing to 
abandon or delete your own code if a better solution is found.

# Fear of New Things

Another reason software engineers write their own solutions and ignore existing solutions,
is that they don't want to learn new things.  If they write the code, then they understand it,
and don't need to struggle with learning something else.  This doesn't help much, because
your co-workers still need to learn your work, except yours is more likely to be 
abandoned within one or two years.

Generally, **software engineers should be willing and eager to learn new software and
technology at any time.**  They should not shy away from new projects or libraries
simply because they are not familiar with them.