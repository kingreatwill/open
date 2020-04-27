# ETL
Extract-Transform-Load的缩写，即数据抽取、转换、装载的过程

ETL，是英文Extract-Transform-Load的缩写，用来描述将数据从来源端经过萃取（extract）、转置（transform）、加载（load）至目的端的过程。ETL一词较常用在数据仓库(Data Warehouse简称DW)，但其对象并不限于数据仓库。

ETL是将业务系统的数据经过抽取（Extract）、清洗转换（Cleaning、Transform）之后加载（Load）到数据仓库的过程，目的是将企业中的分散、零乱、标准不统一的数据整合到一起，为企业的决策提供分析依据。 ETL是商业智能（BI，Business Intelligence）项目重要的一个环节。

BI（Business Intelligence）即商务智能，它是一套完整的解决方案，用来将企业中现有的数据进行有效的整合，快速准确地提供报表并提出决策依据，帮助企业做出明智的业务经营决策。

## kettle
https://github.com/pentaho/pentaho-kettle

下载：http://www.kettle.be/

or(好像只更新到7.1 没有最新版本了) https://sourceforge.net/projects/pentaho/files/Data%20Integration/ 

https://www.toutiao.com/a6818008846477296140

将mysql驱动jar(我这里使用的是5.1.40版本，mysql-connector-java-5.1.40.jar)放到pdi-ce-7.1.0.0-12/data-integration/lib目录下

