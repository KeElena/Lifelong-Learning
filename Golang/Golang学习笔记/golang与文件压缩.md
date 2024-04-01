# zip

## 无密码压缩

**一、创建zip的Writer**

* 使用`zip.NewWriter()`创建zip的writer对象

```go
zipFile, err := os.Create("archive.zip")
if err != nil {
    panic(err)
}
defer zipFile.Close()

zipWriter := zip.NewWriter(zipFile)
```

**二、循环写入文件**

* 使用`zipWriter.Create()`在zip文档里创建新的文件
* 使用`io.Copy()`写入zip文档

```go
func WriteToZip(zipWriter *zip.Writer, address string) {
	f, err := os.Open(address)
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	w, err := zipWriter.Create(f.Name())
	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(w, f); err != nil {
		panic(err)
	}
}
```

**三、使用示例**

```go
func main() {
	zipFile, err := os.Create("archive.zip")
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)

	file := []string{"./example/1.txt", "./example/2.txt"}

	for _, addr := range file {
		WriteToZip(zipWriter, addr)
	}

	zipWriter.Close()
}

func WriteToZip(zipWriter *zip.Writer, address string) {
	f, err := os.Open(address)
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	w, err := zipWriter.Create(f.Name())
	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(w, f); err != nil {
		panic(err)
	}
}
```

## 无密码解压

**一、创建zip的Reader**

* 使用`zip.OpenReader()`创建Reader

```go
var err error
var zipFile *zip.ReadCloser
zipFile, err = zip.OpenReader("archive.zip")
if err != nil {
    panic(err)
}
defer zipFile.Close()
```

**二、循环输出文件**

* 需要判断是否为目录，使用`os.OpenFile()`前需要创建目录地址

```go
for _, f := range zipFile.File {
    //判断是否为目录
    if f.FileInfo().IsDir() {
        err = os.MkdirAll(f.Name, f.Mode())
        if err != nil {
            panic(err)
        }
    }
	//创建文件所在的目录
    if err = os.MkdirAll(filepath.Dir(f.Name), os.ModePerm); err != nil {
        panic(err)
    }
    var output *os.File
    var input io.ReadCloser
    //创建文件，os.O_TRUNC表示已有不再写入
    output, err = os.OpenFile(f.Name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, f.Mode())
    if err != nil {
        panic(err)
    }

    input, err = f.Open()
    if err != nil {
        panic(err)
    }

    if _, err = io.Copy(output, input); err != nil {
        panic(err)
    }
    _ = output.Close()
    _ = input.Close()
}
```

**三、解压示例**

```go
func main() {
	var err error
	var zipFile *zip.ReadCloser
	zipFile, err = zip.OpenReader("archive.zip")
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	for _, f := range zipFile.File {
		if f.FileInfo().IsDir() {
			err = os.MkdirAll(f.Name, f.Mode())
			if err != nil {
				panic(err)
			}
		}

		if err = os.MkdirAll(filepath.Dir(f.Name), os.ModePerm); err != nil {
			panic(err)
		}
		var output *os.File
		var input io.ReadCloser
		output, err = os.OpenFile(f.Name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}
	
		input, err = f.Open()
		if err != nil {
			panic(err)
		}

		if _, err = io.Copy(output, input); err != nil {
			panic(err)
		}
		_ = output.Close()
		_ = input.Close()
	}
}
```











