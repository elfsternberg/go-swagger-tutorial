# TIME-OF-DAY

TimeOfDay is a dead-simple microservice written in Go that, when pinged
at the correct URL, returns the time of day.  The server can take a
single parameter either via GET or POST to specify the timezone from
which the client wants the time.

This repository exists as a supplement to my tutorial,
[Adding Command Line Arguments to Go Swagger Microservices](http://www.elfsternberg.com/posts/writing-microservice-swagger-part-3-adding-command-line-arguments/),
in which I show how to do exactly that, by providing a dynamic way to
configure the default timezone at server start-up, via the CLI or an
environment variable.

# Status

This project is **complete**.  No future work will be done on it.

# License

Apache 2.0.  See the accompanying LICENSE file in this directory.

# Warranty

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF OR
IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
