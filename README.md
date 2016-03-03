# project_cars_udp_recorder
Records data submitted by Project CARS over UDP

The Client component reads raw Project CARS data sent over UDP, and submits the data (as JSON) to the Server component

h3. running tests:
```go test -v ././... -run ^TestParse$```
