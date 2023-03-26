
# How to Build
* build the static site: `hugo`
* test locally: `hugo serve -p 8080`
  * then just open: `http://localhost:8080/`

# How to Publish

* Log in to the aws console.  
  * Get the password: `pass show aws/root`
* Setup the aws CLI.
  * `printenv | grep AWS`
  * If that prints the access and secret key, it is ready.
* Clean:
  * aws s3 rm --recursive 's3://msknapp.com/'
* Publish:
  * aws s3 cp --recursive public/ 's3://msknapp.com/'
* S3 URL:
  * http://msknapp.com.s3-website-us-east-1.amazonaws.com/

# Vital Settings

In AWS, the site is made public by these settings.
* In S3, the msknapp.com bucket:
  * In 'properties', static website hosting is enabled.
* In Route53, 'msknapp.com' is an A record that is an alias pointing to 
  value: 's3-website-us-east-1.amazonaws.com.' and record name 'msknapp.com'
  

# Problems

* next/previous should be available from sections, and start/end of sections.
* left menu should show pages and sections within the current top section, you should
  be able to expand them to see sub-pages within.
* Menus are not right.
* should be able to print pages.
* should be able to view from tablets or phones.
* if in a section, the side menu should be for that section.
  * current page should have its own color.
* Needs tags, and related content. 
* Needs a taxonomy.
* Needs a search bar
* Needs metadata shown on page.

# Fixed
* Needs back and forth buttons.  Also an up button.
* menus should be lists.
* need bread-crumbs to work.
* Pages could have table of contents.