## invoke examples

``` shell
mantil invoke excuses/count
mantil invoke excuses/random
```

View logs:
``` shell
mantil logs -n excuses
mantil logs -n excuses -s 24h
mantil logs -n excuses -s 24h -p "REPORT"
```

Load list off excuses with invoke:
``` shell
mantil invoke excuses/load -l -d '{"URL": "https://gist.githubusercontent.com/fortytw2/78f5f9ef915cb43a3be4/raw/286da386ad35785b2ed9f158e665c8129536e0ce/excuses.txt"}'
mantil invoke excuses/load -l -d '{"URL": "https://gist.githubusercontent.com/orf/db8eb0aaddeea92dfcab/raw/5e9a8958fce65b1fe8f9bbaadeb87c207e5da848/gistfile1.txt"}'
mantil invoke excuses/load -l -d '{"URL": "https://gist.githubusercontent.com/ianic/f3335ba0b7ec63cbb821f8a7b735d86e/raw/066e44b04682295781164c538774db645dfe4cc6/excuses.txt"}'
```
Note that invoke shows each log line from api lambda function.

## curl usage 

To get the api public url:

``` shell
mantil env -u
```
which can be then used in curl:

``` shell
curl -X POST $(mantil env -u)/excuses/count
curl -X POST $(mantil env -u)/excuses/random
curl -X POST $(mantil env -u)/excuses/load -d '{"URL": "https://gist.githubusercontent.com/ianic/f3335ba0b7ec63cbb821f8a7b735d86e/raw/066e44b04682295781164c538774db645dfe4cc6/excuses.txt"}'
```

Get random excuse each 5 seconds:
``` shell
watch -t -n 5 'curl -s -X POST $(mantil env -u)/excuses/random | yq -r .Excuse'
```


## excuses lists

URLs with high quality developer excuses lists:

https://gist.githubusercontent.com/ianic/f3335ba0b7ec63cbb821f8a7b735d86e/raw/066e44b04682295781164c538774db645dfe4cc6/excuses.txt

https://gist.githubusercontent.com/fortytw2/78f5f9ef915cb43a3be4/raw/286da386ad35785b2ed9f158e665c8129536e0ce/excuses.txt

https://gist.githubusercontent.com/orf/db8eb0aaddeea92dfcab/raw/5e9a8958fce65b1fe8f9bbaadeb87c207e5da848/gistfile1.txt

