<!-- gomarkdoc:embed:start -->

<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# MegaUpload

```go
import "github.com/kilianp07/Prestashop-Backup-to-Cloud/providers/Mega"
```

## Index

- [type Client](<#Client>)
  - [func New\(conf config.MegaConf\) \*Client](<#New>)
  - [func \(c \*Client\) Login\(\) error](<#Client.Login>)
  - [func \(c \*Client\) Upload\(filepath string\) error](<#Client.Upload>)


<a name="Client"></a>
## type [Client](<https://github.com/kilianp07/Prestashop-Backup-to-Google-Drive/blob/main/providers/Mega/mega.go#L10-L14>)



```go
type Client struct {
    // contains filtered or unexported fields
}
```

<a name="New"></a>
### func [New](<https://github.com/kilianp07/Prestashop-Backup-to-Google-Drive/blob/main/providers/Mega/mega.go#L16>)

```go
func New(conf config.MegaConf) *Client
```



<a name="Client.Login"></a>
### func \(\*Client\) [Login](<https://github.com/kilianp07/Prestashop-Backup-to-Google-Drive/blob/main/providers/Mega/mega.go#L25>)

```go
func (c *Client) Login() error
```



<a name="Client.Upload"></a>
### func \(\*Client\) [Upload](<https://github.com/kilianp07/Prestashop-Backup-to-Google-Drive/blob/main/providers/Mega/mega.go#L30>)

```go
func (c *Client) Upload(filepath string) error
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- gomarkdoc:embed:end -->