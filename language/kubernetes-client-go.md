[client-go 和 controller的架构设计](https://github.com/kubernetes/sample-controller/blob/master/docs/controller-client-go.md)
![](../img/lng/client-go-controller-interaction.jpeg)

#### client-go组件
- Reflector：通过Kubernetes API监控Kubernetes的资源类型
   - 采用List、Watch机制
   - 可以Watch任何资源包括CRD
   - 添加object对象到FIFO队列，然后Informer会从队列里面取数据
- Informer：controller机制的基础
   - 循环处理object对象
   - 从Reflector取出数据，然后将数据给到Indexer去缓存
   - 提供对象事件的handler接口
- Indexer：提供object对象的索引，是线程安全的，缓存对象信息

#### controller组件

- Informer reference: controller需要创建合适的Informer才能通过Informer reference操作资源对象

- Indexer reference: controller创建Indexer reference然后去利用索引做相关处理

- Resource Event Handlers：Informer会回调这些handlers

- Work queue: Resource Event Handlers被回调后将key写到工作队列
这里的key相当于事件通知，后面根据取出事件后，做后续的处理

- Process Item：从工作队列中取出key后进行后续处理，具体处理可以通过Indexer reference

- controller可以直接创建上述两个引用对象去处理，也可以采用工厂模式，官方都有相关示例