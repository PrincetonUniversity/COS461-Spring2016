

### Assignment 1: HTTP Server and Proxy

#### Due 6:00 p.m. Thursday, March 3, 2016

#### Getting Started
* On your host machine (laptop), go to the course directory. 
```bash 
$ cd COS461-Spring2016
```

* Now, pull the latest update from Github.
```bash
$ git pull
```

* Reprovision your VM as follows: 
```bash
$ vagrant reload --provision
```
* You will find the programming assignment in the vm under the following directory:
```bash
vagrant@cos461:~$ ls /vagrant/notebook/assignment1
README.md dev
```

* You will find the following files in the `dev` directory:
```bash
vagrant@cos461:~$ ls /vagrant/notebook/assignment0/dev
Makefile          www               proxy.go          proxy_parse.h
README            http_server.c     proxy_parse.c     test_scripts
```

#### Overview

In this assignment, you will implement a web server that receives requests and returns files locally stored on the server. You will also implement a web proxy that passes requests and data between multiple web clients and web servers. Both server and proxy should support **concurrent** connections. This assignment will give you a chance to get to know one of the most popular application protocols on the Internet -- the Hypertext Transfer Protocol (HTTP) -- and give you an introduction to the Berkeley sockets API. When you're done with the assignment, you should be able to demonstrate the use of your web server with a standard web browser and configure Firefox to use your personal proxy server as a web proxy.

<a name="Introduction:_The_Hypertext_Transfer_Protocol"></a>

#### Introduction: The Hypertext Transfer Protocol

The Hypertext Transfer Protocol (HTTP) is the protocol used for communication on the web: it defines how your web browser requests resources from a web server and how the server responds. For simplicity, in this assignment, we will be dealing only with version 1.0 of the HTTP protocol, defined in detail in [RFC 1945](http://www.ietf.org/rfc/rfc1945.txt "http://www.ietf.org/rfc/rfc1945.txt"). You may refer to that RFC while completing this assignment, but our instructions should be self-contained.

HTTP communications happen in the form of transactions; a transaction consists of a client sending a request to a server and then reading the response. Request and response messages share a common basic format:

*   An initial line (a request or response line, as defined below)
*   Zero or more header lines
*   A blank line (CRLF)
*   An optional message body.

The initial line and header lines are each followed by a "carriage-return line-feed" (\r\n) signifying the end-of-line.

For most common HTTP transactions, the protocol boils down to a relatively simple series of steps (important sections of [RFC 1945](http://www.ietf.org/rfc/rfc1945.txt "http://www.ietf.org/rfc/rfc1945.txt") are in parenthesis):

1.  A client creates a connection to the server.
2.  The client issues a request by sending a line of text to the server. This **request line** consists of an HTTP _method_ (most often GET, but POST, PUT, and others are possible), a _request URI_ (like a URL), and the protocol version that the client wants to use (HTTP/1.0). The request line is followed by one or more header lines. The message body of the initial request is typically empty. (5.1-5.2, 8.1-8.3, 10, D.1)
3.  The server sends a response message, with its initial line consisting of a **status line**, indicating if the request was successful. The status line consists of the HTTP version (HTTP/1.0), a _response status code_ (a numerical value that indicates whether or not the request was completed successfully), and a _reason phrase_, an English-language message providing description of the status code. Just as with the request message, there can be as many or as few header fields in the response as the server wants to return. Following the CRLF field separator, the message body contains the data requested by the client in the event of a successful request. (6.1-6.2, 9.1-9.5, 10)
4.  Once the server has returned the response to the client, it closes the connection.

It's fairly easy to see this process in action without using a web browser. From a Unix prompt, type:

`telnet www.google.com 80`

This opens a TCP connection to the server at www.google.com listening on port 80 (the default HTTP port). You should see something like this:

```
Trying 74.125.141.103...
Connected to www.google.com.
Escape character is '^]'.
```

type the following:

`GET http://www.google.com/ HTTP/1.0`

and hit enter twice. You should see something like the following:

```
HTTP/1.0 200 OK
Date: Mon, 08 Feb 2016 22:03:31 GMT
(More HTTP headers...)
Content-Type: text/html; charset=ISO-8859-1

<!doctype html><html itemscope="" ...
(More HTML follows)
```

There may be some additional pieces of header information as well, such as setting cookies and/or instructions to the browser or proxy on caching behavior. What you are seeing is exactly what your web browser sees when it goes to the Google home page: the HTTP status line, the header fields, and finally the HTTP message body consisting of the HTML that your browser interprets to create a web page.

<a name="HTTP_Proxies"></a>

##### HTTP Proxies

Ordinarily, HTTP is a client-server protocol. The client (usually your web browser) communicates directly with the server (the web server software). However, in some circumstances it may be useful to introduce an intermediate entity called a proxy. Conceptually, the proxy sits between the client and the server. In the simplest case, instead of sending requests directly to the server, the client sends all of its requests to the proxy. The proxy then opens a connection to the server, and passes on the client's request. The proxy receives the reply from the server, and then sends that reply back to the client. Notice that the proxy is essentially acting like both an HTTP client (to the remote server) and an HTTP server (to the initial client).

Why use a proxy? There are a few possible reasons:

*   **Performance:** By saving a copy of the pages that it fetches, a proxy can reduce the need to create connections to remote servers. This can reduce the overall delay involved in retrieving a page, particularly if a server is remote or under heavy load.
*   **Content Filtering and Transformation:** While in the simplest case the proxy merely fetches a resource without inspecting it, there is nothing that says that a proxy is limited to blindly fetching and serving files. The proxy can inspect the requested URL and selectively block access to certain domains, reformat web pages (for instances, by stripping out images to make a page easier to display on a handheld or other limited-resource client), or perform other transformations and filtering.
*   **Privacy:** Normally, web servers log all incoming requests for resources. This information typically includes at least the IP address of the client, the browser or other client program that they are using (called the User-Agent), the date and time, and the requested file. If a client does not wish to have this personally identifiable information recorded, routing HTTP requests through a proxy is one solution. All requests coming from clients using the same proxy appear to come from the IP address and User-Agent of the proxy itself, rather than the individual clients. If a number of clients use the same proxy (say, an entire business or university), it becomes much harder to link a particular HTTP transaction to a single computer or individual.

**References:**

*   [RFC 1945](http://www.w3.org/Protocols/rfc1945/rfc1945 "http://www.w3.org/Protocols/rfc1945/rfc1945") The Hypertext Transfer Protocol, version 1.0

<a name="HTTP_Server"></a>

#### Part 1: HTTP Server

<a name="The_Basics_P1"></a>

##### The Basics

Your task is to build a web sever capable of accepting HTTP requests and returning response data from locally stored files to a client. The server MUST handle **concurrent** requests by forking a process for each new client request using the `fork()` system call. You will only be responsible for implementing the GET method. All other request methods received by the server should elicit a "Not Implemented" (501) error (see [RFC 1945](http://www.w3.org/Protocols/rfc1945/rfc1945 "http://www.w3.org/Protocols/rfc1945/rfc1945") section 9.5 - Server Error).

Your web server can be completed in either C or C++. It should compile and run (using g++) without errors or warnings on the course VM, producing a binary called `http_server` that takes as its first argument a port to listen from. Don't use a hard-coded port number.

You shouldn't assume that your server will be running on a particular IP address, or that clients will be coming from a pre-determined IP.

<a name="Listening_P1"></a>

##### Listening

When your web server starts, the first thing that it will need to do is establish a socket connection that it can use to listen for incoming connections. Your server should listen on the port specified from the command line and wait for incoming client connections. Each new client request is accepted, and a new process is spawned using `fork()` to handle the request. To avoid overwhelming your server, you should not create more than a reasonable number of child processes (for this assignment, use at most 10), in which case your server should wait until one of its ongoing child processes exits before forking a new one to handle the new request.

Once a client has connected, the server should read data from the client and then check for a properly-formatted HTTP request -- but don't worry, we have provided you with libraries that parse the HTTP request lines and headers. Specifically, you will use our libraries to ensure that the server receives a request that contains a valid request line:
```
<METHOD> <URL> <HTTP VERSION>
```

All other headers just need to be properly formatted:
```
<HEADER NAME>: <HEADER VALUE>
```
Your server should accept requests for files ending in `html, txt, gif, jpeg, jpg, or css` and transmit them to the client with a `Content-Type` of `text/html, text/plain, image/gif, image/jpeg, image/jpeg, or text/css`, respectively. If the client requests a file with any other extension, the web server must respond with a well-formed 400 "Bad Request" code. An invalid request from the client should be answered with an appropriate error code, i.e. "Bad Request" (400) or "Not Implemented" (501) for valid HTTP methods other than GET. If the requested file does not exist, your server should return a well-formed 404 "Not Found" code. Similarly, if headers are not properly formatted for parsing or any other error condition not listed before, your server should also generate a type-400 message.


<a name="Parsing_Library_P1"></a>

##### Parsing Library

We have provided a parsing library to do string parsing on the header of the request. This library is in <tt>proxy_parse.[c|h]</tt> in the skeleton code. The library can parse the request into a structure called <tt>ParsedRequest</tt> which has fields for things like the host name (domain name) and the port. It also parses the custom headers into a set of ParsedHeader structs which each contain a key for the header field name and value corresponding to the value to which the header is set. You can search for headers by the key or header field name and modify them. The library can also recompile the headers into a string given the information in the structs.

**More details as well as an example of how to use the library is included in the header file, <tt>proxy_parse.h</tt>.** This library can also be used to verify that the headers are in the correct format since the parsing functions return error codes if this is not the case.

Your server should answer all valid requests with the following:

<pre>
HTTP/1.0 200 OK
Connection: close
Content-Length: [length of file]
Content-Type: [correct content type] 

[Content of the file]
</pre>

For invalid requests, your server should answer with the following:

<pre>
HTTP/1.0 400 Bad Request [or other error code]
Connection: close
Content-Length: [length of message]
Content-Type: text/html

[a well-formed html message describing the error]
</pre>

We provide the strings for the error messages in the start code.

<a name="Parsing_the_URL_P1"></a>

##### Parsing the URL

Once the web server receives a valid HTTP request, it will need to parse the requested URL. The server needs at least one piece of information: the path of the file. See the `URL (7)` manual page for more info. If the path is just "/", your server should return the content of the file index.html.

<a name="Testing_Your_Server"></a>

##### Testing Your Server

We provide html and image files in the dirctory www/images of the start code directory (dev). Go to the www directory and run your server with the following command (assuming that http_server is in the dev folder of the start code):

`../http_server <port>`, where `port` is the port number that the server should listen on. As a basic test of functionality, try requesting a page using telnet:

```
telnet localhost <port>
Trying 127.0.0.1...
Connected to localhost.localdomain (127.0.0.1).
Escape character is '^]'.
GET / HTTP/1.0
```

If your server is working correctly, the headers and HTML of the provided index.html file should be displayed on your terminal screen. Additionally, try requesting a page using telnet concurrently from two different shells. You can also test your server using a browser. The browser should be able to render the page given in the start code.You can install Firefox on the course VM using the command `sudo apt-get install Firefox`.

<a name="Proxy_Server"></a>

#### Part 2: HTTP Proxy

<a name="The_Basics_P2"></a>

##### The Basics

Your task is to build a web proxy capable of accepting HTTP requests, forwarding requests to remote (origin) servers, and returning response data to a client. The proxy will be implemented in Go and MUST handle **concurrent** requests by creating a Go routine for each new client request. You will only be responsible for implementing the GET method. All other request methods received by the proxy should elicit a "Not Implemented" (501) error (see [RFC 1945](http://www.w3.org/Protocols/rfc1945/rfc1945 "http://www.w3.org/Protocols/rfc1945/rfc1945") section 9.5 - Server Error).

Your proxy implementation should compile and run (using go build) without errors or warnings on the course VM, producing a binary called `proxy` that takes as its first argument a port to listen from. Don't use a hard-coded port number.

You shouldn't assume that your proxy will be running on a particular IP address, or that clients will be coming from a pre-determined IP.

<a name="Listening_P2"></a>

##### Listening

When your proxy starts, the first thing that it will need to do is establish a socket connection that it can use to listen for incoming connections. Your proxy should listen on the port specified from the command line and wait for incoming client connections. Each new client request is accepted, and a new Go routine is spawned to handle the request.

Once a client has connected, the proxy should read data from the client and then check for a properly-formatted HTTP request. Go provides packages to parse the HTTP request lines and headers. Specifically, you will use the package `net/http` to ensure that the proxy receives a request that contains a valid request line (see the sever description above for details about HTTP lines and headers). You should NOT use any Proxy method from the http package (http.*Proxy*).

<a name="Parsing_Library_P2"></a>

##### Parsing and Networking Libraries in Go

For this assignment, you can use the packages `net` and `net/http` for implementing the networking and HTTP functionalities of your proxy. Some of the basic methods of these packages will be covered during the precepts, but it is your responsability to study these packages carefully as part of this assignment.

<a name="Parsing_the_URL_P2"></a>

##### Parsing the URL

Once the proxy receives a valid HTTP request, it will need to parse the requested URL. The proxy needs at least three pieces of information: the requested host, port, and path. See the `URL (7)` manual page for more info. You will need to parse the absolute URL specified in the given request line. **You can use the parsing library to help you.** If the hostname indicated in the absolute URL does not have a port specified, you should use the default HTTP port 80.<a name="Getting_Data_from_the_Remote_Server"></a>

##### Getting Data from the Remote Server

Once the proxy has parsed the URL, it can make a connection to the requested host (using the appropriate remote port, or the default of 80 if none is specified) and send the HTTP request for the appropriate resource. The proxy should always send the request in the relative URL + Host header format regardless of how the request was received from the client:  

Accept from client:

<pre>GET http://www.princeton.edu/ HTTP/1.0
</pre>

Send to remote server:

<pre>GET / HTTP/1.0
Host: www.princeton.edu
Connection: close
(Additional client specified headers, if any...)
</pre>

Note that we always send HTTP/1.0 flags and a <tt>Connection: close</tt> header to the server, so that it will close the connection after its response is fully transmitted, as opposed to keeping open a persistent connection. So while you should pass the client headers you receive on to the server, you should make sure you replace any <tt>Connection</tt> header received from the client with one specifying <tt>close</tt>, as shown. To add new headers or modify existing ones, use the HTTP Request Parsing Library provided by Go.<a name="Returning_Data_to_the_Client"></a>

##### Returning Data to the Client

After the response from the remote server is received, the proxy should send the response message as-is to the client via the appropriate socket. To be strict, the proxy would be required to ensure a <tt>Connection: close</tt> is present in the server's response to let the client decide if it should close it's end of the connection after receiving the response. However, checking this is not required in this assignment for the following reasons. First, a well-behaving server would respond with a <tt>Connection: close</tt> anyway given that we ensure that we sent the server a close token. Second, we configure Firefox to always send a <tt>Connection: close</tt> by setting keepalive to false. Finally, we wanted to simplify the assignment so you wouldn't have to parse the server response.

The following summarizes how status replies should be sent from the proxy to the client:

1.  For any error your proxy should return the status 500 'Internal Error'. This means for any request method other than GET, your proxy should return the status 500 'Internal Error' rather than 501 'Not Implemented'. Likewise, for any invalid, incorrectly formed headers or requests, your proxy should return the status 500 'Internal Error' rather than 400 'Bad Request' to the client. For any error that your proxy has in processing a request such as failed memory allocation or missing files, your proxy should also return the status 500 'Internal Error'. (This is what is done by default in this case.)
2.  Your proxy should simply forward status replies from the remote server to the client. This means most 1xx, 2xx, 3xx, 4xx, and 5xx status replies should go directly from the remote server to the client through your proxy. Most often this should be the status 200 'OK'. However, it may also be the status 404 'Not Found' from the remote server. (While you are debugging, make sure you are getting valid 404 status replies from the remote server and not the result of poorly forwarded requests from your proxy.)

<a name="Testing_Your_Proxy"></a>

##### Testing Your Proxy

Run your proxy with the following command:

`./proxy <port>`, where `port` is the port number that the proxy should listen on. As a basic test of functionality, try requesting a page using telnet:

```
telnet localhost <port>
Trying 127.0.0.1...
Connected to localhost.localdomain (127.0.0.1).
Escape character is '^]'.
GET http://www.google.com/ HTTP/1.0
```

If your proxy is working correctly, the headers and HTML of the Google homepage should be displayed on your terminal screen. Notice here that we request the absolute URL (`http://www.google.com/`) instead of just the relative URL (`/`). A good sanity check of proxy behavior would be to compare the HTTP response (headers and body) obtained via your proxy with the response from a direct telnet connection to the remote server. Additionally, try requesting a page using telnet concurrently from two different shells.

For a slightly more complex test, you can configure Firefox to use your proxy server as its web proxy as follows:

1.  Go to the 'Edit' menu.
2.  Select 'Preferences'. Select 'Advanced' and then select 'Network'.
3.  Under 'Connection', select 'Settings...'.
4.  Select 'Manual Proxy Configuration'. If you are using localhost, remove the default 'No Proxy for: localhost 127.0.0.1.' Enter the hostname and port where your proxy program is running.
5.  Save your changes by selecting 'OK' in the connection tab and then select 'Close' in the preferences tab.

<a name="Socket_Programming"></a>

#### Socket and Multi-Process Programming

You can find details for the Berkeley sockets library in the Unix `man` pages (most of them are in Section 2) and in the Stevens _Unix Network Programming_ book, particularly chapters 3 and 4\. Other sections you may want to browse include the client-server example system in Chapter 5 (you will need to write the server code for this assignment) and the name and address conversion functions in Chapter 9\. Please refer to the precept slides in order to review the [Socket Programming tutorial](./docs/rec01-sockets.pdf "https://www.dropbox.com/s/qvuvlqblr7htsql/rec01-sockets.ppt?dl=0").

In addition to the Berkeley sockets library, there are some functions you will need to use for creating and managing multiple processes: fork and waitpid. These will be reviewed in the precepts.

**References:**

*   [Guide to Network Programming Using Sockets](http://beej.us/guide/bgnet/ "http://beej.us/guide/bgnet/")
*   [HTTP Made Really Easy- A Practical Guide to Writing Clients and Servers](http://www.jmarshall.com/easy/http/ "http://www.jmarshall.com/easy/http/")
*   [Wikipedia page on fork()](http://en.wikipedia.org/wiki/Fork_(operating_system) "http://en.wikipedia.org/wiki/Fork_(operating_system)")

<a name="Grading"></a>

#### Grading

You should submit your completed proxy and web server by the date posted on the course website to [CS Dropbox](https://dropbox.cs.princeton.edu/COS461_S2016/Assignment1). You will need to submit a tarball file containing the following:

*   All of the source code for your proxy and web server.
*   A Makefile that builds your proxy and web server.
*   A README file describing your code and the design decisions that you made.

Your tarball should be named `ass1.tgz`. The sample Makefile in the skeleton tar file we provide will make this tarball for you with the `make tar` command.

Your assignment will be graded out of twenty points, with the following criteria:

**General Tests**

3.  When running `make` on your assignment, it should compile without errors or warnings on the course VM and produce binaries named `proxy` and `http_server`. You will receive one point if your code compile without errors or warnings.
4.  You will receive one point for a well-written README file.
5.  You will receive two points if your server and proxy pass simple `wget` tests provided in the start code.

**Web Server**

9.  You will receive three points if your web server returns the content of the web directory (html, image, etc. files) provided with the start code and Firefox can render the web page correctly.
10.  You will receive two points if your web server can handle concurrent connections
11.  You will receive one point if your web server return the correct error codes.
12.  You will receive two points if your web server passes on our private test cases.

**Proxy**

16.  We'll first check that your proxy works correctly with a small number of major web pages, using the same scripts that we've given you to test your proxy. If your proxy passes all of these 'public' tests, you will get five of the possible points.
17.  We'll then check a number of additional URLs and transactions that you will not know in advance. If your proxy passes **all** of these tests, you get three additional points. These tests will check the overall robustness of your proxy, and how you handle certain edge cases. This may include sending your proxy incorrectly formed HTTP requests, large transfers, etc.

* * *

</div>
