#include <stdio.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <string.h>

int main() {
    int socketFD = socket(AF_INET, SOCK_STREAM, 0);

    struct sockaddr_in address;

    // connect info
    address.sin_family = AF_INET;
    address.sin_port = htons(2000);
    address.sin_addr.s_addr = INADDR_ANY;

    int result = bind(socketFD, &address, sizeof(address));
    if(result == 0) {
        printf("server was bin successfully");
    };

    int result = listen(socketFD, 10);
    if(result == 0){
        printf("listening was successful");
    };

    //! THIS LINE IS WRONG
    // accept(socketFD, &address, sizeof(address));


    return 0;
}