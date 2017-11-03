# etcd backup operator

## Overview

etcd backup operator backups the data of a etcd cluster running on [Kubernetes][Kube] to a remote storage such as AWS [S3][s3].

## Getting Started

Try out etcd backup operator by running it on Kubernetes and then create a `EtcdBackup` Custom Resource which contains the targeting etcd cluster and S3 backup config; the etcd backup operator automatically picks up the `EtcdBackup` Custom Resource, retrieves etcd snapshot, and then saves it to S3.
>Note: The demo uses the `default` namespace.

Prerequisites: 
* A running etcd cluster named `example-etcd-cluster`. See [instructions][etcd_cluster_deploy] to deploy it.

### Deploy etcd backup operator

Create a deployment of etcd backup operator:
> Note: etcd backup operator creates EtcdBackup CRD automatically

```sh
$ kubectl create -f example/etcd-backup-operator/deployment.yaml
$ kubectl get pod
NAME                                    READY     STATUS    RESTARTS   AGE
etcd-backup-operator-1102130733-hhgt7   1/1       Running   0          3s
```

### Setup AWS Secret

Create a Kubernetes secret that contains aws config/credential;
the secret will be used later to save etcd backup into S3.

1. Verify that the local aws config and credentials files exist:

    ```sh
    $ cat $AWS_DIR/credentials
    [default]
    aws_access_key_id = XXX
    aws_secret_access_key = XXX

    $ cat $AWS_DIR/config
    [default]
    region = <region>
    ```

2. Create secret `aws`:

    ```
    kubectl create secret generic aws --from-file=$AWS_DIR/credentials --from-file=$AWS_DIR/config
    ```

### Create EtcdBackup CR

Create EtcdBackup CR:
>Note: this example uses S3 Bucket "mybucket" and k8s secret "aws"

```sh
sed -e 's/<s3-bucket-name>/mybucket/g' \
    -e 's/<aws-secret>/aws/g' \
    example/etcd-backup-operator/backup_cr.yaml \
    | kubectl create -f -
```

### Verify status

Check the `status` section of the `EtcdBackup` CR:

```
$ kubectl get EtcdBackup example-etcd-cluster-backup -o yaml
apiVersion: etcd.database.coreos.com/v1beta2
kind: EtcdBackup
...
status:
  s3Path: mybucket/v1/default/example-etcd-cluster/3.1.8_0000000000000001_etcd.backup
  succeeded: true
```

* `s3Path` is the full S3 object path to the stored etcd backup. 

This demonstrates etcd backup operator's basic one time backup functionality.

### Cleanup

Delete the etcd-backup-operator deployment and the `EtcdBackup` CR.
> Note: Deleting the `EtcdBackup` CR won't delete the backup in S3.

```sh
kubectl delete etcdbackup example-etcd-cluster-backup
kubectl delete -f example/etcd-backup-operator/deployment.yaml
```

[Kube]:https://github.com/kubernetes/kubernetes
[s3]:https://aws.amazon.com/s3/
[etcd_cluster_deploy]:https://github.com/coreos/etcd-operator#create-and-destroy-an-etcd-cluster
[minikube]:https://github.com/kubernetes/minikube