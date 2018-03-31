# Introduction

[${WORK}](http://www.splunk.com) has me writing microservices in Go,
using OpenAPI 2.0 / Swagger.  While I'm not a fan of Go (that's a bit of
an understatement) I get why Go is popular with enterprise managers, it
does exactly what it says it does.  It's syntactically hideous.  I'm
perfectly happy taking a paycheck to write in it, and I'm pretty good at
it already.  I just wouldn't choose it for a personal project.

But if you're writing microservices for enterprise customers, yes, you
should use Go, and yes, you should use OpenAPI and Swagger.  So here's
how it's done.

[Swagger](https://swagger.io/) is a specification that describes the
ndpoints for a webserver's API, usually a REST-based API.  HTTP uses
verbs (GET, PUT, POST, DELETE) and endpoints (/like/this) to describe
things your service handles and the operations that can be performed
against it.

Swagger starts with a **file**, written in JSON or YAML, that names
each and every endpoint, the verbs that endpoint responds to, the
parameters that endpoint requires and takes optionally, and the
possible responses, with type information for every field in the
inputs and outputs.

Swagger *tooling* then takes that file and generates a server ready to
handle all those transactions.  The parameters specified in the
specification file are turned into function calls and populated with
"Not implemented" as the only thing they return.

## Your job

In short, for a basic microservice, it's your job to replace those
functions with your business logic.

There are three things that are your responsibility:

1. Write the specification that describes *exactly* what the server
accepts as requests and returns as responses.

2. Write the business logic.

3. Glue the business logic into the server generated from the
specification.

In Go-Swagger, there is *exactly one* file in the generated code that
you need to change.  Every other file is labeled "DO NOT EDIT."  This
one file, called `configure_project.go`, has a top line that says "This
file is safe to edit, and will not be replaced if you re-run swagger."
That *exactly one* file should be the only one you ever need to change.

## The setup

You'll need Go.  I'm not going to go into setting up Go on your system;
there are
[perfectly adequate guides elsewhere](https://golang.org/doc/install).
You will need to install `swagger` and `dep`.

Once you've set up your Go environment (set up $GOPATH and $PATH), you
can just:

```
$ go get -u github.com/golang/dep/cmd/dep
$ go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

## Initialization

Now you're going to create a new project.  Do it in your src directory
somewhere, under your $GOPATH.

```
$ mkdir project timeofday
$ cd timeofday
$ git init && git commit --allow-empty -m "Init commit: Swagger Time of Day."
$ swagger init spec --format=yaml --title="Timeofday" --description="A silly time-of-day microservice"
```

You will now find a new swagger file in your project directory.  If
you open it up, you'll see short header describing the basic features
Swagger needs to understand your project.

## Operations

Swagger works with **operations**, which is a combination of a verb
and an endpoint.  We're going to have two operations which do the same
thing: return the time of day.  The two operations use the same
endpoint, but different verbs: GET and POST.  The GET argument takes
an optional timezone as a search option; the POST argument takes an
optional timezone as a JSON argument in the body of the POST.

First, let's version our API.  You do that with Basepaths:

<<version the API>>=
basePath: /timeofday/v1
@

Now that we have a base path that versions our API, we want to define
our endpoint.  The URL will ultimately be `/timeofday/v1/time`, and we
want to handle both GET and POST requests, and our responses are going
to be **Success: Time of day** or **Timezone Not Found**.  

<<define the paths>>=
paths:
  /time:
    get:
      operationId: "GetTime"
      parameters:
        - in: path
          name: Timezone
          schema:
            type: string
            minLength: 3
      responses:
        200:
          schema:
            $ref: "#/definitions/TimeOfDay"
        404:
          schema:
            $ref: "#/definitions/NotFound"
    post:
      operationId: "PostTime"
      parameters:
        - in: body
          name: Timezone
          schema:
            $ref: "#/definitions/Timezone"
      responses:
        200:
          schema:
            $ref: "#/definitions/TimeOfDay"
        404:
          schema:
            $ref: "#/definitions/NotFound"
@

The `$ref` entries are a YAML thing for referring to something else.
The octothorpe symbol `(#)` indicates "look in the current file.  So
now we have to create those paths:

<<schemas>>=
definitions:
  NotFound:
    type: object
    properties:
      code:
        type: integer
      message:
        type: string
  Timezone:
    type: object
    properties:
      Timezone:
        type: string
        minLength: 3
  TimeOfDay:
    type: object
    properties:
      TimeOfDay: string
@

This is *really verbose*, but on the other hand it is *undeniably
complete*: these are the things we take in, and the things we respond
with.

So now your file looks like this:

<<swagger.yml>>=
swagger: "2.0"
info:
  version: 0.1.0
  title: timeofday
produces:
  - application/json
consumes:
  - application/json
schemes:
  - http

# Everything above this line was generated by the swagger command.
# Everything below this line you have to add:

<<version the API>>

<<schemas>>

<<define the paths>>
@

Now that you have that, it's time to generate the server!  

`$ swagger generate server -f swagger.yml`

It will spill out the actions it takes as it generates your new REST
server.  **Do not** follow the advice at the end of the output.
There's a better way.

`$ dep init`

Dependency management in Go is a bit of a mess, but the accepted
solution now is to use `dep` rather than `go get`.  This creates a
pair of files, one describing the Go packages that your file uses, and
one describing the exact *versions* of those packages that you last
downloaded and used in the `./vendor/` directory under your project
root.

Now you can build the server:

`$ go build ./cmd/timeofday-server/`

And then you can run it.  Feel free to change the port number:

`$ ./timeofday-server --port=8082`

You can now tickle the server:

```
$ curl http://localhost:8082/
{"code":404,"message":"path / was not found"}
$ curl http://localhost:8082/timeofday/v1/time
" function .GetTime is not implemented"
```

Congratulations!  You have a working REST server that does, well,
nothing.

For part two, we'll make our server actually do things.
