K8S工程师必备问题排查手册（作者:ZHDYA)
一、Pod相关问题及排查: 
1.Pod无法启动，如何查找原因?
使用kubectl describe pod [pod_name] -n [namespace_name]命令查看该Pod 的状态信息，检查容器的状态和事件信息，判断是否出现问题。
·使用kubectl logs [pod_name] -n [namespace_name]命令查看该Pod容器的日志信息，判断是否有错误或异常信息。
使用kubectl get events --field-selector involvedObject.name=[pod_name] -n [namespace_name]命令查看该Pod相关的事件信息，判断是否有异常事件发生。
2.Pod无法连接到其他服务，如何排查?
使用kubectl exec -it [pod_name] -n [namespace_name] -- /bin/bash命令进入该Pod 所在的容器，尝试使用 ping 或 telnet.等命令测试与其他服务的网烙连接情况。
便用kubectl describe pod [pod_name] -n [namespace_name]命令检查Pod 的NetworkPo1icy 配置，判断是否阻止了该Pod访问其他服务。
·使用kubectl describe service [service_name] -n [namespace_name]命令检查目标服务的配置和状态信息，判断是否存在故障。
3.Pod运行缓慢或异常，如何排查?
便用kubectl top pod [pod_name] -n [namespace_name]命令查看该Pod 的CPU和内存便用情况，判断是否存在性能瓶颈。
使用kubectl exec -it [pod_name] -n [namespace_name] -- /bin/bash命令进入该Pod所在的容器，使用 top 或 htop 命令查看容器内部进程的CPU和内存使用情况，找出可能存在的瓶颈。
。使用kubectl logs [pod_name] -n [namespace_name]命令查看该Pod容器的日志信息，寻找可能的错误或异常信息。
4.Pod无法被调度到节点上运行，如何排查?
使用kubectl describe pod [pod_name] -n [namespace_name]命令查看Pod的调度情况，判断是否存在资源不足、调度策略等问题。
使用kubectl get nodes和kubectl describe node [node_name]命令查看所有节点的资源使用情况，判断是否存在节点资源不足或故障的情况。使用kubectl describe pod [pod_name] -n [namespace_name]命令检查Pod 所需的标签和注释，以及节点的标签和注释，判断是否匹配
5.Pod状态一直是Pending，怎么办?
·查看该Pod的事件信息:kubectl describe pod <pod-name>
·查看该节点资源利用率是否过高:kubectl top node
·如果是调度问题，可以通过以下方式解决:
。确保有足够的节点资源满足该Pod调度需求
。检查该节点的 taints 和tolerations是否与Pod 的selector 匹配。
。调整Pod的调度策略，如使用NodeSelector、Affinity等
6.Pod无法访问外部服务，怎么办?
。查看Pod中的DNS配置是否正确
·检查Pod所在的命名空间中是否存在Service服务
·确认该Pod是否具有网络访问权限
·查看Pod所在的节点是否有对外的访问权限
·检查网络策略是否阻止了Pod对外的访问
7.Pod启动后立即退出，怎么办?
。查看该Pod的事件信息:kubectl describe pod <pod-name>
。查看该Pod的日志: kubectl logs <pod-name>
·检查容器镜像是否正确、环境变量是否正确、入口脚本是否正常
·尝试在本地便用相同的镜像运行该容器，查看是否有报错信息，如执行docker run <image-name>
8.Pod启动后无法正确运行应用程序，怎么办?
。查看Pod中的应用程序日志: kubectl logs <pod-name>。
查看该Pod 的事件信息:kubectl describe pod <pod-name>
·检查应用程序的配置文件是否正确
·检查应用程序的依赖是否正常
尝试在本地使用相同的镜像运行该容器，查看是否有报错信息，如执行docker run <image-name>
·确认该应用程序是否与Pod的资源限制相符
9.Kubernetes集群中的Service不可访问，怎么办?
·检查Service的定义是否正确
·检查endpoint是否正确生成·检查网络插件配置是否正确
·确保防火墙配置允许Service对外开放
二、Node相关问题及排查:
1.Node状态异常，如何排查?
·使用kubectl get nodes命令查看集群中所有节点的状态和信息，判断是否存在故障。
·使用kubectl describe node [node_name]命令查看目标节点的详细信息，包括CPU、内存、磁盘够硬件资源的使用情况，判断是否存在性能瓶颈
使用 kubectl get pods -o wide --all-namespaces 命令查看集群中所有Pod 的状态信息，判断是否有Pod运行在目栎节点上导致资源紧张。
2.Node 上运行的Pod无法访问网络，如何排查?
·使用kubectl describe node [node_name]命令查看目标节点的信息，检查节点是否正常连接到网烙。
·使用kubectl describe pod [pod_name] -n [namespace_name]命令查看Pod所运行的节点信息，判断是否因为节点状态异常导数网络访问失败。
·使用kubectl logs [pod_name] -n [namespace_name]命令查看Pod容器的日志信息，寻找可能的错误或异常信息。
3.Node 上的 Pod无法访问存储，如何排查?
便用kubectl describe pod [pod_name] -n [namespace_name]命令检查Pod的 volumes配置信息，判断是否存在存储挂载失败的情况。
使用kubectl exec -it [pod_name] -n [namespace_name] -- /bin/bash命令进入Pod所在的容器，尝试使用ls和cat等命令访问挂载的文件系统，判断是否存在读写错误。
使用kubectl describe persistentvolumeclaim [pvc_name] -n [namespace_name]命令查看相关PVC配置和状态信息，判新是否存在故障
4.存储卷挂载失败，如何处理?
。使用 kubectl describe pod [pod_name] -n [namespace_name]命令检查Pod的 volumes配置信息，判断是否存在存储卷定义错误.
·使用 kubectl describe persistentvolumeclaim [pvc_name] -n [namespace_name]命令检查PVC 的状态和信息，判断是否存在存储配额不足或存储资源故障等原因。
·如果是NFS或Ceph等网络存储，需要确认网络连接是否正常，以及存储服务器的服务是否正常。
5. Node节点加入 Kubernetes集群后无法被调度，怎么办?
·检查该节点的 taints 和 tolerations 是否与Pod的 selector匹配
·检查该节点的资源便用情况是否满足Pod的调度要求
·确保该节点与Kubernetes API server的连接正常
6.Kubernetes集群中的 PersistentVolume挂载失败，怎么办?
·检查PersistentVolume和Pod之间的匹配关系是否正确
·检查PersistentVolumeClaim 中的 storageClassName是否与PersistentVolume的 storageClassName 匹配
·检查节点存储配置和PersistentVolume 的定义是否正确
·白动供给层面的权限是否已经给到位
三、集群层面问题及排查:
1.集群中很多Pod运行缓慢，如何排查?
·便用kubectl top pod -n [namespace_name]命令查看所有Pod的CPU和内存使用情况，判断是否存在资源瓶颈。
使用kubectl get nodes和kubectl describe node [node_name]命令查看所有节点的资源使用情况，判断是否存在单个节点资源紧张的情况。
·使用kubectl logs [pod_name] -n [namespace_name]命令查看Pod容器的日志信息，寻找可能的错误或异常信息。
2.集群中某个服务不可用，如何排查?
。使用kubectl get pods -n [namespace_name]命令查看相关服务的所有Pod的状态信息，判断是否存在故障。
使用kubectl describe pod [pod_name〕 -n [namespace_name]命令检查Pod的网络连接和存储访问等问题，寻找故障原因。
·便用kubectl describe service [service_name] -n [namespace_name]命令查看服务的配置和状态信息，判断是否存在故障
3.集群中的Node 和Pod不平衡，如何排查?
使用kubectl get nodes和kubectl get pods -o wide --all-namespaces 命令查看所有Node和Pod的状态信息，判断是否存在分布不均的情况.
·使用 kubectl top pod -n [namespace_name]命令查看所有Pod的CPU和内存便用情况，判断是否存在资源瓶颈导致Pod分布不均,
使用kubectl describe pod [pod_name] -n [namespace_name]命令查看Pod 所运行的节点信忌，并使用kubectl describe node [node_name]命令查看相关节点的状态信息，判断是否存在节点不平衡的情况。
。使用kubectl describe pod / node [node_name]查看当前Pod / Node上是否有相关的亲和或反亲和策略导致固定调度。
4.使群中某个节点宕机,如何处理?
●使用kubectl get nodes 命令检查节点状态，找到异常节点。
●使用kubectl drain [node_name] --ignore-daemonsets 命令将节点上的Pod驱逐出去，并将其部署到其他节点上。添加--ignore-daemonsets 参数可以忽略DaemonSet资源。
●如果需要对节点进行维护或替换硬件，则便用kubectl delete node [node_name]命令删除该节点。此时该节点上运行的Pod会自动调度到其他节点上。
5. Kubernetes API Server不可用，如何排查?
●使用kubectl cluster -info命令查看集群状态，判断是否存在API Server不可用的情况。
●便用kubectl version命令查看集群版本，确认Kubernetes API Server和kubelet版本是否匹配。
●使用systemctl status kube-apiserver命令检查API Server运行状态，确认是否存在故障或错误。
●结合apiServer所在的节点查看系统层面的日志，进一步 定位问题点。
6. Kubernetes命令执行失败，怎么办?
●检查Kubernetes API server是否可用: kubectl cluster-info
●检查当前用户对集群的权限是否足够: kubectl auth can-i <verb> <resource>
●检查kubeconfig文件中的登录信息是否正确: kubectl config view
7. Kubernetes master节点不可用，怎么办?
●检查kube-apiserver, kube-scheduler, kube-controller-manager是否都在运行状态
，检查etcd存储系统是否可用
, 尝试重新启动master节点上的kubelet和容器运行时
8. Kubernetes集群绕过了LoadBalancer,直接访问Pod,怎么办?
●检查Service和Pod的通信是否使用了ClusterIP类型的Service
确认该Service的selector是否匹配到了正确的Pod
9. Kubernetes集群中的Deployment自动更新失败，怎么办?
●检查更新策略是否设置正确，如rollingUpdate或recreate
●检查Kubernetes API server和kubelet之间的连接是否正常
●检查Pod的定义是否正确
10. Kubernetes集群中的状态检查错误，怎么办?
●检查节点日志和事件信息，并确认错误类型
●确认该状态检查是否与kubelet的版本兼容
●尝试升级kubelet和容器运行时等组件
11. Kubernetes集群中的授权配置有误，怎么办?
●检查RoleBinding和ClusterRoleBinding定义是否正确
●检查用户或服务账号所绑定的角色是否正确
●检查kubeconfig文件中的用户和访问权限是否正确
12. Kubernetes集群无法连接etcd存储系统，怎么办?
●检查etcd存储系统是否正常运行
●检查kube apiserver配置文件中etcd的连接信息是否正确
●尝试手动连接etcd集群,如执行etcdctl cluster-health

