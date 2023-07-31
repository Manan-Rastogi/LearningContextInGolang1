# Context in Golang

## Introduction

### What is the context package in Go?
```
The context package in Go provides functions to handle timeouts, cancelation signals, and other request-scoped values across API boundaries and between processes. In terms of Go's concurrent programming model, you can think of context as a way to manage state between multiple goroutines that have a parent-child relationship.
```

### Why use context?
```
Let's say you have a web server that handles incoming requests and each request starts a goroutine. If the client cancels the request or there's a timeout, you'll want to stop any work the server is doing that's related to that request. This could involve stopping long-running goroutines, I/O operations, database queries, and so on. This is where the context package comes in handy.
```

## Types of Context
### context.Background()
```
This returns a non-nil, empty Context. This is typically used when main function begins and other contexts are derived from it.
```

### context.TODO()
```
This also returns a non-nil, empty Context. Developers are encouraged to use context.TODO when it's unclear which Context to use or it is not yet available.
```

### context.WithDeadline()
```
The context.WithDeadline() function returns a copy of the parent context that will be cancelled if the deadline exceeds or if the cancel function is called, whichever happens first. 
```

### context.WithTimeout()
```
The context.WithTimeout() function is a shortcut for context.WithDeadline(parent, time.Now().Add(timeout)). It returns a derived context that gets cancelled if the timeout exceeds or if the cancel function is called, whichever happens first.
```

### context.WithValue()
```
This function is used to attach a key-value pair to a context. The key-value pair can then be retrieved anywhere the context is passed. You should be careful while using context.WithValue() because it can easily lead to tightly coupled and hard-to-maintain code. It's typically used in situations where you want to pass request-scoped data that's used across API boundaries, but not suitable for function parameters.
```

## When to call the cancel() function?
```
The cancel function returned by context.WithCancel, context.WithDeadline, or context.WithTimeout should be called whenever the work associated with the context has finished, regardless of whether it completed successfully or failed.

In a real-world application, this could be when:

*   You've finished processing an HTTP request and have sent the response.
*   You've completed a database query and have processed the results.
*   A user operation is finished.
*   An error occurs that requires the cancellation of associated operations.
```