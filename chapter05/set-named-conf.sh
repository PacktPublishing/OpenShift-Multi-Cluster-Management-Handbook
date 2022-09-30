#!/bin/bash

# Adding bastion IP (192.168.1.200) in the listen-on port line
sudo sed -i 's/53 { 127.0.0.1;/53 { 127.0.0.1; 192.168.1.200;/g' /etc/named.conf

# Adding the new zones
sudo cat <<EOF >> /etc/named.conf
zone "ocp.hybridmycloud.com" IN {
type master;
file "/var/named/ocp.hybridmycloud.com.db";
allow-query { any; };
allow-transfer { none; };
allow-update { none; };
};

zone "1.168.192.in-addr.arpa" IN {
type master;
file "/var/named/ /var/named/1.168.192.in-addr.arpa";
allow-update { none; };
};
EOF

