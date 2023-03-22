import { PrometheusDriver, QueryResult } from 'prometheus-query';

const prom = new PrometheusDriver({
    endpoint: 'http://prometheus:9090/',
    /*endpoint: 'http://localhost:9090/',
    requestInterceptor: {
        //onFulfilled: (config: AxiosRequestConfig) => {
        //    return config;
        //},
        onRejected: (error: any) => {
            return Promise.reject(error.message);
        }
    },
    responseInterceptor: {
        //onFulfilled: (res: AxiosResponse) => {
        //    return res;
        //},
        onRejected: (error: any) => {
            return Promise.reject(error.message);
        }
    }*/
    //endpoint: "https://prometheus.demo.do.prometheus.io/",
});
/*const q = 'up{instance="demo.do.prometheus.io:9090",job="node"}';
prom.instantQuery(q)
    .then((res) => {
        const series = res.result;
        series.forEach((serie) => {
            console.log("Serie:", serie.metric.toString());
            console.log("Time:", serie.value.time);
            console.log("Value:", serie.value.value);
        });
    })
    .catch(console.error);*/
const cpu_query = 'avg by (instance) (rate(node_cpu_seconds_total{mode!="idle"}[1m])) * 100';
prom.instantQuery(cpu_query)
    .then((res: QueryResult) => {
        const series = res.result;
        series.forEach((serie) => {
            console.log("Last query time:", serie.value.time);
            console.log(serie.metric.toString(), "cpu_usage", serie.value.value);
        });
    })
    .catch(console.error);

const mem_query = 'node_memory_active_bytes/node_memory_total_bytes * 100';
prom.instantQuery(mem_query)
    .then((res: QueryResult) => {
        const series = res.result;
        series.forEach((serie) => {
            console.log(serie.metric.toString(), "mem_usage", serie.value.value);
        });
    })
    .catch(console.error);

const disk_query = '(node_filesystem_avail_bytes{mountpoint="/",fstype!="rootfs"} * 100) / node_filesystem_size_bytes{mountpoint="/",fstype!="rootfs"}';
prom.instantQuery(disk_query)
    .then((res: QueryResult) => {
        const series = res.result;
        series.forEach((serie) => {
            console.log(serie.metric.toString(), "disk_usage", serie.value.value);
        });
    })
    .catch(console.error);