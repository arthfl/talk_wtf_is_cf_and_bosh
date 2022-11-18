% What is Cloud Foundry and why should i care?
% ...and what the hell is BOSH?

# whoami
- Florian Arthofer
- Principal Software Engineer @ Dynatrace
- Developing cloud integrations and automating everything

# What is Cloud Foundry?
- Open-Source PaaS solution
- Runs nearly everywhere, from vSphere to $whatevercloud
- Multi-cloud is built-in
- Initially created by VMWare as a side project (B29)
- Released in 2011
- open-sourced in 2014 and transferred to Pivotal
- Pivotal got consumed by VMWare again in 2019

# The two dimensions of Cloud Foundry
- dev dimension, where the "end user" lives
    - "cf" CLI tool
    - Buildpacks
    - Docker images
- infrastructure dimension, where the operator lives
    - BOSH
    - lots of YAML

- devs don't need to know or care about the infra dimension
- the dev dimension resides inside the ops dimension. Just like in real life :-)

# dev dimension
![cf platform](pics/powerofplatform.png)

source: https://docs.cloudfoundry.org/concepts/overview.html

# dev dimension
## cf CLI
- How do i deploy an app?
```sh
cf push
```

# dev dimension
## What actually happens during `cf push`?
\includegraphics[width = \textwidth]{pics/cf_push_flow.png}

source: https://docs.cloudfoundry.org/concepts/how-applications-are-staged.html

# dev dimension
## Buildpacks overview
- Just  create the app, don't care about running it
- Cloud Foundry does the rest:
    - create container image
    - store image centrally
    - start app container
    - scale app container
    - monitor app container

# dev dimension
## demo time
- DEMOTIME

# ops dimension
## There's a lot of things running
\includegraphics[width = \textwidth, height = 0.8\textheight]{pics/diego_architecture.png}

# ops dimension
## What the hell is BOSH?
- Tool chain for release engineering, deployment and life cycle management
- Compareable to Ansible/Terraform/Puppet/Chef/etc but does even more
- Launched/released in 2012 as Open Source software
- Abstracts the whole IaaS platform away from you

# ops dimension
## What does BOSH do?
- Deploy infrastructure and services, e.g. Cloud Foundry as whole
- Manage the deployed infrastructure: scale up/down, restart, rebuild, destroy...
- Manage the deployed software
- Keep everything running ("self healing")

# ops dimension
## How does BOSH work?
- Agent/Server based
    - BOSH director
    - bosh-agent
- Stemcells
- Interaction via CLI
- YAML-based Manifests
- CPI
    - AWS, GCP, Azure, OpenStack, RackHD, SoftLayer, VirtualBox, vCloud, vSphere, Warden

# ops dimension
## Example BOSH deployment manifest
- [safe-boshrelease](https://raw.githubusercontent.com/cloudfoundry-community/safe-boshrelease/master/manifests/safe.yml)

# ops dimension
## BOSH demo
- `emtpyvm` BOSH deployment

# Positive/Negative CF
## Positive
- is dead easy for devs
- runs very stable
- does a LOT in terms of container security

## Negative
- Infrastructure footprint isn't tiny
- It's complex (but complex things tend to be)

# Positive/Negative BOSH
## Positive
- abstraction to the max
- actually can delete things
- self-healing
- very mature
- enforced immutable infrastructure

## Negative
- can be quite confusing at first

# Questions?
- Don't be shy

# Links

## Docs:
[https://docs.cloudfoundry.org/](https://docs.cloudfoundry.org/)

[https://ultimateguidetobosh.com/](https://ultimateguidetobosh.com/)

[https://bosh.io/](https://bosh.io/)

## Me:
https://twitter.com/arthfl

https://github.com/arthfl
