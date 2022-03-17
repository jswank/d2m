# d2m

[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/jswank/d2m)

Given a directory or URL with a Maven2 component, create an effective manifest.

## Install
```console
$ go install github.com/jswank/d2m/d2m@latest
```

## Usage

```console
$ d2m maven2-component-dir
$ d2m https://repo1.maven.org/maven2/com/github/120011676/cipher/0.0.7
{
    "timestamp": "",
    "version": "",
    "coordinates": {
      "group": "com.github.120011676",
      "artifact": "cipher",
      "version": "0.0.7"
    },
    "files": [
      {
        "size": 117461,
        "filename": "cipher-0.0.7-javadoc.jar",
        "mime_type": "application/x-java-archive",
        "hashes": {
          "md5": "5fbeea7da1a003f588ce66d84139304c",
        }
      }
    ]
    }
... 
```
