# cf-foundation-report-plugin (WIP)

``` bash
>_ cf foundation-report | jq
{
  "apps": [
    {
      "name": "my-app",
      "instances": 2
    },
    {
      "name": "another-app",
      "instances": 3
    }
  ],
  "orgs": [
    {
      "name": "mamachanko"
    },
    {
      "name": "cloud-friends"
    }
  ],
  "spaces": [
    {
      "name": "mamachanko-personal"
    }
  ]
}
```
## Installation
```bash
cf install-plugin https://github.com/mamachanko/cf-foundation-report-plugin/releases/download/0.0.1alpha/main
```
