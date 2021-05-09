
# Application Programming Interfaces[APIs]:
    * An interface which has a set of functions that allow programmers/programs to access specific features or data of an application, operating system or other services. 
    * Provide the platform and the medium for applications to talk and understand each other.

# Types of APIs based on:
    * Availability: Private, Partner, Public/External
    * Use cases   : Database APIs, Operating System APIs, Remote APIs, Web APIs

# Web API Architecture styles or Specifications:
    * Remote Procedure Call [RPC]
    * Simple Object Access Protocol [SOAP]
    * REpresentational State Transfer [REST]
    * GraphQL

# HTTP:
Underlying protocol used to transfer data. It would have a request and responses.
```sh
Request:
    METHOD  URL
    HEADER Fields
    BODY
Response:
    ResponseCode ResponseString
    HEADER Fields
    BODY
```
The BODY would consist of data and JSON/XML are used to transfer the data.
    
# REST:
* REST stands for REpresentational State Transfer. It is an Architectural style, or design pattern, for an API.
* It means when a RESTful API is called, the server will transfer to the client a representation of the state of the requested resource.
* The representation of state can be in JSON/XML/HTML format.


# References:
```sh 
https://www.altexsoft.com/blog/engineering/what-is-api-definition-types-specifications-documentation/
https://www.youtube.com/watch?v=lsMQRaeKNDk

https://tutorialedge.net/software-eng/what-is-a-rest-api/

https://cloud.google.com/blog/products/api-management/google-cloud-api-design-tips

https://blog.dreamfactory.com/grpc-vs-rest-how-does-grpc-compare-with-traditional-rest-apis/
https://cloud.google.com/blog/products/api-management/understanding-grpc-openapi-and-rest-and-when-to-use-them
```

## HTTP:
```sh
https://developer.mozilla.org/en-US/docs/Web/HTTP/Status
https://golang.org/pkg/net/http/
```
## Golang
* Can find packages at https://golang.org/pkg/ and https://pkg.go.dev/ (for third party packages)