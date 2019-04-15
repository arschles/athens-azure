# Open Census with Azure App Insights 

>This code was copied and adapted from [github.com/bketelsen/censusai](https://github.com/bketelsen/censusai) with permission. Thanks so much to [bketelsen](https://github.com/bketelsen) for this work!

## Getting Started

* Follow the guide [here](https://cda.ms/H8) to create an Application Insights instance
* Go up one directory and export your instrumentation key as an environment variable:
  * `export APPINSIGHTS_INSTRUMENTATIONKEY=xxxxxxx`
  * I use [direnv](http://direnv.net) and MAKE SURE that my `.envrc` file is in `.gitignore`
* Go up one directory and build / run the local forwarder process (which is in [Dockerfile](./Dockerfile)):
  * `make build-forward`
  * `make forward` 

## Shutdown / Cleanup

Be sure to delete the App Insights Resource Group you created in the Azure Portal when you are done experimenting so you won't have a lingering billable service you're not using.

## Links and Documentation

[App Insights - Go + OpenCensus](https://cda.ms/H8)

