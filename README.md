### 阿里云物联网套件IOT SDK Go语言版本

* 简单易使用
* 优秀的随机字符串生成算法，良好的性能，即便 `time.Now().UnixNano()` 生成的随机种子，同时启动100个协程也不会重复，详情请查看测试用例


```
go get -u github.com/cheerego/aliyun-iot
```

```
import (
	"github.com/cheerego/aliyun-iot"
)
var c iot.Client = Client{
	AccessKeyId:     "{AccessKeyId}",
	AccessKeySecret: "{AccessKeySecret}",
	Version:         "{2018-01-20}",
	RegionId:        "{cn-shanghai}",
}


func main(){
    request := make(map[string]string)
    request["ProductKey"] = "xxxxx"
    request["Action"] = "RRpc"
    request["RequestBase64Byte"] = base64.StdEncoding.EncodeToString([]byte("1312312"))
    request["DeviceName"] = "MACHINE_100023"
    request["Timeout"] = "5000"
    res := c.Send(request)
    fmt.Println(res)
}
```
