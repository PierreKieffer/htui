

```
                                          _   _        _ 
                                         | |_| |_ _  _(_)                 
                                         | ' \  _| || | |
                                         |_||_\__|\_,_|_|
                                    
```
<div align="center">

**Heroku Terminal User Interface**


<!-- <img src="" /> -->


</div>

---



## Install 

**Note**: Prebuilt binaries for 64-bit operating systems, doesn't require Go.

### 64-bit
#### Linux 
```bash 
curl -sSL https://raw.githubusercontent.com/PierreKieffer/htui/master/install/install_htui64.sh | bash
```
#### Mac OS 
```bash 
curl -sSL https://raw.githubusercontent.com/PierreKieffer/htui/master/install/install_htui64_osx.sh | bash
```

## Run 
```bash
htui
```

## Move around
```
     -----------------------------
     -        Move around        -
     -----------------------------
     go up               ▲  or 'k'
     go down             ▼  or 'j'
     go to the top       'gg'
     go to the bottom    'G'
     select item         'enter'
     Quit htui           'q'
```

## Authentication 
htui uses API token mechanism for authentication to Heroku, with `HEROKU_API_KEY` environment variable. 
If `~/.netrc` file exists (UNIX), `HEROKU_API_KEY` is set automatically. 
If `~/.netrc` doesn't exist, you need to set `HEROKU_API_KEY` manually : 
- Retrieve the API token : 
  - heroku CLI : `heroku auth:token`
  - heroku account setting web page : API Key
- `export HEROKU_API_KEY="api token"` 


## Built With

- [gizak/termui](https://github.com/gizak/termui)


## License 
BSD




