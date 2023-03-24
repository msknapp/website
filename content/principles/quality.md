---
title: Judging Quality
summary: You should routinely use automated tools to assess your code quality.
weight: 9
description: Assessing code quality.
lastmod: 2023-03-18
---

Lots of people write code, and most of them think they are experts.  How do you know
that code is actually GOOD?  Often times, you may review somebody else's code, and 
you feel that it's poorly done.  If you want to criticize it, then you risk alienating
that co-worker, and convincing team members that you are difficult to work with.

# Constructive Feedback

It's important that, when you give feedback, you are not stating your own opinion.
If it sounds like your own crazy bull-headed opinion, people will just think you are 
stubborn.

So you need to make sure that:
* Your feedback can be tied to well established industry best practices and principles.
* Your feedback can be tied to documented requirements from product owners, management, or team members.
* Your concerns are measurable.
* Your concerns are material.  Don't block a PR because a variable name is too terse.

# Measurable Checks

To judge code, you should consider:
* Unit test coverage: This is measurable, most languages have tools to check coverage.
* Documentation: There doesn't need to be a ton, but if a section of code is confusing 
  or counter-intuitive, then you should insist on it being documented.  Add function 
  documentation, and a README.
* Low Duplication (DRY code): This can be measured, there are programs like "dupl" and "CPD" that are
  able to search for code blocks that are highly similar, like they were copied and pasted.
* Simplicity: This can be measured with "cyclomatic complexity", which Sonar can evaluate.
* Loose Coupling (single responsibility): This may be the most difficult one to measure or prove.
  It could also be the most neglected aspect, and the one most likely to affect long term maintenance.
  Look for cases where one class is taking responsibility for more things than necessary, 
  like if it manages database updates and authorization simultaneously.

It's also fair to say that anything reported in Sonar could be an objection.

# Immaterial Issues

If you are reviewing a pull request, please keep in mind that the developer's time is
valuable.  There are other stories to take on, and it might feel like it takes forever 
just to get their PRs reviewed.

Typically I leave **suggestions** for how code can be improved, but I will not block the
PR unless I think it is really necessary.  

Be sure that you do not reject PRs for things like the following:
* bad naming conventions.
* excessive spaces.
* failing to document every minor function.
* Not making code defensive in nature.