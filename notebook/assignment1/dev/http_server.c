#include "proxy_parse.h"

char *badReqMsg   = "<html><head>\r\b<title>400 Bad Request</title>\r\n"\
	            "</head><body>\r\n<h1>Bad Request</h1>\r\n"\
	             "</body></html>\r\n";

char *notFoundMsg = "<html><head>\r\b<title>404 Not Found</title>\r\n"\
	            "</head><body>\r\n<h1>Not Found</h1>\r\n"\
	            "</body></html>\r\n";

char *notImpMsg   = "<html><head>\r\b<title>501 Not Implemented</title>\r\n"\
	            "</head><body>\r\n<h1>Not Found</h1>\r\n"\
	            "</body></html>\r\n";
	
int main(int argc, char * argv[])
{
	return 0;
}
