---
title: "Transcribe Scanned Documents"
weight: 9
description: my description.
summary: a useful page.
lastmod: 2023-03-14
date: 2023-03-14
tags: []
categories: []
series: []
keywords: []
draft: true
---

On Ubuntu, ImageMagick should be installed out of the box.
Unfortunately it's policy does not allow you to read PDFs by default.

First switch to sudo:
```
$ sudo su -
```

Take a backup of the policy before editing it:
cp /etc/ImageMagick-6/policy.xml /etc/ImageMagick-6/policy.xml.backup

edit the policy xml:
$ /etc/ImageMagick-6/policy.xml

You should find a line like this:
<policy domain="coder" rights="none" pattern="PDF" />

Replace it with this:
<policy domain="module" rights="read" pattern="PDF" />

You also need tesseract to be installed:
$ apt-get install -y tesseract-ocr

You can exit sudo to return to your own user.

Now there are two steps to convert your scanned PDFs:

$ convert -density 300 file.pdf -depth 8 file.tiff  
$ tesseract file.tiff output

You can replace the file names.  'output' is a single text file.
