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

resClusterStats = {
    "timestamp": 1515140759136,
    "cluster_name": "elasticsearch",
    "status": "green",
    "indices": {
        "count": 3,
        "shards": {
            "total": 15,
            "primaries": 15,
            "replication": 0,
            "index": {
                "shards": {
                    "min": 5,
                    "max": 5,
                    "avg": 5
                },
                "primaries": {
                    "min": 5,
                    "max": 5,
                    "avg": 5
                },
                "replication": {
                    "min": 0,
                    "max": 0,
                    "avg": 0
                }
            }
        },
        "docs": {
            "count": 10,
            "deleted": 0
        },
        "store": {
            "size_in_bytes": 40039,
            "throttle_time_in_millis": 0
        },
        "fielddata": {
            "memory_size_in_bytes": 0,
            "evictions": 0
        },
        "filter_cache": {
            "memory_size_in_bytes": 132,
            "evictions": 0
        },
        "id_cache": {
            "memory_size_in_bytes": 0
        },
        "completion": {
            "size_in_bytes": 0
        },
        "segments": {
            "count": 10,
            "memory_in_bytes": 52564,
            "index_writer_memory_in_bytes": 0,
            "index_writer_max_memory_in_bytes": 7680000,
            "version_map_memory_in_bytes": 0,
            "fixed_bit_set_memory_in_bytes": 0
        },
        "percolate": {
            "total": 0,
            "time_in_millis": 0,
            "current": 0,
            "memory_size_in_bytes": -1,
            "memory_size": "-1b",
            "queries": 0
        }
    },
    "nodes": {
        "count": {
            "total": 1,
            "master_only": 0,
            "data_only": 0,
            "master_data": 1,
            "client": 0
        },
        "versions": [
            "1.7.3"
        ],
        "os": {
            "available_processors": 8,
            "mem": {
                "total_in_bytes": 67584700416
            },
            "cpu": [
                {
                    "vendor": "Intel",
                    "model": "Xeon",
                    "mhz": 3699,
                    "total_cores": 8,
                    "total_sockets": 8,
                    "cores_per_socket": 32,
                    "cache_size_in_bytes": 10240,
                    "count": 1
                }
            ]
        },
        "process": {
            "cpu": {
                "percent": 0
            },
            "open_file_descriptors": {
                "min": 265,
                "max": 265,
                "avg": 265
            }
        },
        "jvm": {
            "max_uptime_in_millis": 1044574181,
            "versions": [
                {
                    "version": "1.8.0_72",
                    "vm_name": "Java HotSpot(TM) 64-Bit Server VM",
                    "vm_version": "25.72-b15",
                    "vm_vendor": "Oracle Corporation",
                    "count": 1
                }
            ],
            "mem": {
                "heap_used_in_bytes": 136690512,
                "heap_max_in_bytes": 1037959168
            },
            "threads": 90
        },
        "fs": {
            "total_in_bytes": 51470012416,
            "free_in_bytes": 5698859008,
            "available_in_bytes": 3077476352,
            "disk_reads": 1490266,
            "disk_writes": 44818691,
            "disk_io_op": 46308957,
            "disk_read_size_in_bytes": 27341878272,
            "disk_write_size_in_bytes": 182892331008,
            "disk_io_size_in_bytes": 210234209280,
            "disk_queue": "0",
            "disk_service_time": "0"
        },
        "plugins": []
    }
}

resStatsIndexing = {
    "_shards": {
        "total": 30,
        "successful": 15,
        "failed": 0
    },
    "_all": {
        "primaries": {
            "indexing": {
                "index_total": 0,
                "index_time_in_millis": 0,
                "index_current": 0,
                "delete_total": 0,
                "delete_time_in_millis": 0,
                "delete_current": 0,
                "noop_update_total": 0,
                "is_throttled": 'false',
                "throttle_time_in_millis": 0
            }
        },
        "total": {
            "indexing": {
                "index_total": 0,
                "index_time_in_millis": 0,
                "index_current": 0,
                "delete_total": 0,
                "delete_time_in_millis": 0,
                "delete_current": 0,
                "noop_update_total": 0,
                "is_throttled": 'false',
                "throttle_time_in_millis": 0
            }
        }
    },
    "indices": {
        "piratepu_pp2": {
            "primaries": {
                "indexing": {
                    "index_total": 0,
                    "index_time_in_millis": 0,
                    "index_current": 0,
                    "delete_total": 0,
                    "delete_time_in_millis": 0,
                    "delete_current": 0,
                    "noop_update_total": 0,
                    "is_throttled": 'false',
                    "throttle_time_in_millis": 0
                }
            },
            "total": {
                "indexing": {
                    "index_total": 0,
                    "index_time_in_millis": 0,
                    "index_current": 0,
                    "delete_total": 0,
                    "delete_time_in_millis": 0,
                    "delete_current": 0,
                    "noop_update_total": 0,
                    "is_throttled": 'false',
                    "throttle_time_in_millis": 0
                }
            }
        },
        "anarchop_apn": {
            "primaries": {
                "indexing": {
                    "index_total": 0,
                    "index_time_in_millis": 0,
                    "index_current": 0,
                    "delete_total": 0,
                    "delete_time_in_millis": 0,
                    "delete_current": 0,
                    "noop_update_total": 0,
                    "is_throttled": 'false',
                    "throttle_time_in_millis": 0
                }
            },
            "total": {
                "indexing": {
                    "index_total": 0,
                    "index_time_in_millis": 0,
                    "index_current": 0,
                    "delete_total": 0,
                    "delete_time_in_millis": 0,
                    "delete_current": 0,
                    "noop_update_total": 0,
                    "is_throttled": 'false',
                    "throttle_time_in_millis": 0
                }
            }
        },
        "readme": {
            "primaries": {
                "indexing": {
                    "index_total": 0,
                    "index_time_in_millis": 0,
                    "index_current": 0,
                    "delete_total": 0,
                    "delete_time_in_millis": 0,
                    "delete_current": 0,
                    "noop_update_total": 0,
                    "is_throttled": 'false',
                    "throttle_time_in_millis": 0
                }
            },
            "total": {
                "indexing": {
                    "index_total": 0,
                    "index_time_in_millis": 0,
                    "index_current": 0,
                    "delete_total": 0,
                    "delete_time_in_millis": 0,
                    "delete_current": 0,
                    "noop_update_total": 0,
                    "is_throttled": 'false',
                    "throttle_time_in_millis": 0
                }
            }
        }
    }
}

resCatIndices='''health status index        pri rep docs.count docs.deleted store.size pri.store.size 
yellow open   anarchop_apn   5   1          8            0     30.3kb         30.3kb 
yellow open   readme         5   1          1            0      4.4kb          4.4kb 
yellow open   piratepu_pp2   5   1          1            0      4.2kb          4.2kb'''

if __name__ == '__main__':
    import json
    print json.dumps(resClusterStats)
    print json.dumps(resStatsIndexing)
