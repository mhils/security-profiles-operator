# create cluster
gcloud beta container --project "pse-shared" clusters create "mhils-test" --zone "europe-west4-b" --tier "standard" --no-enable-basic-auth --cluster-version "1.32.4-gke.1106006" --release-channel "regular" --machine-type "e2-standard-4" --image-type "COS_CONTAINERD" --disk-type "pd-balanced" --disk-size "100" --metadata disable-legacy-endpoints=true --num-nodes "1" --logging=SYSTEM,WORKLOAD --monitoring=SYSTEM,STORAGE,POD,DEPLOYMENT,STATEFULSET,DAEMONSET,HPA,JOBSET,CADVISOR,KUBELET,DCGM --enable-ip-alias --network "projects/pse-shared/global/networks/default" --subnetwork "projects/pse-shared/regions/europe-west4/subnetworks/default" --no-enable-intra-node-visibility --default-max-pods-per-node "110" --enable-ip-access --security-posture=standard --workload-vulnerability-scanning=disabled --no-enable-google-cloud-access --addons HorizontalPodAutoscaling,HttpLoadBalancing,GcePersistentDiskCsiDriver --enable-autoupgrade --enable-autorepair --max-surge-upgrade 1 --max-unavailable-upgrade 0 --binauthz-evaluation-mode=DISABLED --enable-managed-prometheus --enable-shielded-nodes --shielded-integrity-monitoring --no-shielded-secure-boot --node-locations "europe-west4-b"
# install cert-manager
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.17.2/cert-manager.yaml


# later: patch daemonset priortiy
kubectl apply -f rq.yaml
