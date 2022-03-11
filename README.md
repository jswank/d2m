# D2M

Given a directory with a Maven2 component, create an effective manifest.

(godoc link)

# x2j
Convert a Maven POM file to JSON

## Install
```console
$ go install install github.com/jswank/d2m@latest
```

## Usage

```console
$ d2m maven2-component-dir
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
