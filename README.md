## Phrase hasher
#### Usage:
- to run with docker: ```make```
- to run without docker: ```make run``` (port is hardcoded: 80)
- send POST-request to ```/get-phrase-hash``` with following ```application/json``` body: 
```
[{
    "phrase": "hello"
},
{
    "phrase": "world"  
}]
```

#### [DockerHub](https://hub.docker.com/repository/docker/ilyushagod/hasher)