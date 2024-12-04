# MicroStrategy REST API for Go

`mstr-rest-client` is a framework for interacting with MicroStrategy's REST API.

It currently implements Anonymous and Standard authentication, and handles the tokens and cookies
necessary for an interaction of multiple requests.

## Getting started

Here is some sample code demonstrating how to use the module.

```
package main

import (
    "context"

    "github.com/paulbailey/mstr-rest-client/client"
)

func main() {
    ctx := context.Background()
    demoClient := client.NewStandardMstrRestClient("dave", "dave'sPassword", "https://demo.microstrategy.com/MicroStrategyLibrary/api/")
    demoClient.Login(ctx)
    projects, projectsErr := demoClient.Project.GetProjects(ctx)
    if projectsErr != nil {
        panic(projectsErr)
    }
    demoClient.Logout()
}
```