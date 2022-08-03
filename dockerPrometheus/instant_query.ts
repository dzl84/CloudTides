import { PrometheusDriver, QueryResult } from 'prometheus-query';

const prom = new PrometheusDriver({
    endpoint: 'https://prometheus.demo.do.prometheus.io/',
});


// function Instant(query: string){
//     // last value
//     prom.instantQuery(query)
//         .then((res: QueryResult) => {
//             const series = res.result;
//             series.forEach((serie) => {
//                 // console.log("[instantQuery] Serie:", serie.metric.toString());
//                 // console.log("[instantQuery] Time:", serie.value.time);
//                 console.log(query, serie.value.value);
//             });
//         })
//         .catch(console.error);
// }

const cpu_query = 'sum by (cpu)(rate(node_cpu_seconds_total{mode!="idle"}[5m]))*100';
prom.instantQuery(cpu_query)
    .then((res: QueryResult) => {
        const series = res.result;
        series.forEach((serie) => {
            console.log("Last query time:", serie.value.time);
            console.log("cpu_usage", serie.value.value);
        });
    })
    .catch(console.error);

const mem_query = 'node_memory_Active_bytes/node_memory_MemTotal_bytes*100';
prom.instantQuery(mem_query)
    .then((res: QueryResult) => {
        const series = res.result;
        series.forEach((serie) => {
            console.log("mem_usage", serie.value.value);
        });
    })
    .catch(console.error);

const disk_query = '(node_filesystem_avail_bytes{mountpoint="/",fstype!="rootfs"} * 100) / node_filesystem_size_bytes{mountpoint="/",fstype!="rootfs"}';
prom.instantQuery(disk_query)
    .then((res: QueryResult) => {
        const series = res.result;
        series.forEach((serie) => {
            console.log("disk_usage", serie.value.value);
        });
    })
    .catch(console.error);