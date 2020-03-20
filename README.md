# go-fuzz の試用

## 試し方

```
go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build
git clone https://github.com/nus/trygofuzz.git
cd trygofuzz
go-fuzz-build github.com/nus/trygofuzz/lib
go-fuzz -bin=trygofuzz-fuzz.zip
```
