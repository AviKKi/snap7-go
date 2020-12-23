# snap7-go
snap7 golang binding

#### Updates:
1. Logo interface added, similar to python-snap7



## Usage

```go
host := "109.90.104.232"

plc := snap7.NewLogo()                    // create new instance
err := plc.Connect(host, 0x0200, 0x0300) // connect to default port 102
// err := plc.Connect(host, 0x0200, 0x0300, 1009) // custom port 1009


```