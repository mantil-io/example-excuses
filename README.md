## About

This is a simple Mantil project template that will demonstrate a few Mantil
concepts. It has a single Lambda function that holds a list of programming excuses
(strings). API has methods for:
* getting [number of items](https://github.com/mantil-io/template-excuses/blob/601410bb2c25d1ea9c825c026087ffde5edcae1f/api/excuses/excuses.go#L45) in the list
* [clearing](https://github.com/mantil-io/template-excuses/blob/601410bb2c25d1ea9c825c026087ffde5edcae1f/api/excuses/excuses.go#L50) list
* [loading](https://github.com/mantil-io/template-excuses/blob/601410bb2c25d1ea9c825c026087ffde5edcae1f/api/excuses/excuses.go#L67) into list from some URL   
* [getting](https://github.com/mantil-io/template-excuses/blob/601410bb2c25d1ea9c825c026087ffde5edcae1f/api/excuses/excuses.go#L56) random item from the list. 


The first concept to show is the use of environment variables. In this case we will use
a project-wide environment variable. They can be also set at the individual stage
level which is a way to configure the same Lambda function to work differently in
different stages.   
In our case we will use an environment variable to set
[preload_url](https://github.com/mantil-io/template-excuses/blob/601410bb2c25d1ea9c825c026087ffde5edcae1f/config/environment.yml#L36)
which will be used during Lambda function cold start to load the initial list of
excuses. If _preload_url_ is not set application will start with an empty list of
excuses.

The second concept is integration between UI and API. The project has a simple [web page](https://github.com/mantil-io/template-excuses/blob/master/public/index.html) which will show random excuses, and on each click call API to get a new random excuse.  

## Prerequisites

This example is created with Mantil. To download [Mantil CLI](https://github.com/mantil-io/mantil#installation) on Mac or Linux use Homebrew 
```
brew tap mantil-io/mantil
brew install mantil
```
or check [direct download links](https://github.com/mantil-io/mantil#installation).

To deploy this application you will need an [AWS account](https://aws.amazon.com/premiumsupport/knowledge-center/create-and-activate-aws-account/).

## Installation

Note: If this is the first time you are using Mantil you will need to install Mantil Node on your AWS account. For detailed instructions please follow the [one-step setup](https://github.com/mantil-io/mantil/blob/master/docs/getting_started.md#setup)
```
mantil aws install
```

`mantil new` command has a flag `--from` for creating a new project from existing
template.

We will use this set of commands to create a project from a template, create new
stage and deploy a project to the stage.

```
mantil new my-excuses --from https://github.com/mantil-io/template-excuses
cd my-excuses
mantil deploy
```

After that we can load the project web page at URL:
```
open $(mantil env --url)
```

The web page should look like:

![web page](/excuses.png)

Click on the excuse text to get a new random text.

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

Explore the `mantil logs` command. It will show logs from the Lambda function
execution. The first form will show all logs. Second only REPORT lines.

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

_index.html_ page from project _/public_ folder is available at root URL.
You can get the root URL by:

```
mantil env --url
```

or open in the browser with this terminal command:
```
open $(mantil env --url)
```


## Random excuse in terminal

If you have [jq](https://github.com/stedolan/jq) (JSON processor), this
is a useful one-liner to be ready when some manager comes into the room:

```
watch -t -n 5 'curl -s -X POST $(mantil env -u)/excuses/random | jq -r .Excuse'
```


## Environment variable

Try to remove _preload_url_ from config/environment.yml. After deploy function
will always start with an empty list.

## Test

Run tests with:
```
mantil test
```

Explore test/excuses_test.go file to get the feeling of how to use integration
tests in Mantil.

## Cleanup

Remove created stage with:
```
mantil stage destroy development
```

After that all resources created in the AWS account are removed. You can delete
this test project folder.

```
cd ..
rm -rf my-excuses
```
