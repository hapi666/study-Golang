# http学习心得

### 1.GET请求

get请求有好几种方式，

1.直接使用net/http包内函数请求

```go
import(
	"net/http"
)
...
resp,err:=http.Get("http://www.baidu.com")
```

2.利用http.client结构体来请求

```go
import(
	"net/http"
)
...
clt:=http.Client{}
resp,err:=clt.Get("http://www.baidu.com")
```

3.最本质的请求方式

看源代码可以发现，以上两种方式都是用了一下这个最本质的请求方式，即使用http.NewRequest函数。

```go
req,err:=http.NewRequest("GET","http://baidu.com",nil)
//然后http.Client  结构体Do方法
//http.DefaultClient可以换成另外一个http.Client(自己定义)
resp,err:=http.DefaultClient.Do(req)
```

Go的get请求表面上有好几种请求方式，实则只有一种:

######  就是使用http.NewRquest函数获得request实体，然后利用http.Client结构体的Do方法，将request实体传入Do方法中。

### 2.POST请求

和get请求一样，post请求也有好多种方式，但本质还是使用了http.NewRquest函数和http.Client的Do方法。

1.使用net/http包带的post方法

```go
import(
	"net/http"
  	"net/url"
)
...
data:=url.Values{"start":{"0"},"offset":{"xxxx"}}
body:=strings.NewReader(data,Encode())
resp,err:=http.Post("xxxxxx","application/x-www-form-urlencoded",body)
```

或者还可以

```go
import(
	"net/http"
  	"net/url"
)
...
var r http.Request
r.ParseForm()
r.Form.Add("xxx","xxx")
body:=strings.NewReader(r.Form.Encode())
http.Post("xxxx","application/x-www-form-urlencoded",body)
```

2.使用实例化的http.Client的post方法

###### 其实本质上直接使用包函数 和 实例化http.Client是一样的，包函数的底层也仅仅是实例化了一个DefaultClient,然后调用DefaultClient的方法。

下面给出使用实例化的http client的post方法：

```go
import(
	"net/http"
  	"net/url"
)
...
data:=url.Values{"start":{"0"},"offset":{"xxxx"}}
body:=strings.NewReader(data.Encode())
clt:=http.Client{}
resp,err:=clt.Post("xxxxxx","application/x-www-form-urlencoded",body)
```

还有

```go
import(
	"net/http"
  	"net/url"
)
...
var r http.Request
r.ParseForm()
r.Form.Add("xxx","xxx")
body:=strings.NewReader(r.Form.Encode())
clt:=http.Client{}
clt.Post("xxxx","application/x-www-form-urlencoded",body)
```

