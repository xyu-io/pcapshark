# pcapshark
pcap file parser  implement by gin api

## note
- This project is based on github.com/wader/fq project modifications

### use
#### 1、run app and listen on port 8081
```shell
go run app.go
```
#### 2、post to upload pcap file. exp:
```shell
curl --location 'http://127.0.0.1:8081/upload' \
--header 'User-Agent: Reqable/2.30.3' \
--form 'upload=@"/C:/Users/xuanyu/Desktop/surica.pcap"'
```
