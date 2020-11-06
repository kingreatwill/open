# jmeter

https://jmeter.apache.org/
http://www.jmeter.com.cn/

https://www.cnblogs.com/poloyy/tag/Jmeter/

## 其它产品

- LoadRunner（商用版）
- Load impact（免费使用）
  https://github.com/loadimpact/
  https://github.com/loadimpact/k6
  
- Locust
  https://github.com/locustio/locust python 14.2k

- OpenSTA
  http://www.opensta.org/

## 下载安装

JMETER_HOME

和 Path

jmeter.bat
run JMeter (in GUI mode by default)
jmeterw.cmd （其实就是 javaw.exe 运行）
run JMeter without the windows shell console (in GUI mode by default)
jmeter-n.cmd
drop a JMX file on this to run a CLI mode test
jmeter-n-r.cmd
drop a JMX file on this to run a CLI mode test remotely
jmeter-t.cmd
drop a JMX file on this to load it in GUI mode
jmeter-server.bat
start JMeter in server mode
mirror-server.cmd
runs the JMeter Mirror Server in CLI mode
shutdown.cmd
Run the Shutdown client to stop a CLI mode instance gracefully
stoptest.cmd
Run the Shutdown client to stop a CLI mode instance abruptly

## 使用

[demo jmx](yml/登录获取token.jmx)

## Grafana+Jmeter+Influxdb（推荐）

看压测过程中参数的变化
https://www.cnblogs.com/poloyy/p/12219145.html

https://jmeter.apache.org/usermanual/component_reference.html#Backend_Listener

- Graphite
- Influxdb

## Command

```
bin>jmeter -?
    _    ____   _    ____ _   _ _____       _ __  __ _____ _____ _____ ____
   / \  |  _ \ / \  / ___| | | | ____|     | |  \/  | ____|_   _| ____|  _ \
  / _ \ | |_) / _ \| |   | |_| |  _|    _  | | |\/| |  _|   | | |  _| | |_) |
 / ___ \|  __/ ___ \ |___|  _  | |___  | |_| | |  | | |___  | | | |___|  _ <
/_/   \_\_| /_/   \_\____|_| |_|_____|  \___/|_|  |_|_____| |_| |_____|_| \_\ 5.3

Copyright (c) 1999-2020 The Apache Software Foundation

        --?
                print command line options and exit
        -h, --help
                print usage information and exit
        -v, --version
                print the version information and exit
        -p, --propfile <argument>
                the jmeter property file to use
        -q, --addprop <argument>
                additional JMeter property file(s)
        -t, --testfile <argument>
                the jmeter test(.jmx) file to run. "-t LAST" will load last
                used file
        -l, --logfile <argument>
                the file to log samples to
        -i, --jmeterlogconf <argument>
                jmeter logging configuration file (log4j2.xml)
        -j, --jmeterlogfile <argument>
                jmeter run log file (jmeter.log)
        -n, --nongui
                run JMeter in nongui mode
        -s, --server
                run the JMeter server
        -E, --proxyScheme <argument>
                Set a proxy scheme to use for the proxy server
        -H, --proxyHost <argument>
                Set a proxy server for JMeter to use
        -P, --proxyPort <argument>
                Set proxy server port for JMeter to use
        -N, --nonProxyHosts <argument>
                Set nonproxy host list (e.g. *.apache.org|localhost)
        -u, --username <argument>
                Set username for proxy server that JMeter is to use
        -a, --password <argument>
                Set password for proxy server that JMeter is to use
        -J, --jmeterproperty <argument>=<value>
                Define additional JMeter properties
        -G, --globalproperty <argument>=<value>
                Define Global properties (sent to servers)
                e.g. -Gport=123
                 or -Gglobal.properties
        -D, --systemproperty <argument>=<value>
                Define additional system properties
        -S, --systemPropertyFile <argument>
                additional system property file(s)
        -f, --forceDeleteResultFile
                force delete existing results files and web report folder if
                 present before starting the test
        -L, --loglevel <argument>=<value>
                [category=]level e.g. jorphan=INFO, jmeter.util=DEBUG or com
                .example.foo=WARN
        -r, --runremote
                Start remote servers (as defined in remote_hosts)
        -R, --remotestart <argument>
                Start these remote servers (overrides remote_hosts)
        -d, --homedir <argument>
                the jmeter home directory to use
        -X, --remoteexit
                Exit the remote servers at end of test (non-GUI)
        -g, --reportonly <argument>
                generate report dashboard only, from a test results file
        -e, --reportatendofloadtests
                generate report dashboard after load test
        -o, --reportoutputfolder <argument>
                output folder for report dashboard
```

jmeter -h

```
...

To run Apache JMeter in NON_GUI mode:
Open a command prompt (or Unix shell) and type:

jmeter.bat(Windows)/jmeter.sh(Linux) -n -t test-file [-p property-file] [-l results-file] [-j log-file]

--------------------------------------------------

To run Apache JMeter in NON_GUI mode and generate a report at end :
Open a command prompt (or Unix shell) and type:

jmeter.bat(Windows)/jmeter.sh(Linux) -n -t test-file [-p property-file] [-l results-file] [-j log-file] -e -o [Path to output folder]

--------------------------------------------------
To generate a Report from existing CSV file:
Open a command prompt (or Unix shell) and type:

jmeter.bat(Windows)/jmeter.sh(Linux) -g [csv results file] -o [path to output folder (empty or not existing)]

...
```

jmeter -n -t script.jmx -r
jmeter -n -t script.jmx -R server1,server2...

- n to start Jmeter in a non-gui mode
- t to define the test file
- r to start the remote server as defined in the [JMeter properties file](https://datacadamia.com/jmeter/remote#property)
- R to define the list of servers and start them
- -Gproperty=value - define a property in all the servers (may appear more than once)
- -Z - Exit remote servers at the end of the test

[JMeter - 2.9 - (Remote Test|Distributed testing)](https://datacadamia.com/jmeter/remote)

## docker

https://github.com/justb4/docker-jmeter
https://github.com/pedrocesar-ti/distributed-jmeter-docker

https://www.github.com/cagdascirit/docker-jmeter
