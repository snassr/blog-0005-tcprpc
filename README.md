# controlled TCP RPC server example

RPC (Remote Procedure Call): Information is placed on a call stack then control flow is transferred to another part of the program.

## Summary
- Go RPC client will only talk to a Go server (using gob serilization).
- Restrictions
    - Function must be public.
    - Two Arguments (pointer to a value for data to be received by client function, and pointer to a value to hold response returned to the client).
    - Have a return value of type error. ```func exFunc(&R, &W) error {}```

## Run
```bash
# start server
cd $GOPATH/src/github.com/snassr/blog-0003-gorpc/server && go run server.go
# run client
cd $GOPATH/src/github.com/snassr/blog-0003-gorpc/client && go run client.go
```