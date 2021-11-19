## About

This is simple Mantil project template which will demonstrate few Mantil
concepts. It has single Lambda function which holds list of programming excuses
(strings). API has methods for:
* getting [number of items](https://github.com/mantil-io/template-excuses/blob/601410bb2c25d1ea9c825c026087ffde5edcae1f/api/excuses/excuses.go#L45) in the list
* [clearing](https://github.com/mantil-io/template-excuses/blob/601410bb2c25d1ea9c825c026087ffde5edcae1f/api/excuses/excuses.go#L50) list
* [loading](https://github.com/mantil-io/template-excuses/blob/601410bb2c25d1ea9c825c026087ffde5edcae1f/api/excuses/excuses.go#L67) into list from some URL   
* [getting](https://github.com/mantil-io/template-excuses/blob/601410bb2c25d1ea9c825c026087ffde5edcae1f/api/excuses/excuses.go#L56) random item from the list. 


First concept to show is use of environment variables. In this case we will use
project wide environment variable. They can be also set at individual stage
level which is way to configure same Lambda function to work differently in
different stages.   
In our case we will use environment variable to set
[preload_url](https://github.com/mantil-io/template-excuses/blob/601410bb2c25d1ea9c825c026087ffde5edcae1f/config/environment.yml#L36)
which will be used during Lambda function cold start to load initial list of
excuses. If _preload_url_ is not set application will start with empty list of
excuses.

Second concept is integration between UI and API. Project has simple [web page](https://github.com/mantil-io/template-excuses/blob/master/public/index.html) which will show random excuse, and on each click call API to get new random excuse.  

## Create new project from template

`mantil new` command has flag `--from` for creating a new project from existing
template.

We will use this set of commands to create project from template, create new
stage and deploy project to the stage.

```
mantil new my-excuses --from https://github.com/mantil-io/template-excuses
cd my-excses
mantil stage new development
```

After that we can load project web page at URL:
```
open $(mantil env --url)
```

Web page should look like:

![web page](/excuses.png)

Click on the excuse text to get new random text.

## Invoking Methods

Let's call some API methods:

```
mantil invoke excuses/count
mantil invoke excuses/random
```

We can use equivalent curl methods:
```
curl -X POST $(mantil env -u)/excuses/count
curl -X POST $(mantil env -u)/excuses/random
```

Explore `mantil logs` command. It will show logs from the Lambda function
execution. First form will show all logs. Second only REPORT lines.

```
mantil logs excuses
mantil logs excuses --filter-pattern "REPORT"
```

## Loading list of excuses

To clear current list execute:
```
mantil invoke excuses/clear
```
and check:
```
➜ mantil invoke excuses/random
Error: 500 Internal Server Error
X-Api-Error: no excuses
```

To load new list of excuses:
```
mantil invoke excuses/load --data '{"URL": "https://gist.githubusercontent.com/orf/db8eb0aaddeea92dfcab/raw/5e9a8958fce65b1fe8f9bbaadeb87c207e5da848/gistfile1.txt"}'
```

Some URLs for loading excuses lists:

https://gist.githubusercontent.com/ianic/f3335ba0b7ec63cbb821f8a7b735d86e/raw/066e44b04682295781164c538774db645dfe4cc6/excuses.txt

https://gist.githubusercontent.com/fortytw2/78f5f9ef915cb43a3be4/raw/286da386ad35785b2ed9f158e665c8129536e0ce/excuses.txt

https://gist.githubusercontent.com/orf/db8eb0aaddeea92dfcab/raw/5e9a8958fce65b1fe8f9bbaadeb87c207e5da848/gistfile1.txt

## Web interface

_index.html_ page from project _/public_ folder is availabe at root URL.
You can get root URL by:

```
mantil env --url
```

or open in the browser with this terminal command:
```
open $(mantil env --url)
```


## Random excuse in terminal

If you have [jq](https://github.com/stedolan/jq) (JSON processor), this
is usefull one liner to be ready when some manager comes into room:

```
watch -t -n 5 'curl -s -X POST $(mantil env -u)/excuses/random | jq -r .Excuse'
```


## Environment variable

Try to remove _preload_url_ from config/environment.yml. After deploy function
will always start with empty list.

## Test

Run tests with:
```
mantil test
```

Explore test/excuses_test.go file to get the feeling how to use integration
tests in Mantil.

## Cleanup

Remove created stage with:
```
mantil stage destroy development
```

After that all resource created in the AWS account are remove. You can delete
this test project folder.

```
cd ..
rm -rf my-excuses
```
