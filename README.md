# Go Workshop

Laast year started using quite a lot of Golang and was really enjoying it. At the time I thought my Go usage would increase, but in the end I left the service alone as it worked perfectly, and ended up doing 100% Rails for the last few months.

This workshop is going to be a learning experience for us all.

## Schedule

**Part One**

1. Setting up Go GB
2. Installing the dependencies before the Internet dies
3. Using Gorrila Mux to route requests to handlers
4. Basic hello world in JSON

**Part Deux**

1. Installing elasticsearch and managing with forego
2. Importing data into elasticsearch and getting it back out via postman
  - http://blog.qbox.io/multi-field-partial-word-autocomplete-in-elasticsearch-using-ngrams
3. Switch to using elasticseatch go client
4. Output hardcodes elasticsearch search to logs

**Part Three**

1. Format that output as JSON and check in the browser
  - respond with JSON
  - Accept status codes
2. Switch to using query strings
3. Handling errors with HTTP for missing content

**Part Four**

1. Calling from PHP
2. HTTP Middlewares ala http://www.alexedwards.net/blog/making-and-using-middleware and  https://gist.github.com/alexedwards/6f9496caecb2996ac61d
3. Using `.env` for useful settings
