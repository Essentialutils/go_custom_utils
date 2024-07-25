[![Go Reference](https://pkg.go.dev/badge/github.com/Essentialutils/go_custom_utils?status.svg)](https://pkg.go.dev/github.com/Essentialutils/go_custom_utils?tab=doc)
[![Release](https://img.shields.io/github/release/Essentialutils/go_custom_utils.svg?style=flat-square)](https://github.com/Essentialutils/go_custom_utils/releases)

## Import

```go
import (
    gcu "github.com/Essentialutils/go_custom_utils"
)
```

## Print Server Details

```go
gcu.InfoPrint(gcu.InfoPrintData{
    ServerRunningPort: ":0000",
    ServerRunningMode: "Debug",
    OpenApiPath:       "/postman",
    Developer:         "Sharafas OM",
    Website:           "https://repad.dev/",
    Version:           "0.0.0",
})
```
output : 

<img src="img/print_info.png" alt="Alt text" height="300">
