[TOC]
# Software release
[Blue-green Deployments, A/B Testing, and Canary Releases](https://blog.christianposta.com/deploy/blue-green-deployments-a-b-testing-and-canary-releases/)
[蓝绿部署、金丝雀发布（灰度发布）、A/B测试的准确定义](https://www.cnblogs.com/lijiaocn/p/9986880.html)

## 蓝绿发布(Blue-green deployment)
蓝绿部署中，一共有两套系统：一套是正在提供服务系统(也就是上面说的旧版)，标记为“绿色”；另一套是准备发布的系统，标记为“蓝色”。两套系统都是功能完善的，并且正在运行的系统，只是系统版本和对外服务情况不同。正在对外提供服务的老系统是绿色系统，新部署的系统是蓝色系统。

蓝色系统不对外提供服务，用来做啥？

用来做发布前测试，测试过程中发现任何问题，可以直接在蓝色系统上修改，不干扰用户正在使用的系统。

蓝色系统经过反复的测试、修改、验证，确定达到上线标准之后，直接将用户切换到蓝色系统, 切换后的一段时间内，依旧是蓝绿两套系统并存，但是用户访问的已经是蓝色系统。这段时间内观察蓝色系统（新系统）工作状态，如果出现问题，直接切换回绿色系统。

当确信对外提供服务的蓝色系统工作正常，不对外提供服务的绿色系统已经不再需要的时候，蓝色系统正式成为对外提供服务系统，成为新的绿色系统。 原先的绿色系统可以销毁，将资源释放出来，用于部署下一个蓝色系统。

### 蓝绿发布特点
- 蓝绿部署的目的是减少发布时的中断时间、能够快速撤回发布。
- 两套系统没有耦合的时候才能百分百保证不干扰

### 蓝绿发布注意事项
蓝绿部署只是上线策略中的一种，它不是可以应对所有情况的万能方案。 蓝绿部署能够简单快捷实施的前提假设是目标系统是非常内聚的(**简单的说单体应用就很适合**)，如果目标系统相当复杂，那么如何切换、两套系统的数据是否需要以及如何同步等，都需要仔细考虑。

当你切换到蓝色环境时，需要妥当处理未完成的业务和新的业务。如果你的数据库后端无法处理，会是一个比较麻烦的问题；

- 可能会出现需要同时处理“微服务架构应用”和“传统架构应用”的情况，如果在蓝绿部署中协调不好这两者，还是有可能会导致服务停止。
- 需要提前考虑数据库与应用部署同步迁移 /回滚的问题。
- 蓝绿部署需要有基础设施支持。
- 在非隔离基础架构（ VM 、 Docker 等）上执行蓝绿部署，蓝色环境和绿色环境有被摧毁的风险。

## 滚动发布（Rolling release）
一般是取出一个或者多个服务器停止服务，执行更新，并重新将其投入使用。周而复始，直到集群中所有的实例都更新成新版本。

发布流程:
相对于蓝绿发布需要一套完备的机器不同, 滚动发布只需要一台机器(这儿这是为了理解, 实际可能是多台), 我们只需要将部分功能部署在这台机器上, 然后去替换正在运行的机器, 将更新后的功能部署在Server1 上, 然后Server1去替换正在运行的Server, 替换下来的物理机又可以继续部署Server2的新版本, 然后去替换正在工作的Server2 , 以此类推, 直到替换完所有的服务器, 至此 ,服务更新完成。

### 滚动发布特点
- 这种部署方式相对于蓝绿部署，更加节约资源——它不需要运行两个集群、两倍的实例数。我们可以部分部署，例如每次只取出集群的20%进行升级。

- 回滚困难
### 滚定发布注意事项
- 滚动发布没有一个确定可行的环境。使用蓝绿部署，我们能够清晰地知道老版本是可行的，而使用滚动发布，我们无法确定。

- 修改了现有的环境。

- 回滚困难。举个例子，在某一次发布中，我们需要更新100个实例，每次更新10个实例，每次部署需要5分钟。当滚动发布到第80个实例时，发现了问题，需要回滚，这个回滚却是一个痛苦，并且漫长的过程。

- 有的时候，我们还可能对系统进行动态伸缩，如果部署期间，系统自动扩容/缩容了，我们还需判断到底哪个节点使用的是哪个代码。尽管有一些自动化的运维工具，但是依然令人心惊胆战。

- 因为是逐步更新，那么我们在上线代码的时候，就会短暂出现新老版本不一致的情况，如果对上线要求较高的场景，那么就需要考虑如何做好兼容的问题。
## 灰度发布(Grayscale release) /金丝雀部署(Canary deployment)
灰度发布, 也叫金丝雀发布。是指在黑与白之间，能够平滑过渡的一种发布方式。AB test就是一种灰度发布方式，让一部分用户继续用A，一部分用户开始用B，如果用户对B没有什么反对意见，那么逐步扩大范围，把所有用户都迁移到B上面来。灰度发布可以保证整体系统的稳定，在初始灰度的时候就可以发现、调整问题，以保证其影响度，而我们平常所说的金丝雀部署也就是灰度发布的一种方式。

具体到服务器上, 实际操作中还可以做更多控制，譬如说，给最初更新的10台服务器设置较低的权重、控制发送给这10台服务器的请求数，然后逐渐提高权重、增加请求数。一种平滑过渡的思路, 这个控制叫做“流量切分”。

> 17世纪，英国矿井工人发现，金丝雀对瓦斯这种气体十分敏感。空气中哪怕有极其微量的瓦斯，金丝雀也会停止歌唱；而当瓦斯含量超过一定限度时，虽然鲁钝的人类毫无察觉，金丝雀却早已毒发身亡。当时在采矿设备相对简陋的条件下，工人们每次下井都会带上一只金丝雀作为“瓦斯检测指标”，以便在危险状况下紧急撤离。

过程:

- 准备好部署各个阶段的工件，包括：构建工件，测试脚本，配置文件和部署清单文件。
- 将“金丝雀”服务器部署进服务器中, 测试。
- 从负载均衡列表中移除掉“金丝雀”服务器。
- 升级“金丝雀”应用（排掉原有流量并进行部署）。
- 对应用进行自动化测试。
- 将“金丝雀”服务器重新添加到负载均衡列表中（连通性和健康检查）。
- 如果“金丝雀”在线使用测试成功，升级剩余的其他服务器。（否则就回滚）


## A/B测试

A/B测试和蓝绿发布、滚动发布以及金丝雀发布，完全是两回事。

蓝绿发布、滚动发布和金丝雀是发布策略，目标是确保新上线的系统稳定，关注的是新系统的BUG、隐患。

A/B测试是效果测试，同一时间有多个版本的服务对外服务，这些服务都是经过足够测试，达到了上线标准的服务，有差异但是没有新旧之分（它们上线时可能采用了蓝绿部署的方式）。

**A/B测试关注的是不同版本的服务的实际效果，譬如说转化率、订单情况等。**

A/B测试时，线上同时运行多个版本的服务，这些服务通常会有一些体验上的差异，譬如说页面样式、颜色、操作流程不同。相关人员通过分析各个版本服务的实际效果，选出效果最好的版本。