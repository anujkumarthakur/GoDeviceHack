

#####How To use Package to detect Device:
```go
  if deviceType == goDevice.MOBILE {
    fmt.Fprintf(w,"<h1>Mobile</h1>")
  } else if deviceType == goDevice.WEB {
    fmt.Fprintf(w,"<h1>Web</h1>")
  } else if deviceType == goDevice.TABLET {
    fmt.Fprintf(w,"<h1>Tablet</h1>")
  }
  ```
```
For now goDevice is detecting Mobile,Desktop and Tablet.In future we will add support for SmartTV, Watch etc

```
