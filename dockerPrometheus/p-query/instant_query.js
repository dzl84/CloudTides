"use strict";
exports.__esModule = true;
var prometheus_query_1 = require("prometheus-query");
var prom = new prometheus_query_1.PrometheusDriver({
    endpoint: 'http://localhost:9090/'
});
var cpu_query = 'avg by (instance) (rate(node_cpu_seconds_total{mode!="idle"}[1m])) * 100';
prom.instantQuery(cpu_query)
    .then(function (res) {
    var series = res.result;
    series.forEach(function (serie) {
        console.log("Last query time:", serie.value.time);
        console.log(serie.metric.toString(), "cpu_usage", serie.value.value);
    });
})["catch"](console.error);
var mem_query = 'node_memory_active_bytes/node_memory_total_bytes * 100';
prom.instantQuery(mem_query)
    .then(function (res) {
    var series = res.result;
    series.forEach(function (serie) {
        console.log(serie.metric.toString(), "mem_usage", serie.value.value);
    });
})["catch"](console.error);
var disk_query = '(node_filesystem_avail_bytes{mountpoint="/",fstype!="rootfs"} * 100) / node_filesystem_size_bytes{mountpoint="/",fstype!="rootfs"}';
prom.instantQuery(disk_query)
    .then(function (res) {
    var series = res.result;
    series.forEach(function (serie) {
        console.log(serie.metric.toString(), "disk_usage", serie.value.value);
    });
})["catch"](console.error);
