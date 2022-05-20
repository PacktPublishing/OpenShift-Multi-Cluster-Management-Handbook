Discussing OpenShift Networking Model
The Container Network Interface (CNI) is a common interface between the network provider and the container execution and is implemented as network plug-ins. The CNI provides the specification for plug-ins to configure network interfaces inside containers. Plug-ins written to the specification allow different network providers to control the OpenShift cluster network.

The SDN uses CNI plug-ins to create Linux namespaces to partition the usage of resources and processes on physical and virtual hosts. This implementation allows containers inside pods to share network resources, such as devices, IP stacks, firewall rules, and routing tables. The SDN allocates a unique routable IP to each pod so that you can access the pod from any other service in the same network.

Some common CNI plug-ins used in OpenShift are:

OpenShift SDN

OVN-Kubernetes

Kuryr

In OpenShift 4.6, both OpenShift SDN and OVN-Kubernetes are the default network providers.

The OpenShift SDN network provider uses Open vSwitch (OVS) to connect pods on the same node and Virtual Extensible LAN (VXLAN) tunneling to connect nodes. OVN-Kubernetes uses Open Virtual Network (OVN) to manage the cluster network. OVN extends OVS with virtual network abstractions. Kuryr provides networking through the Neutron and Octavia Red Hat OpenStack Platform services.



# Link importante para o capitulo 9 network

https://role.rhu.redhat.com/rol-rhu/app/courses/do280-4.6/pages/ch05

https://kubernetes.io/docs/concepts/cluster-administration/networking/

https://github.com/ovn-org/ovn-kubernetes


    