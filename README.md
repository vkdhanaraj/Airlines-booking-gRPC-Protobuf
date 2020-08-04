# Airlines booking (gRPC,Protobuf)

## **Objective**

Learn the working of gRPC and Protobuf using Golang.

### **Problems with using REST APIs**

- Relies on human agreement
- Relies on Documentation (requires maintenance/update)
- Need a lot of &quot;formatting, parsing&quot; from and to the agreement (both services)
- Most of the development time is spent on agreeing and formatting not business logic

Using RPC client/library from the creator of the service will ensure correctness when calling the service. If we want to use RPC and REST as the medium, developer B will have to write the client code for developer A to use. If both developers use different languages of choice, this is a major problem for developer B because he needs to write clients in another language that he is not accustomed to. And if different services also need to use service B, developer B will have to spend a lot of time making RPC client in different languages and has to maintain it.

### **gRPC?**

gRPC is a high-performance RPC (Remote Procedure call) framework created by Google. It runs on top of HTTP2. gRPC Automatically generates servers and clients in 10+ languages so you can pick up any language of your choice. It is an efficient way to connect services written in different languages with pluggable support for load balancing, tracing, health checking, and authentication. By default, gRPC uses protocol buffers for serializing structured data. Generally, gRPC is considered as a better alternative to the REST protocol for microservice architecture.

### **Protobuf?**

Protocol buffers are Google&#39;s language-neutral, platform-neutral, extensible mechanism for serializing structured data. gRPC uses protobuf as a language to define data structures and services.

### **Advantages of Protobuf**

- Data is fully typed
- Data is compressed automatically
- Documentation can be embedded in schema (.proto file)
- Data can be read by any language
- Faster
- Code is generated automatically
- Payload size of Protobuf is much smaller that Json (Hence saving network bandwidth)
- Parsing Protobuf data is less CPU intensive because it&#39;s closer to how a machine represents data. ( Json is closer to human readable format)

### **gRPC vs REST**

Unlike REST, which uses JSON (mostly), gRPC uses protocol buffers, which are a better way of encoding data. As JSON is a text-based format, it will be much heavier than compressed data in protobuf format.

Another significant improvement of gRPC over conventional REST is that it uses HTTP 2 as its transfer protocol. HTTP 1.1, which is mainly used by REST, is basically a request-response model. (The REST can also be implemented with HTTP2 ) gRPC makes uses of the bidirectional communication feature of HTTP 2 along with the traditional response-request structure. In HTTP 1.1, when multiple requests come from multiple clients, they are served one by one. This can slow down the system. HTTP 2 allows multiplexing, so multiple requests and responses can be served at the same time.
