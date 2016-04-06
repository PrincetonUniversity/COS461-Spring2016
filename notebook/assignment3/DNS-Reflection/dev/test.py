import sys
import threading
from scapy.all import *
from scapy.all import sniff as scasniff


class PacketHandler:
        
    def __init__(self, intf, mac):
        self.intf = intf
        self.mac = mac
        
    def start(self):
        t = threading.Thread(target = self.sniff)
        t.start()
    
    def incoming(self, pkt):
        return pkt[Ether].dst == self.mac


    def handle_packet(self, pkt):
        #TODO: compute and print the number of ping and DNS replies
        #      received in each minutes. (See the writeup for more
        #      details.)
        print "received packet"  

    def sniff(self):
        scasniff(iface=self.intf, prn = self.handle_packet,
                  lfilter = self.incoming) 
    
if __name__ == "__main__":
    if len(sys.argv) < 3:
        print "usage: python test.py intf mac"
        sys.exit(0)
    
    intf = sys.argv[1]
    mac = sys.argv[2]
    handler = PacketHandler(intf, mac)
    handler.start() 

