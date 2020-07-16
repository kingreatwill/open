# Spark-submit 参数调优完整攻略

### --sparksubmit

--num-executors  

该参数主要用于设置该应用总共需要多少executors来执行，Driver在向集群资源管理器申请资源时需要根据此参数决定分配的Executor个数，并尽量满足所需。在不带的情况下只会分配少量Executor。这个值得设置还是要看分配的队列的资源情况，太少了无法充分利用集群资源，太多了则难以分配需要的资源。



--executor-memory

设置每个executor的内存，对Spark作业运行的性能影响很大。一般4-8G就差不多了，当然还要看资源队列的情况。num-executor*executor-memory的大小绝不能超过队列的内存总大小。



--executor-cores

设置每个executor的cpu核数，其决定了每个executor并行执行task的能力。Executor的CPU core数量设置为2-4个即可。但要注意，num-executor*executor-cores也不能超过分配队列中cpu核数的大小。具体的核数的设置需要根据分配队列中资源统筹考虑，取得Executor，核数，及任务数的平衡。对于多任务共享的队列，更要注意不能将资源占满    



--driver-memory

运行sparkContext的Driver所在所占用的内存，通常不必设置，设置的话1G就足够了，除非是需要使用collect之类算子经常需要将数据提取到driver中的情况。

  --total-executor-cores    

是所有executor总共使用的cpu核数 standalone default all cores



### --conf

 --conf spark.default.parallelism

此参数用于设置每个stage经TaskScheduler进行调度时生成task的数量，此参数未设置时将会根据读到的RDD的分区生成task，即根据源数据在hdfs中的分区数确定，若此分区数较小，则处理时只有少量task在处理，前述分配的executor中的core大部分无任务可干。通常可将此值设置为num-executors*executor-cores的2-3倍为宜，如果与其相近的话，则对于先完成task的core则无任务可干。2-3倍数量关系的话即不至于太零散，又可使得任务执行更均衡。



--conf spark.storage.memoryFraction

参数说明：该参数用于设置RDD持久化数据在Executor内存中能占的比例，默认是0.6。也就是说，默认Executor 60%的内存，可以用来保存持久化的RDD数据。根据你选择的不同的持久化策略，如果内存不够时，可能数据就不会持久化，或者数据会写入磁盘。

参数调优建议：如果Spark作业中，有较多的RDD持久化操作，该参数的值可以适当提高一些，保证持久化的数据能够容纳在内存中。避免内存不够缓存所有的数据，导致数据只能写入磁盘中，降低了性能。但是如果Spark作业中的shuffle类操作比较多，而持久化操作比较少，那么这个参数的值适当降低一些比较合适。此外，如果发现作业由于频繁的gc导致运行缓慢（通过spark web ui可以观察到作业的gc耗时），意味着task执行用户代码的内存不够用，那么同样建议调低这个参数的值。个人不太建议调该参数



--conf spark.shuffle.memoryFraction

参数说明：该参数用于设置shuffle过程中一个task拉取到上个stage的task的输出后，进行聚合操作时能够使用的Executor内存的比例，默认是0.2。也就是说，Executor默认只有20%的内存用来进行该操作。shuffle操作在进行聚合时，如果发现使用的内存超出了这个20%的限制，那么多余的数据就会溢写到磁盘文件中去，此时就会极大地降低性能。

参数调优建议：如果Spark作业中的RDD持久化操作较少，shuffle操作较多时，建议降低持久化操作的内存占比，提高shuffle操作的内存占比比例，避免shuffle过程中数据过多时内存不够用，必须溢写到磁盘上，降低了性能。此外，如果发现作业由于频繁的gc导致运行缓慢，意味着task执行用户代码的内存不够用，那么同样建议调低这个参数的值。个人不太建议调该参数



--conf spark.sql.codegen

默认值为false，当它设置为true时，Spark SQL会把每条查询的语句在运行时编译为java的二进制代码。这有什么作用呢？它可以提高大型查询的性能，但是如果进行小规模的查询的时候反而会变慢，就是说直接用查询反而比将它编译成为java的二进制代码快。所以在优化这个选项的时候要视情况而定。

这个选项可以让Spark SQL把每条查询语句在运行前编译为java二进制代码，由于生成了专门运行指定查询的代码，codegen可以让大型查询或者频繁重复的查询明显变快，然而在运行特别快（1-2秒）的即时查询语句时，codegen就可能增加额外的开销（将查询语句编译为java二进制文件）。codegen还是一个实验性的功能，但是在大型的或者重复运行的查询中使用codegen



--conf spark.sql.inMemoryColumnStorage.compressed

默认值为false 它的作用是自动对内存中的列式存储进行压缩



--conf spark.sql.inMemoryColumnStorage.batchSize    

默认值为1000 这个参数代表的是列式缓存时的每个批处理的大小。如果将这个值调大可能会导致内存不够的异常，所以在设置这个的参数的时候得注意你的内存大小

在缓存SchemaRDD（Row RDD）时，Spark SQL会安照这个选项设定的大小（默认为1000）把记录分组，然后分批次压缩。

太小的批处理会导致压缩比过低，而太大的话，比如当每个批处理的数据超过内存所能容纳的大小时，也有可能引发问题。

如果你表中的记录比价大（包含数百个字段或者包含像网页这样非常大的字符串字段），就可能需要调低批处理的大小来避免内存不够（OOM）的错误。如果不是在这样的场景下，默认的批处理 的大小是比较合适的，因为压缩超过1000条压缩记录时也基本无法获得更高的压缩比了。



--conf spark.sql.parquet.compressed.codec

默认值为snappy 这个参数代表使用哪种压缩编码器。可选的选项包括uncompressed/snappy/gzip/lzo        uncompressed这个顾名思义就是不用压缩的意思



--conf spark.speculation 

推测执行优化机制采用了典型的以空间换时间的优化策略，它同时启动多个相同task（备份任务）处理相同的数据块，哪个完成的早，则采用哪个task的结果，这样可防止拖后腿Task任务出现，进而提高作业计算速度，但是，这样却会占用更多的资源，在集群资源紧缺的情况下，设计合理的推测执行机制可在多用少量资源情况下，减少大作业的计算时间。

检查逻辑代码中注释很明白，当成功的Task数超过总Task数的75%(可通过参数spark.speculation.quantile设置)时，再统计所有成功的Tasks的运行时间，得到一个中位数，用这个中位数乘以1.5(可通过参数spark.speculation.multiplier控制)得到运行时间门限，如果在运行的Tasks的运行时间超过这个门限，则对它启用推测。简单来说就是对那些拖慢整体进度的Tasks启用推测，以加速整个Stage的运行

    - spark.speculation.interval    100毫秒    Spark经常检查要推测的任务。

    - spark.speculation.multiplier    1.5    任务的速度比投机的中位数慢多少倍。

    - spark.speculation.quantile    0.75    在为特定阶段启用推测之前必须完成的任务的分数。

        

--conf spark.shuffle.consolidateFiles

默认值：false

参数说明：如果使用HashShuffleManager，该参数有效。如果设置为true，那么就会开启consolidate机制，会大幅度合并shuffle write的输出文件，对于shuffle read task数量特别多的情况下，这种方法可以极大地减少磁盘IO开销，提升性能。

调优建议：如果的确不需要SortShuffleManager的排序机制，那么除了使用bypass机制，还可以尝试将spark.shffle.manager参数手动指定为hash，使用HashShuffleManager，同时开启consolidate机制。在实践中尝试过，发现其性能比开启了bypass机制的SortShuffleManager要高出10%~30%。    



--conf spark.shuffle.file.buffer    

默认值：32k

参数说明：该参数用于设置shuffle write task的BufferedOutputStream的buffer缓冲大小。将数据写到磁盘文件之前，会先写入buffer缓冲中，待缓冲写满之后，才会溢写到磁盘。

调优建议：如果作业可用的内存资源较为充足的话，可以适当增加这个参数的大小（比如64k，一定是成倍的增加），从而减少shuffle write过程中溢写磁盘文件的次数，也就可以减少磁盘IO次数，进而提升性能。在实践中发现，合理调节该参数，性能会有1%~5%的提升。



--conf spark.reducer.maxSizeInFlight

默认值：48m

参数说明：该参数用于设置shuffle read task的buffer缓冲大小，而这个buffer缓冲决定了每次能够拉取多少数据。

调优建议：如果作业可用的内存资源较为充足的话，可以适当增加这个参数的大小（比如96m），从而减少拉取数据的次数，也就可以减少网络传输的次数，进而提升性能。在实践中发现，合理调节该参数，性能会有1%~5%的提升。



--conf spark.shuffle.io.maxRetries

默认值：3

参数说明：shuffle read task从shuffle write task所在节点拉取属于自己的数据时，如果因为网络异常导致拉取失败，是会自动进行重试的。该参数就代表了可以重试的最大次数。如果在指定次数之内拉取还是没有成功，就可能会导致作业执行失败。

调优建议：对于那些包含了特别耗时的shuffle操作的作业，建议增加重试最大次数（比如60次），以避免由于JVM的full gc或者网络不稳定等因素导致的数据拉取失败。在实践中发现，对于针对超大数据量（数十亿~上百亿）的shuffle过程，调节该参数可以大幅度提升稳定性。

taskScheduler不负责重试task，由DAGScheduler负责重试stage



--conf spark.shuffle.io.retryWait

默认值：5s

参数说明：具体解释同上，该参数代表了每次重试拉取数据的等待间隔，默认是5s。

调优建议：建议加大间隔时长（比如60s），以增加shuffle操作的稳定性。



--conf spark.shuffle.memoryFraction

默认值：0.2

参数说明：该参数代表了Executor内存中，分配给shuffle read task进行聚合操作的内存比例，默认是20%。

调优建议：在资源参数调优中讲解过这个参数。如果内存充足，而且很少使用持久化操作，建议调高这个比例，给shuffle read的聚合操作更多内存，以避免由于内存不足导致聚合过程中频繁读写磁盘。在实践中发现，合理调节该参数可以将性能提升10%左右。    



--conf spark.shuffle.manager

默认值：sort|hash

参数说明：该参数用于设置ShuffleManager的类型。Spark 1.5以后，有三个可选项：hash、sort和tungsten-sort。HashShuffleManager是Spark 1.2以前的默认选项，但是Spark 1.2以及之后的版本默认都是SortShuffleManager了。tungsten-sort与sort类似，但是使用了tungsten计划中的堆外内存管理机制，内存使用效率更高。

调优建议：由于SortShuffleManager默认会对数据进行排序，因此如果你的业务逻辑中需要该排序机制的话，则使用默认的SortShuffleManager就可以；而如果你的业务逻辑不需要对数据进行排序，那么建议参考后面的几个参数调优，通过bypass机制或优化的HashShuffleManager来避免排序操作，同时提供较好的磁盘读写性能。这里要注意的是，tungsten-sort要慎用，因为之前发现了一些相应的bug。



--conf spark.shuffle.sort.bypassMergeThreshold

默认值：200

参数说明：当ShuffleManager为SortShuffleManager时，如果shuffle read task的数量小于这个阈值（默认是200），则shuffle write过程中不会进行排序操作，而是直接按照未经优化的HashShuffleManager的方式去写数据，但是最后会将每个task产生的所有临时磁盘文件都合并成一个文件，并会创建单独的索引文件。

调优建议：当你使用SortShuffleManager时，如果的确不需要排序操作，那么建议将这个参数调大一些，大于shuffle read task的数量。那么此时就会自动启用bypass机制，map-side就不会进行排序了，减少了排序的性能开销。但是这种方式下，依然会产生大量的磁盘文件，因此shuffle write性能有待提高。



--conf spark.shuffle.consolidateFiles

默认值：false

参数说明：如果使用HashShuffleManager，该参数有效。如果设置为true，那么就会开启consolidate机制，会大幅度合并shuffle write的输出文件，对于shuffle read task数量特别多的情况下，这种方法可以极大地减少磁盘IO开销，提升性能。

调优建议：如果的确不需要SortShuffleManager的排序机制，那么除了使用bypass机制，还可以尝试将spark.shffle.manager参数手动指定为hash，使用HashShuffleManager，同时开启consolidate机制。在实践中尝试过，发现其性能比开启了bypass机制的SortShuffleManager要高出10%~30%。