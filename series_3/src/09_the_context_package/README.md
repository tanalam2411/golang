### Context Package

Three ways to create context:
- Background
- WithCancel
- WithTimeout
- WithDeadline

Http Context example:

Server:
```bash
server$ go run main.go 


2019/12/10 10:44:11 Handler started...
2019/12/10 10:44:16 Handler ended...
2019/12/10 10:44:48 Handler started...
2019/12/10 10:44:48 context canceled
2019/12/10 10:44:48 Handler ended...

```

Client:
```bash
client$ go run main.go 
hello
client$ go run main.go 
^Csignal: interrupt

```

---
- What ever is there in the context should be request specific and shouldn't impact the flow fo you're program.

- Values do not propagate from client to server unlike cancellation.

---
[ref](https://www.youtube.com/watch?v=LSzR0VEraWw&t=405s)

[git](https://github.com/campoy/justforfunc/tree/master/09-context) 