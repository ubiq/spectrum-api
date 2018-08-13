# Spectrum-API

API for spectrum-interface written in golang. Serves data from mongodb that has been populated by spectrum-crawler.

### Build

clone to ```$GOPATH/src/github.com/ubiq/spectrum-api```

```
cd $GOPATH/src/github.com/ubiq/spectrum-api
go get
go build
```

### Configure

Configure via config.toml

```
port: port to listen on (e.g :3000)  
server: mongodb server (e.g localhost)
database: mongodb database (e.g spectrumdb)
```

### Run

```
./spectrum-api
```
