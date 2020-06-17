http://spark.apache.org/docs/latest/running-on-kubernetes.html

https://github.com/GoogleCloudPlatform/spark-on-k8s-operator
https://github.com/GoogleCloudPlatform/spark-on-k8s-operator/blob/master/docs/quick-start-guide.md
[Spark in action on Kubernetes - Spark Operator的原理解析](https://yq.aliyun.com/articles/695315)


[Spark on K8S环境部署细节](https://www.cnblogs.com/lanrish/p/12267623.html)



![](img/k8s-cluster-mode.png)
上面这张图是Spark中kubernetes的集成图，也就是说当我们通过spark-submit提交作业的时候，会自动生成driver pod与exector pods