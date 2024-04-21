# https://blog.51cto.com/u_93011/10139111
# pvc 存在数据来，不能修改 第一次与第二次的初始化参数不一样。因为会对数据产生影响
#  nfs 安装  rpcbind  nfs-utils
# 服务启动  systemctl enable rpcbind     nfs-utils 需要启动这个 nfs-utils.service 网上有不是这个文件的。
# 配置 /etc/export
# 我们也可以不实用动态制备 可以先创建 pv 然后使用 volumeClaimTemplates 不创建  StorageClass
# kubectl get  pod  -n nfs-client-provisioner
#  Deployment  StatefulSet 使用  secretKeyRef 存在bug
#  StorageClass volumeClaimTemplates 可以动态声明 pvc,pv,他们可以使用
# k8s-sigs.io/nfs-subdir-external-provisioner 驱动制作 。到 GitHub nfs-subdir-external-provisioner 仓库克隆仓库。
# 然后给 查看  跟目录 .github/workflows  release.yml 文件  on 为 action的出发规则。 文件是把镜像推送到  quay.io 仓库
# 制备好 镜像。使用仓库下面的  deploy 文件 里面的  kustomization.yaml 部署驱动（class.yaml是一个测试的 storageClass）可以只运行
# rbac.yaml deployment.yaml  objects 文件是拆细的对象，可以学习
# on:
#  push:
#    tags:
#      - 'gh-v*.*.*'

kubectl get cm/mysql-cm -o yaml
kubectl create configmap  mysql-cm --from-file ./mysql.sh