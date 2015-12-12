# Dashing

This is a go version of http://dashing.io, but completly rewrited with [golang](http://golang.org) and [vuejs](http://vuejs.org). By using [webpack](https://webpack.github.io) and [vue-loader](https://github.com/vuejs/vue-loader), as a designer or developer you could write a widget in single file and share it easily.

## Demo

- http://dash.4leaf.me/sample
- http://dash.4leaf.me/swiming
- http://dash.4leaf.me/map

## Install

```bash
go get -u github.com/exherb/dashing-cli
```

## Usage

```bash
dashing-cli new dashboard
cd dashboard
npm install
npm run build
go run server.go
```

## Deploy

TODO

## Dashboards

[Examples](https://github.com/exherb/Dashing/tree/master/assets/dashboards)

## Widgets

[Examples](https://github.com/exherb/Dashing/tree/master/assets/src/widgets)

## Getting Data Into Your Widgets

### Jobs

[Examples](https://github.com/exherb/Dashing/tree/master/assets/jobs)

### API

```bash
curl -d '{"token": "{token}", ...}' http://dash.4leaf.me/dashboards/{dashboard_id}/widgets/{widget_id}
```

example

```bash
curl -d '{"token": "okiedokie", "color": "yellow"}' http://dash.4leaf.me/dashboards/*/widgets/sperm
```

if dashboard_id is '*', then the data will send to all dashboards

## Todo

- [x] widget token
- [ ] deploy document
- [ ] save history to database
- [ ] restore historical datas
- [ ] webgl demo
- [ ] video demo
- [ ] kinect demo
- [ ] fasthttp
