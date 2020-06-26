FROM registry.svc.ci.openshift.org/openshift/release:golang-1.13 AS builder
WORKDIR /go/src/github.com/openshift/azure-disk-csi-driver-operator
COPY . .
RUN make

FROM registry.svc.ci.openshift.org/openshift/origin-v4.0:base
COPY --from=builder /go/src/github.com/openshift/azure-disk-csi-driver-operator/azure-disk-csi-driver-operator /usr/bin/
COPY bundle /bundle
ENTRYPOINT ["/usr/bin/azure-disk-csi-driver-operator"]
LABEL com.redhat.delivery.appregistry=true
LABEL io.k8s.display-name="OpenShift azure-disk-csi-driver-" \
      io.k8s.description="The azure-disk-csi-driver-operator installs and maintains the Azure Disk CSI Driver on a cluster."
