## simple-fileserver


### Usage

#### upload one file
```
curl -X POST http://localhost:8001/upload \
  -F "files=@/home/test/demo.txt" \
  -H "Content-Type: multipart/form-data"
```
#### upload multiple file
```
curl -X POST http://localhost:8001/upload \
  -F "files=@/home/test/demo.txt" \
  -F "files=@/home/test/demo2.txt" \
  -H "Content-Type: multipart/form-data"
```