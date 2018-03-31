# Review of Part One

In [Part One of Go-Swagger](TK:), we generated a on OpenAPI 2.0 server
with REST endpoints.  The server builds and responds to queries, but
every valid query ends with "This feature has not yet been
implemented."

It's time to implement the feature.

I want to emphasize that with Go Swagger there is *only* one generated
file you need to touch.  Since our project is named `timezone`, the
file will be named `restapi/configure_timezone.go`.  Our first step
will be to break those "not implemented" functions out into their own
Go package.  That package will be our business logic.  The configure
file and the business logic package will be the *only* things we
change.

## Break out the business logic

Create a new folder in your project root and call it `timeofday`.

Open up your editor and find the file `restapi/configure_timeofday.go`.
In your `swagger.yml` file you created two endpoints and gave them each
an `operationId`: `TimekPost` and `TimeGet`.  Inside
`configure_timeofday.go`, you should find two corresponding assignments
in the function `configureAPI()`: `TimeGetHandlerFunc` and
`ClockPostHandlerFunc`.  Inside those function calls, you'll find
anonymous functions.

I want you to take those anonymous functions, cut them out, and paste
them into a new file inside the `timeofday/` folder.  You will also have
to create a package name and import any packages being used.  Now your
file, which I've called `timeofday/handlers.go`, looks like this:

<<handlers.go before implementation>>=
package timeofday

import(
  "github.com/go-openapi/runtime/middleware"
  "github.com/elfsternberg/timeofday/restapi/operations"
)

func GetTime(params operations.TimeGetParams) middleware.Responder {
  return middleware.NotImplemented("operation .TimeGet has not yet been implemented")
}

func PostTime(params operations.TimePostParams) middleware.Responder {
  return middleware.NotImplemented("operation .TimePost has not yet been implemented")
}
@

And now go back to `restapi/configure_timeofday.go`, add
`github.com/elfsternberg/timeofday/clock` to the imports, and change the
handler lines to look like this:

<<configuration lines before implementation>>=
	api.TimeGetHandler = operations.TimeGetHandlerFunc(timeofday.GetTime)
	api.TimePostHandler = operations.TimePostHandlerFunc(timeofday.PostTime)
@

## Implementation

Believe it or not, you've now done everything you need to do except the
business logic.  We're going to honor the point of OpenAPI and the `//
DO NOT EDIT`` comments, and not modify anything exceept the contents of
our handler.

To understand our code, though, we're going to have to *read* some of
those files.  Let's go look at `/models`.  In here, you'll find the
schemas you outlined in the `swagger.yml` file turned into source code.
If you open one, like many files generated by Swagger, you'll see it
reads `// DO NOT EDIT`.  But then there's that function there,
`Validate()`.  What if you want to do advanced validation for custom
patterns or inter-field relations not covered by Swagger's validators?

Well, you'll have to edit this file.  And figure out how to live with
it.  We're not going to do that here.  This exercise is about *not*
editing those files.  But we can see, for example, that the `Timezone`
object has a field, `Timezone.Timezone`, which is a string, and which
has to be at least three bytes long.

The other thing you'll have to look at is the `restapi/operations`
folder.  In here you'll find GET and POST operations, the parameters
they accept, the responses they deliver, and lots of functions only
Swagger cares about.  But there are a few we care about.

Here's how we craft the GET response.  Inside `handlers.go`, you're
going to need to extract the requested timezone, get the time of day,
and then return either a success message or an error message.  Looking
in the operations files, there are a methods for good and bad returns,
as we described in the swagger file.

<<gettime implementation>>=
func GetTime(params operations.TimeGetParams) middleware.Responder {
	var tz *string = nil

	if (params.Timezone != nil) {
		tz = params.Timezone
	}

	thetime, err := getTimeOfDay(params.Timezone)

@

The first thing to notice here is the `params` field: we're getting a
customized, tightly bound object from the server.  There's no hope of
abstraction here.  The next is that we made the Timezone input optional,
so here we have to check if it's `nil` or not.  if it isn't, we need to
set it.  We do this here because we need to *cast* params.Timezone into
a pointer to a string, because Go is weird about types.

We then call a (thus far undefined) function called `getTimeOfDay`.

Let's deal with the error case:

<<gettime implementation>>=
	if err != nil {
		return operations.NewTimeGetNotFound().WithPayload(
			&models.ErrorResponse {
				int32(operations.TimeGetNotFoundCode),
				swag.String(fmt.Sprintf("%s", err)),
			})
	}
@

That's a lot of references.  We have a model, an operation, and what's
that "swag" thing?  In order to satisfy Swagger's strictness, we use
only what Swagger offers: for our 404 case, we didn't find the timezone
requested, so we're returning the ErrorResponse model populated with a
numeric code and a string, extracted via `fmt`, from the err returned
from our time function.  The 404 case for get is called, yes,
`NewClockGetNotFound`, and then `WithPayload()` decorates the body of
the response with content.

The good path is similar:

<<gettime implementation>>=
	return operations.NewClockGetOK().WithPayload(
		&models.Timeofday{
			Timeofday: *thetime,
		})
}
@

Now might be a good time to go look in `models/` and `/restapi/options`,
to see what's available to you.  You'll need to do so anyway, because
unless you go to the
[git repository](https://github.com/elfsternberg/go-swagger-tutorial)