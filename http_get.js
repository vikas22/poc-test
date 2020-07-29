import http from 'k6/http';
import { check } from 'k6';
import { Rate } from 'k6/metrics';

export let errorRate = new Rate('errors');

export let options = {
  stages:[
    {"duration" : '1m', "target": 500},
    {"duration" : '2m', "target": 1000},
    {"duration" : '2m', "target": 1500},
    {"duration" : '2m', "target": 2000},
    {"duration" : '2m', "target": 2500},
    {"duration" : '1m', "target": 3000},
  ]
}

export default function() {
  var url = 'http://localhost:1234/poc';
  var params = {
    headers: {
    }
  };
  check(http.get(url, params), {
    'status is 200': r => r.status == 200
  }) || errorRate.add(1);
}
