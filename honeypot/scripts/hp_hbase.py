body="""<!--[if IE]>
<!DOCTYPE html>
<![endif]-->
<?xml version=\"1.0\" encoding=\"UTF-8\" ?>
<html lang=\"en\">
  <head>
    <meta charset=\"utf-8\">
    <title>Master: nnops_db01</title>
    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">
    <meta name=\"description\" content=\"\">
    <link href=\"/static/css/bootstrap.min.css\" rel=\"stylesheet\">
    <link href=\"/static/css/bootstrap-theme.min.css\" rel=\"stylesheet\">
    <link href=\"/static/css/hbase.css\" rel=\"stylesheet\">
  </head>

  <body>

    <div class=\"navbar  navbar-fixed-top navbar-default\">
        <div class=\"container\">
            <div class=\"navbar-header\">
                <button type=\"button\" class=\"navbar-toggle\" data-toggle=\"collapse\" data-target=\".navbar-collapse\">
                    <span class=\"icon-bar\"></span>
                    <span class=\"icon-bar\"></span>
                    <span class=\"icon-bar\"></span>
                </button>
                <a class=\"navbar-brand\" href=\"/master-status\"><img src=\"/static/hbase_logo_small.png\" alt=\"HBase Logo\"/></a>
            </div>
            <div class=\"collapse navbar-collapse\">
                <ul class=\"nav navbar-nav\">
                <li class=\"active\"><a href=\"/\">Home</a></li>
                <li><a href=\"/tablesDetailed.jsp\">Table Details</a></li>
                <li><a href=\"/logs/\">Local Logs</a></li>
                <li><a href=\"/logLevel\">Log Level</a></li>
                <li><a href=\"/dump\">Debug Dump</a></li>
                <li><a href=\"/jmx\">Metrics Dump</a></li>
                
                <li><a href=\"/conf\">HBase Configuration</a></li>
                
               </ul>
            </div><!--/.nav-collapse -->
        </div>
    </div>

    <div class=\"container\">
	
        <div class=\"row inner_header\">
            <div class=\"page-header\">
                <h1>Master <small>nnops_db01 </small></h1>
            </div>
        </div>

        <div class=\"row\">
        <!-- Various warnings that cluster admins should be aware of -->
        
        

        <section>
            <h2>Region Servers</h2>
            

<div class=\"tabbable\">
    <ul class=\"nav nav-pills\">
        <li class=\"active\"><a href=\"#tab_baseStats\" data-toggle=\"tab\">Base Stats</a></li>
        <li class=\"\"><a href=\"#tab_memoryStats\" data-toggle=\"tab\">Memory</a></li>
        <li class=\"\"><a href=\"#tab_requestStats\" data-toggle=\"tab\">Requests</a></li>
        <li class=\"\"><a href=\"#tab_storeStats\" data-toggle=\"tab\">Storefiles</a></li>
        <li class=\"\"><a href=\"#tab_compactStas\" data-toggle=\"tab\">Compactions</a></li>
    </ul>
    <div class=\"tab-content\" style=\"padding-bottom: 9px; border-bottom: 1px solid #ddd;\">
        <div class=\"tab-pane active\" id=\"tab_baseStats\">
            <table class=\"table table-striped\">
<tr>
    <th>ServerName</th>
    <th>Start time</th>
    <th>Requests Per Second</th>
    <th>Num. Regions</th>
</tr>
<tr>
    <td>
            <a href=\"//nnops_db01:33998/rs-status\">nnops_db01,34047,1521441995960</a>
        
</td>
    <td>Mon Mar 19 14:46:35 CST 2018</td>
    <td>0</td>
    <td>418</td>
</tr>
<tr><td>Total:1</td>
<td></td>
<td>0</td>
<td>418</td>
</tr>
</table>

        </div>
        <div class=\"tab-pane\" id=\"tab_memoryStats\">
            <table class=\"table table-striped\">
<tr>
    <th>ServerName</th>
    <th>Used Heap</th>
    <th>Max Heap</th>
    <th>Memstore Size</th>

</tr>
<tr>
    <td>
            <a href=\"//nnops_db01:33998/rs-status\">nnops_db01,34047,1521441995960</a>
        
</td>
    <td>393m</td>
    <td>1950m</td>
    <td>0m</td>

</tr>
</table>

        </div>
        <div class=\"tab-pane\" id=\"tab_requestStats\">
            <table class=\"table table-striped\">
<tr>
    <th>ServerName</th>
    <th>Request Per Second</th>
    <th>Read Request Count</th>
    <th>Write Request Count</th>
</tr>
<tr>
<td>
            <a href=\"//nnops_db01:33998/rs-status\">nnops_db01,34047,1521441995960</a>
        
</td>
<td>0</td>
<td>11943</td>
<td>5814</td>
</tr>
</table>

        </div>
        <div class=\"tab-pane\" id=\"tab_storeStats\">
            <table class=\"table table-striped\">
<tr>
    <th>ServerName</th>
    <th>Num. Stores</th>
    <th>Num. Storefiles</th>
    <th>Storefile Size Uncompressed</th>
    <th>Storefile Size</th>
    <th>Index Size</th>
    <th>Bloom Size</th>
</tr>
<tr>
<td>
            <a href=\"//nnops_db01:33998/rs-status\">nnops_db01,34047,1521441995960</a>
        
</td>
<td>418</td>
<td>581</td>
<td>634m</td>
<td>642mb</td>
<td>803k</td>
<td>1523k</td>
</tr>
</table>

        </div>
        <div class=\"tab-pane\" id=\"tab_compactStas\">
            <table class=\"table table-striped\">
<tr>
    <th>ServerName</th>
    <th>Num. Compacting KVs</th>
    <th>Num. Compacted KVs</th>
    <th>Remaining KVs</th>
    <th>Compaction Progress</th>
</tr>
<tr>
<td>
            <a href=\"//nnops_db01:33998/rs-status\">nnops_db01,34047,1521441995960</a>
        
</td>
<td>1008038</td>
<td>1008038</td>
<td>0</td>
<td>100.00%</td>
</tr>
</table>

        </div>
    </div>
</div>





            
                
<h2>Dead Region Servers</h2>
<table class=\"table table-striped\">
    <tr>
        <th></th>
        <th>ServerName</th>
        <th>Stop time</th>
    </tr>
    <tr>
    	<th></th>
        <td>nnops_db01,40862,1521432797018</td>
        <td>Mon Mar 19 14:46:42 CST 2018</td>
    </tr>
    <tr>
        <th>Total: </th>
        <td>servers: 1</td>
        <th></th>
    </tr>
</table>


            
        </section>
        <section>
            
    <h2>Backup Masters</h2>

    <table class=\"table table-striped\">
    <tr>
        <th>ServerName</th>
        <th>Port</th>
        <th>Start Time</th>
    </tr>
    <tr><td>Total:0</td>
    </table>


        </section>
        <section>
            <h2>Tables</h2>
            <div class=\"tabbable\">
                <ul class=\"nav nav-pills\">
                    <li class=\"active\">
                        <a href=\"#tab_userTables\" data-toggle=\"tab\">User Tables</a>
                    </li>
                    <li class=\"\">
                        <a href=\"#tab_catalogTables\" data-toggle=\"tab\">System Tables</a>
                    </li>
                    <li class=\"\">
                        <a href=\"#tab_userSnapshots\" data-toggle=\"tab\">Snapshots</a>
                    </li>
                </ul>
                <div class=\"tab-content\" style=\"padding-bottom: 9px; border-bottom: 1px solid #ddd;\">
                    <div class=\"tab-pane active\" id=\"tab_userTables\">
                        
                            
<table class=\"table table-striped\">
    <tr>
        <th>Namespace</th>
        <th>Table Name</th>
        
        <th>Online Regions</th>
        <th>Description</th>
    </tr>
    
    <tr>
        <td>default</td>
        <td><a href=table.jsp?name=AgentEvent>AgentEvent</a> </td>
        
        <td>1
        <td>'AgentEvent', {NAME => 'E', TTL => '1296000 SECONDS (15 DAYS)'}</td>
    </tr>
    
    <tr>
        <td>default</td>
        <td><a href=table.jsp?name=AgentInfo>AgentInfo</a> </td>
        
        <td>1
        <td>'AgentInfo', {NAME => 'Info', TTL => '1296000 SECONDS (15 DAYS)'}</td>
    </tr>
    
    <tr>
        <td>default</td>
        <td><a href=table.jsp?name=AgentLifeCycle>AgentLifeCycle</a> </td>
        
        <td>1
        <td>'AgentLifeCycle', {NAME => 'S', TTL => '1296000 SECONDS (15 DAYS)'}</td>
    </tr>
    
    <tr>
        <td>default</td>
        <td><a href=table.jsp?name=AgentStatV2>AgentStatV2</a> </td>
        
        <td>64
        <td>'AgentStatV2', {NAME => 'S', TTL => '1296000 SECONDS (15 DAYS)'}</td>
    </tr>
    
    <tr>
        <td>default</td>
        <td><a href=table.jsp?name=ApiMetaData>ApiMetaData</a> </td>
        
        <td>8
        <td>'ApiMetaData', {NAME => 'Api', TTL => '5184000 SECONDS (60 DAYS)'}</td>
    </tr>
    
    <tr>
        <td>default</td>
        <td><a href=table.jsp?name=ApplicationIndex>ApplicationIndex</a> </td>
        
        <td>1
        <td>'ApplicationIndex', {NAME => 'Agents', TTL => '1296000 SECONDS (15 DAYS)'}</td>
    </tr>
    
    <tr>
        <td>default</td>
        <td><a href=table.jsp?name=ApplicationMapStatisticsCallee_Ver2>ApplicationMapStatisticsCallee_Ver2</a> </td>
        
        <td>16
        <td>'ApplicationMapStatisticsCallee_Ver2', {NAME => 'C', TTL => '1296000 SECONDS (15 DAYS)'}</td>
    </tr>
    
    <tr>
        <td>default</td>
        <td><a href=table.jsp?name=ApplicationMapStatisticsCaller_Ver2>ApplicationMapStatisticsCaller_Ver2</a> </td>
        
        <td>16
        <td>'ApplicationMapStatisticsCaller_Ver2', {NAME => 'C', TTL => '1296000 SECONDS (15 DAYS)'}</td>
    </tr>
    
    <tr>
        <td>default</td>
        <td><a href=table.jsp?name=ApplicationMapStatisticsSelf_Ver2>ApplicationMapStatisticsSelf_Ver2</a> </td>
        
        <td>8
        <td>'ApplicationMapStatisticsSelf_Ver2', {NAME => 'C', TTL => '1296000 SECONDS (15 DAYS)'}</td>
    </tr>
    
    <tr>
        <td>default</td>
        <td><a href=table.jsp?name=ApplicationTraceIndex>ApplicationTraceIndex</a> </td>
        
        <td>16
        <td>'ApplicationTraceIndex', {NAME => 'I', TTL => '1296000 SECONDS (15 DAYS)'}</td>
    </tr>
    
    <tr>
        <td>default</td>
        <td><a href=table.jsp?name=HostApplicationMap_Ver2>HostApplicationMap_Ver2</a> </td>
        
        <td>4
        <td>'HostApplicationMap_Ver2', {NAME => 'M', TTL => '1296000 SECONDS (15 DAYS)'}</td>
    </tr>
    
    <tr>
        <td>default</td>
        <td><a href=table.jsp?name=SqlMetaData_Ver2>SqlMetaData_Ver2</a> </td>
        
        <td>16
        <td>'SqlMetaData_Ver2', {NAME => 'Sql', TTL => '5184000 SECONDS (60 DAYS)'}</td>
    </tr>
    
    <tr>
        <td>default</td>
        <td><a href=table.jsp?name=StringMetaData>StringMetaData</a> </td>
        
        <td>8
        <td>'StringMetaData', {NAME => 'Str', TTL => '5184000 SECONDS (60 DAYS)'}</td>
    </tr>
    
    <tr>
        <td>default</td>
        <td><a href=table.jsp?name=TraceV2>TraceV2</a> </td>
        
        <td>256
        <td>'TraceV2', {NAME => 'S', TTL => '1296000 SECONDS (15 DAYS)'}</td>
    </tr>
    
    <p>14 table(s) in set. [<a href=tablesDetailed.jsp>Details</a>]</p>
</table>


                        
                    </div>
                    <div class=\"tab-pane\" id=\"tab_catalogTables\">
                        
                            
<table class=\"table table-striped\">
<tr>
    <th>Table Name</th>
    
    <th>Description</th>
</tr>

<tr>
<td><a href=\"table.jsp?name=hbase:meta\">hbase:meta</a></td>
    
    <td>The hbase:meta table holds references to all User Table regions</td>
</tr>

<tr>
<td><a href=\"table.jsp?name=hbase:namespace\">hbase:namespace</a></td>
    
    <td>The .NAMESPACE. table holds information about namespaces.</td>
</tr>

</table>


                        
                    </div>
                    <div class=\"tab-pane\" id=\"tab_userSnapshots\">
                        

                    </div>
                </div>
            </div>
        </section>
        
        


        
	


        <section>
            
<h2>Tasks</h2>
  <ul class=\"nav nav-pills\">
    <li ><a href=\"?filter=all\">Show All Monitored Tasks</a></li>
    <li class=\"active\"><a href=\"?filter=general\">Show non-RPC Tasks</a></li>
    <li ><a href=\"?filter=handler\">Show All RPC Handler Tasks</a></li>
    <li ><a href=\"?filter=rpc\">Show Active RPC Calls</a></li>
    <li ><a href=\"?filter=operation\">Show Client Operations</a></li>
    <li><a href=\"?format=json&filter=general\">View as JSON</a></li>
  </ul>
  
    <p>No tasks currently running on this node.</p>
  




        </section>

        <section>
            <h2>Software Attributes</h2>
            <table id=\"attributes_table\" class=\"table table-striped\">
                <tr>
                    <th>Attribute Name</th>
                    <th>Value</th>
                    <th>Description</th>
                </tr>
                <tr>
                    <td>HBase Version</td>
                    <td>1.0.3, revision=f1e1312f9790a7c40f6a4b5a1bab2ea1dd559890</td><td>HBase version and revision</td>
                </tr>
                <tr>
                    <td>HBase Compiled</td>
                    <td>Tue Jan 19 19:26:53 PST 2016, enis</td>
                    <td>When HBase version was compiled and by whom</td>
                </tr>
                <tr>
                    <td>HBase Source Checksum</td>
                    <td>c4e192f054afd9e526a93f3c19f65a22</td>
                    <td>HBase source MD5 checksum</td>
                </tr>
                <tr>
                    <td>Hadoop Version</td>
                    <td>2.5.1, revision=2e18d179e4a8065b6a9f29cf2de9451891265cce</td>
                    <td>Hadoop version and revision</td>
                </tr>
                <tr>
                    <td>Hadoop Compiled</td>
                    <td>2014-09-05T23:05Z, kasha</td>
                    <td>When Hadoop version was compiled and by whom</td>
                </tr>
                <tr>
                    <td>Hadoop Source Checksum</td>
                    <td>6424fcab95bfff8337780a181ad7c78</td>
                    <td>Hadoop source MD5 checksum</td>
                </tr>
                <tr>
                    <td>ZooKeeper Client Version</td>
                    <td>3.4.6, revision=1569965</td>
                    <td>ZooKeeper client version and revision</td>
                </tr>
                <tr>
                    <td>ZooKeeper Client Compiled</td>
                    <td>02/20/2014 09:09 GMT</td>
                    <td>When ZooKeeper client version was compiled</td>
                </tr>
                <tr>
                    <td>Zookeeper Quorum</td>
                    <td> localhost:2181 </td>
                    <td>Addresses of all registered ZK servers. For more, see <a href=\"/zk.jsp\">zk dump</a>.</td>
                </tr>
                <tr>
                    <td>Zookeeper Base Path</td>
                    <td> /hbase</td>
                    <td>Root node of this cluster in ZK.</td>
                </tr>
                <tr>
                    <td>HBase Root Directory</td>
                    <td>file:/software/hbase/data/hbase</td>
                    <td>Location of HBase home directory</td>
                </tr>
                <tr>
                    <td>HMaster Start Time</td>
                    <td>Mon Mar 19 14:46:34 CST 2018</td>
                    <td>Date stamp of when this HMaster was started</td>
                </tr>
                
	                <tr>
	                    <td>HMaster Active Time</td>
	                    <td>Mon Mar 19 14:46:35 CST 2018</td>
	                    <td>Date stamp of when this HMaster became active</td>
	                </tr>
	                <tr>
	                    <td>HBase Cluster ID</td>
	                    <td>1b07f0f9-f34e-46b9-aa3d-d3c81c425f2c</td>
	                    <td>Unique identifier generated for each HBase cluster</td>
	                </tr>
	                <tr>
	                    <td>Load average</td>
	                    <td>418.00</td>
	                    <td>Average number of regions per regionserver. Naive computation.</td>
	                </tr>
	                
	                <tr>
	                    <td>Coprocessors</td>
	                    <td>[]</td>
	                    <td>Coprocessors currently loaded by the master</td>
	                </tr>
                
            </table>
        </section>
        </div>
    </div> <!-- /container -->

    <script src=\"/static/js/jquery.min.js\" type=\"text/javascript\"></script>
    <script src=\"/static/js/bootstrap.min.js\" type=\"text/javascript\"></script>
    <script src=\"/static/js/tab.js\" type=\"text/javascript\"></script>
  </body>
</html>"""