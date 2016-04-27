#!/usr/bin/python

"""
Example to create a Mininet topology and connect it to the internet via NAT
through eth0 on the host.

Glen Gibb, February 2011

(slight modifications by BL, 5/13)
"""

from mininet.cli import CLI
from mininet.log import lg, info
from mininet.node import Node, OVSController, CPULimitedHost
from mininet.util import quietRun
from mininet.topo import Topo
from mininet.net import Mininet
from mininet.link import TCLink

#################################
class NetTopo(Topo):

    def __init__(self):
        super(NetTopo, self).__init__()

        client = self.addHost('client', ip = "192.168.0.1")
        proxy = self.addHost('proxy', ip = "192.168.0.2")
        
        s1 = self.addSwitch('s1')

        self.addLink(s1, client)
        self.addLink(s1, proxy)

def startNAT( root, inetIntf='eth0', subnet='192.168.0/24' ):
    """Start NAT/forwarding between Mininet and external network
    root: node to access iptables from
    inetIntf: interface for internet access
    subnet: Mininet subnet (default 10.0/8)="""

    # Identify the interface connecting to the mininet network
    localIntf =  root.defaultIntf()

    # Flush any currently active rules
    root.cmd( 'iptables -F' )
    root.cmd( 'iptables -t nat -F' )

    # Create default entries for unmatched traffic
    root.cmd( 'iptables -P INPUT ACCEPT' )
    root.cmd( 'iptables -P OUTPUT ACCEPT' )
    root.cmd( 'iptables -P FORWARD DROP' )

    # Configure NAT
    root.cmd( 'iptables -I FORWARD -i', localIntf, '-d', subnet, '-j DROP' )
    root.cmd( 'iptables -A FORWARD -i', localIntf, '-s', subnet, '-j ACCEPT' )
    root.cmd( 'iptables -A FORWARD -i', inetIntf, '-d', subnet, '-j ACCEPT' )
    root.cmd( 'iptables -t nat -A POSTROUTING -o ', inetIntf, '-j MASQUERADE' )

    # Instruct the kernel to perform forwarding
    root.cmd( 'sysctl net.ipv4.ip_forward=1' )

def stopNAT( root ):
    """Stop NAT/forwarding between Mininet and external network"""
    # Flush any currently active rules
    root.cmd( 'iptables -F' )
    root.cmd( 'iptables -t nat -F' )

    # Instruct the kernel to stop forwarding
    root.cmd( 'sysctl net.ipv4.ip_forward=0' )

def fixNetworkManager( root, intf ):
    """Prevent network-manager from messing with our interface,
       by specifying manual configuration in /etc/network/interfaces
       root: a node in the root namespace (for running commands)
       intf: interface name"""
    cfile = '/etc/network/interfaces'
    line = '\niface %s inet manual\n' % intf
    config = open( cfile ).read()
    if ( line ) not in config:
        print '*** Adding', line.strip(), 'to', cfile
        with open( cfile, 'a' ) as f:
            f.write( line )
    # Probably need to restart network-manager to be safe -
    # hopefully this won't disconnect you
    root.cmd( 'service network-manager restart' )

def connectToInternet( network, switch='s1', routerip='192.168.0.254', subnet='192.168.0/24'):
    """Connect the network to the internet
       switch: switch to connect to root namespace
       routerip: address for interface in root namespace
       subnet: Mininet subnet"""

    switch = network.get( switch )
    prefixLen = subnet.split( '/' )[ 1 ]
    routes = [ subnet ]  # host networks to route to

    # Create a node in root namespace
    router = Node( 'router', inNamespace=False )
    
    # Prevent network-manager from interfering with our interface
    fixNetworkManager( router, 'router-eth0' )

    # Create link between root NS and switch
    link = network.addLink( router, switch, delay='100ms')
    link.intf1.setIP( routerip, prefixLen )

    # Start network that now includes link to root namespace
    network.start()

    # Start NAT and establish forwarding
    startNAT( router )

    # Establish routes from end hosts
    for host in network.hosts:
        host.cmd( 'ip route flush root 0/0' )
        host.cmd( 'route add -net', subnet, 'dev', host.defaultIntf() )
        host.cmd( 'route add default gw', routerip )

    return router

if __name__ == '__main__':
    lg.setLogLevel( 'info')

    topo = NetTopo()
    net = Mininet(topo=topo, link=TCLink, host=CPULimitedHost, controller = OVSController)
    # Configure and start NATted connectivity
    router = connectToInternet(net)
    print "*** Hosts are running and should have internet connectivity"
    print "*** Type 'exit' or control-D to shut down network"
    CLI(net)
    # Shut down NAT
    stopNAT(router)
    net.stop()
