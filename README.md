### 阿里云物联网套件IOT SDK Go语言版本

```
go get -u github.com/cheerego/aliyun-iot
```

```
import (
	iot "github.com/cheerego/aliyun-iot"
)
c := iot.Client{
	AccessKeyId:     "{AccessKeyId}",
	AccessKeySecret: "{AccessKeySecret}",
	Version:         "{2018-01-20}",
	RegionId:        "{cn-shanghai}",
}

request := make(map[string]string)
request["ProductKey"] = "xxxxx"
request["Action"] = "RRpc"
request["RequestBase64Byte"] = base64.StdEncoding.EncodeToString([]byte("1312312"))
request["DeviceName"] = "MACHINE_100023"
request["Timeout"] = "5000"
res := c.Send(request)
fmt.Println(res)
```