package memcached

const (
	commandVersion = "VERSION 1.4.13\r\n"
	commandGet     = "VALUE %s 0 7\r\njeffery\r\nEND\r\n"
	commandGets    = "VALUE %s 0 10 50\r\nsome_value\r\nEND\r\n"

	commandStats = "STAT pid 2080\r\nSTAT uptime 3151236\r\nSTAT time 1520550684\r\nSTAT version 1.4.13\r\nSTAT libevent 2.0.16-stable\r\nSTAT pointer_size 64\r\nSTAT rusage_user 371.247201\r\nSTAT rusage_system 1839.982991\r\nSTAT curr_connections 8\r\nSTAT total_connections 5547233\r\nSTAT connection_structures 55\r\nSTAT reserved_fds 20\r\nSTAT cmd_get 22076096\r\nSTAT cmd_set 21\r\nSTAT cmd_flush 3\r\nSTAT cmd_touch 0\r\nSTAT get_hits 22076066\r\nSTAT get_misses 30\r\nSTAT delete_misses 0\r\nSTAT delete_hits 0\r\nSTAT incr_misses 0\r\nSTAT incr_hits 0\r\nSTAT decr_misses 0\r\nSTAT decr_hits 0\r\nSTAT cas_misses 0\r\nSTAT cas_hits 0\r\nSTAT cas_badval 0\r\nSTAT touch_hits 0\r\nSTAT touch_misses 0\r\nSTAT auth_cmds 0\r\nSTAT auth_errors 0\r\nSTAT bytes_read 286857265\r\nSTAT bytes_written 129670828957\r\nSTAT limit_maxbytes 67108864\r\nSTAT accepting_conns 1\r\nSTAT listen_disabled_num 0\r\nSTAT threads 4\r\nSTAT conn_yields 0\r\nSTAT hash_power_level 16\r\nSTAT hash_bytes 524288\r\nSTAT hash_is_expanding 0\r\nSTAT expired_unfetched 0\r\nSTAT evicted_unfetched 0\r\nSTAT bytes 29828\r\nSTAT curr_items 5\r\nSTAT total_items 21\r\nSTAT evictions 0\r\nSTAT reclaimed 3\r\nEND\r\n"
)
