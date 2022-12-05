# keeper-active-users

### Server for storing active IP users list

    git clone git@github.com:Kolyan4ik99/keeper-active-users.git
    cd keeper-active-users
    make

Server listen:

    127.0.0.1:8080

For add or update expiration date of your ip address in the list of users, send 'GET' request:

    curl 127.0.0.1:8080/user/ping
    Response: {}

For get all active users:

    curl 127.0.0.1:8080/admin/users
    Response: [{"ip_address":"127.0.0.1","since":1670270078}]

Every user will be deleted after expiration date. The current one is 30 minute

