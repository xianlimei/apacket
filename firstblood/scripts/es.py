#!/usr/bin/env python
# -*- coding: UTF-8 -*-

resCmd = {
  "took" : 659,
  "timed_out" : "false",
  "_shards" : {
    "total" : 15,
    "successful" : 15,
    "failed" : 0
  },
  "hits" : {
    "total" : 10,
    "max_score" : 1.0,
    "hits" : [ {
      "_index" : "anarchop_apn",
      "_type" : "post",
      "_id" : "86267",
      "_score" : 1.0,
      "fields" : {
        "iswin" : [ "root" ]
      }
    } ]
  }
}

resBanner = {
  "status" : 200,
  "name" : "Blob",
  "cluster_name" : "elasticsearch",
  "version" : {
    "number" : "1.7.3",
    "build_hash" : "05d4530971ef0ea46d0f4fa6ee64dbc8df659682",
    "build_timestamp" : "2015-10-15T09:14:17Z",
    "build_snapshot" : "false",
    "lucene_version" : "4.10.4"
  },
  "tagline" : "You Know, for Search"
}

resNodes = {
    "cluster_name": "elasticsearch",
    "nodes": {
        "gW9D9GRNR4m_0j1-7m_utQ": {
            "name": "Blob",
            "transport_address": "inet[/10.20.149.158:9300]",
            "host": "server1",
            "ip": "10.20.140.1",
            "version": "1.7.3",
            "build": "05d4530",
            "http_address": "inet[/10.20.149.158:9200]",
            "settings": {
                "pidfile": "/var/run/elasticsearch/elasticsearch.pid",
                "path": {
                    "conf": "/etc/elasticsearch",
                    "data": "/var/lib/elasticsearch",
                    "logs": "/var/log/elasticsearch",
                    "home": "/usr/share/elasticsearch"
                },
                "cluster": {
                    "name": "elasticsearch"
                },
                "name": "Blob",
                "client": {
                    "type": "node"
                },
                "foreground": "yes",
                "config.ignore_system_properties": "true",
                "config": "/etc/elasticsearch/elasticsearch.yml",
                "script": {
                    "inline": "on",
                    "indexed": "on"
                }
            },
            "os": {
                "refresh_interval_in_millis": 1000,
                "available_processors": 8,
                "cpu": {
                    "vendor": "Intel",
                    "model": "Xeon",
                    "mhz": 3761,
                    "total_cores": 8,
                    "total_sockets": 8,
                    "cores_per_socket": 32,
                    "cache_size_in_bytes": 10240
                },
                "mem": {
                    "total_in_bytes": 67584700416
                },
                "swap": {
                    "total_in_bytes": 1071636480
                }
            },
            "process": {
                "refresh_interval_in_millis": 1000,
                "id": 1158,
                "max_file_descriptors": 65535,
                "mlockall": "false"
            },
            "jvm": {
                "pid": 1158,
                "version": "1.8.0_72",
                "vm_name": "Java HotSpot(TM) 64-Bit Server VM",
                "vm_version": "25.72-b15",
                "vm_vendor": "Oracle Corporation",
                "start_time_in_millis": 1511653511901,
                "mem": {
                    "heap_init_in_bytes": 268435456,
                    "heap_max_in_bytes": 1037959168,
                    "non_heap_init_in_bytes": 2555904,
                    "non_heap_max_in_bytes": 0,
                    "direct_max_in_bytes": 1037959168
                },
                "gc_collectors": [
                    "ParNew",
                    "ConcurrentMarkSweep"
                ],
                "memory_pools": [
                    "Code Cache",
                    "Metaspace",
                    "Compressed Class Space",
                    "Par Eden Space",
                    "Par Survivor Space",
                    "CMS Old Gen"
                ]
            },
            "thread_pool": {
                "percolate": {
                    "type": "fixed",
                    "min": 8,
                    "max": 8,
                    "queue_size": "1k"
                },
                "fetch_shard_started": {
                    "type": "scaling",
                    "min": 1,
                    "max": 16,
                    "keep_alive": "5m",
                    "queue_size": -1
                },
                "listener": {
                    "type": "fixed",
                    "min": 4,
                    "max": 4,
                    "queue_size": -1
                },
                "index": {
                    "type": "fixed",
                    "min": 8,
                    "max": 8,
                    "queue_size": "200"
                },
                "refresh": {
                    "type": "scaling",
                    "min": 1,
                    "max": 4,
                    "keep_alive": "5m",
                    "queue_size": -1
                },
                "suggest": {
                    "type": "fixed",
                    "min": 8,
                    "max": 8,
                    "queue_size": "1k"
                },
                "generic": {
                    "type": "cached",
                    "keep_alive": "30s",
                    "queue_size": -1
                },
                "warmer": {
                    "type": "scaling",
                    "min": 1,
                    "max": 4,
                    "keep_alive": "5m",
                    "queue_size": -1
                },
                "search": {
                    "type": "fixed",
                    "min": 13,
                    "max": 13,
                    "queue_size": "1k"
                },
                "flush": {
                    "type": "scaling",
                    "min": 1,
                    "max": 4,
                    "keep_alive": "5m",
                    "queue_size": -1
                },
                "optimize": {
                    "type": "fixed",
                    "min": 1,
                    "max": 1,
                    "queue_size": -1
                },
                "fetch_shard_store": {
                    "type": "scaling",
                    "min": 1,
                    "max": 16,
                    "keep_alive": "5m",
                    "queue_size": -1
                },
                "management": {
                    "type": "scaling",
                    "min": 1,
                    "max": 5,
                    "keep_alive": "5m",
                    "queue_size": -1
                },
                "get": {
                    "type": "fixed",
                    "min": 8,
                    "max": 8,
                    "queue_size": "1k"
                },
                "merge": {
                    "type": "scaling",
                    "min": 1,
                    "max": 4,
                    "keep_alive": "5m",
                    "queue_size": -1
                },
                "bulk": {
                    "type": "fixed",
                    "min": 8,
                    "max": 8,
                    "queue_size": "50"
                },
                "snapshot": {
                    "type": "scaling",
                    "min": 1,
                    "max": 4,
                    "keep_alive": "5m",
                    "queue_size": -1
                }
            },
            "network": {
                "refresh_interval_in_millis": 5000,
                "primary_interface": {
                    "address": "10.20.140.1",
                    "name": "eth0",
                    "mac_address": "0C:C4:7A:04:91:7C"
                }
            },
            "transport": {
                "bound_address": "inet[/0:0:0:0:0:0:0:0:9300]",
                "publish_address": "inet[/10.20.149.158:9300]",
                "profiles": {}
            },
            "http": {
                "bound_address": "inet[/0:0:0:0:0:0:0:0:9200]",
                "publish_address": "inet[/10.20.149.158:9200]",
                "max_content_length_in_bytes": 104857600
            },
            "plugins": []
        }
    }
}

if __name__ == '__main__':
    import json
    print json.dumps(resCmd)
