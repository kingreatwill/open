

## Docker镜像的导入导出，用于迁移、备份、升级等场景
### save 
示例
docker save -o nginx.tar nginx:latest
或
docker save > nginx.tar nginx:latest

### load

示例
docker load -i nginx.tar
或
docker load < nginx.tar

### export
示例
docker export -o nginx-test.tar nginx-test(容器)


### import

示例
docker import nginx-test.tar nginx:imp
或
cat nginx-test.tar | docker import - nginx:imp

## 区别

- export命令导出的tar文件略小于save命令导出的
- export命令是从容器（container）中导出tar文件，而save命令则是从镜像（images）中导出
- 基于第二点，export导出的文件再import回去时，无法保留镜像所有历史（即每一层layer信息，不熟悉的可以去看Dockerfile），不能进行回滚操作；而save是依据镜像来的，所以导入时可以完整保留下每一层layer信息。如下图所示，nginx:latest是save导出load导入的，nginx:imp是export导出import导入的。