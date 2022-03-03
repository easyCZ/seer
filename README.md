# Seer - synthetic monitoring for APIs

<img src="./assets/logo.png" align="right" style="border:2px solid black; border-radius: 50%; padding: 2px"
     alt="Synthetic Biology by Laymik from NounProject.com" width="150" height="150">

**Seer** is an open-source implementation of [synthetic monitoring](https://en.wikipedia.org/wiki/Synthetic_monitoring) for HTTP APIs.

**Goals**:

- Self-hosted - deploy on your existing infrastructure (e.g. kubernetes)
- Private networks - run inside VPCs without exposing the service to public internet.
- Low-overhead - utilize unused capacity in your existing compute (kubernetes clusters).

**Motivation**

There are a a number of commercial products offering synthetic monitoring ([DataDog](https://docs.datadoghq.com/synthetics/), [NewRelic](https://docs.newrelic.com/docs/synthetics/synthetic-monitoring/getting-started/get-started-synthetic-monitoring/), [Sematext](https://sematext.com/synthetic-monitoring/)). This tool aims to provide an alternative which can utilize your existing infrastrucuture excess capacity to reduce extra costs (and dependencies) on third party providers.

## Maturity

Seer is in very early stages of maturity and unsuitable for most workloads.

## Developing




The repository is setup to be developed inside GitPod.

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/easyCZ/seer)

## Credits

- Logo by [Synthetic Biology by Laymik from NounProject.com](https://thenounproject.com/icon/synthetic-biology-4116522/)


