[Docs Site](https://msknapp.github.com/website)

# How to Build

## For Public Website

* In config.toml, set the parameter "root" to an empty string or remove it.
* build the static site: `hugo`
* test locally: `hugo serve -p 8080`
  * then just open: `http://localhost:8080/`

## For Github Pages

* `hugo -d docs -b /website/`
* Note: the config.toml file has a parameter called "root", which prefixes all URLs.
  This must match the repository name.

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

* left menu should show pages and sections within the current top section, you should
  be able to expand them to see sub-pages within.
* Needs tags, and related content. 
* Needs a taxonomy.
* Needs a search bar

# Plans

* applications for:
  * fibonacci
  * large factorials
  * permutations, combinations
  * evaluating expressions on documents.
* Graphical interactions with javascript and wasm.
* documentation on:
  * mastering the api server
  * SQL
  * Designing relational tables.
  * Postgres
  * grpc, protobuff
  * Using Gin
  * Cloud architecture.
  * google cloud
  * bash scripting
  * Kubernetes
  * golang concurrency (tidy the docs)
  * golang gotchas.
* add in react

# Plans Before Starting Job

* Practice GRPC and document it.
* Learn React.
* GCP architect certification.

Lower priority:
* Review gin
* API best practices.
* learn angular.
* learn Rust

# Code to html

`code-to-html -l golang < /path/to/file | xclip -selection clipboard`
