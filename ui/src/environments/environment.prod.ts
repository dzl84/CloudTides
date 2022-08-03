import { base } from './base';

const hostName = window.location.host
const http = window.location.protocol
export const environment = {
  production: false,
  apiPrefix: http+"//"+hostName+'/api/v1',
  apiIp: hostName
};

// export const environment = {
//   //apiPrefix: 'http://106.14.190.68:8081/api/v1',
//   apiPrefix: 'http://127.0.0.1/api/v1',
//   webssh: 'http://127.0.0.1/',
//   apiIp: '127.0.0.1:8045',
//   production: true,
// };
