## ⚠️ Deprecated!

**This repository is deprecated.** <br />
A more up-to-date version of examples is available here: https://github.com/khulnasoft/kengine-plugin-examples

**Available examples:**
- [datasource-basic](https://github.com/khulnasoft/kengine-plugin-examples/blob/master/examples/datasource-basic) - demonstrates how to build a basic data source plugin with a backend
- [datasource-http](https://github.com/khulnasoft/kengine-plugin-examples/blob/master/examples/datasource-http) - demonstrates how to query data from HTTP-based APIs.
- [datasource-streaming-websocket](https://github.com/khulnasoft/kengine-plugin-examples/blob/master/examples/datasource-streaming-websocket) - demonstrates how to create an event-based data source plugin using RxJS and web sockets.

---

# Kengine Data Source Backend Plugin Template

[![Build](https://github.com/khulnasoft/kengine-starter-datasource-backend/workflows/CI/badge.svg)](https://github.com/khulnasoft/kengine-datasource-backend/actions?query=workflow%3A%22CI%22)

This template is a starting point for building Kengine Data Source Backend Plugins

## What is Kengine Data Source Backend Plugin?

Kengine supports a wide range of data sources, including Prometheus, MySQL, and even Datadog. There’s a good chance you can already visualize metrics from the systems you have set up. In some cases, though, you already have an in-house metrics solution that you’d like to add to your Kengine dashboards. Kengine Data Source Plugins enables integrating such solutions with Kengine.

For more information about backend plugins, refer to the documentation on [Backend plugins](https://kengine.com/docs/kengine/latest/developers/plugins/backend/).

## Getting started

A data source backend plugin consists of both frontend and backend components.

### Frontend

1. Install dependencies

   ```bash
   yarn install
   ```

2. Build plugin in development mode or run in watch mode

   ```bash
   yarn dev
   ```

   or

   ```bash
   yarn watch
   ```

3. Build plugin in production mode

   ```bash
   yarn build
   ```

### Backend

1. Update [Kengine plugin SDK for Go](https://kengine.com/docs/kengine/latest/developers/plugins/backend/kengine-plugin-sdk-for-go/) dependency to the latest minor version:

   ```bash
   go get -u github.com/khulnasoft/kengine-plugin-sdk-go
   go mod tidy
   ```

2. Build backend plugin binaries for Linux, Windows and Darwin:

   ```bash
   mage -v
   ```

3. List all available Mage targets for additional commands:

   ```bash
   mage -l
   ```

## Learn more

- [Build a data source backend plugin tutorial](https://kengine.com/tutorials/build-a-data-source-backend-plugin)
- [Kengine documentation](https://kengine.com/docs/)
- [Kengine Tutorials](https://kengine.com/tutorials/) - Kengine Tutorials are step-by-step guides that help you make the most of Kengine
- [Kengine UI Library](https://developers.kengine.com/ui) - UI components to help you build interfaces using Kengine Design System
- [Kengine plugin SDK for Go](https://kengine.com/docs/kengine/latest/developers/plugins/backend/kengine-plugin-sdk-for-go/)
