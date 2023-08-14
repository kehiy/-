#include <stdio.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <string.h>

int main() {
    // init socket
    int socketFD = socket(AF_INET, SOCK_STREAM, 0);
    struct sockaddr_in address;

    // connect info
    address.sin_family = AF_INET;
    address.sin_port = htons(80);
    char* IP = "216.239.38.120";
    inet_pton(AF_INET, IP, &address.sin_addr.s_addr);

    // connect
    int result = connect(socketFD, &address, sizeof address);
    if(result == 0){
        printf("connection was successful\n");
    };

    char* message;
    message = "GET \\ HTTP/1.1\r\nHost:google.com\r\n\r\n";
    send(socketFD, message, strlen(message), 0);

    char buffer[1024];
    recv(socketFD, buffer, 1024, 0);
    printf("%s", buffer);

    return 0;
}