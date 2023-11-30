# Vitess

## Vitess with Docker
Make sure that Docker is installed and you Docker Daemon is up and running.

### Docker Image for vitess
`docker pull vitess/vttestserver:mysql80`

### Environments Variables
The docker image expects some of the environment variables to be set to function properly. The following table lists all the environment variables available along with their usages.
|Environment  |variable |Required Use |
| --- | :--: | :-- |
| KEYSPACES | yes | Specifies the names of the keyspaces to be created as a comma separated value. |
| NUM_SHARDS | yes | Specifies the number of shards in each keyspace. It is a comma separated value as well, read in conjunction with the KEYSPACES. | 
| PORT | 	yes	| The starting of the port addresses that vitess will use to register its components like vtgate, etc.| 
| MYSQL_MAX_CONNECTIONS	| no	Maximum number of connections that the MySQL instance will support. If unspecified, it defaults to 1000.| 
| MYSQL_BIND_HOST	| no	| Which host to bind the MySQL listener to. If unspecified, it defaults to 127.0.0.1.| 
| MYSQL_SERVER_VERSION | no	| MySQL server version to advertise. If unspecified, it defaults to 8.0.31-vitess or 5.7.9-vitess according to the version of vttestserver run.| 
| CHARSET	| no	| Default charset to use. If unspecified, it defaults to utf8mb4.| 
| FOREIGN_KEY_MODE | no	| This is to provide how to handle foreign key constraint in create/alter table. Valid values are: allow (default), disallow.| 
| ENABLE_ONLINE_DDL	| no	| Allow users to submit, review and control Online DDL. Valid values are: true (default), false.| 
| ENABLE_DIRECT_DDL	| no	| Allow users to submit direct DDL statements. Valid values are: true (default), false.| 
| PLANNER_VERSION	| no | Sets the default planner to use when the session has not changed it. Valid values are: Gen4, Gen4Greedy, Gen4Left2Right.| 
| TABLET_REFRESH_INTERVAL	| no	| Interval at which vtgate refreshes tablet information from topology server.| 

Environment variables in docker can be specified using the -e aka --env flag.

### Ports
The vtgate listens for MySQL connections on 3 + the PORT environment variable specified. i.e. if you specify PORT to be 33574,  
then vtgate will be listening to connections on 33577, on the MYSQL_BIND_HOST which defaults to localhost.  
But this port will be on the docker container side. To connect to vtgate externally from a MySQL client, you will need to publish that port as well and  
specify the MYSQL_BIND_HOST to 0.0.0.0. This can be done via the -p aka --publish flag to docker.  
For eg: adding -p 33577:33577 to the docker run command will publish the container's 33577 port to your local 33577 port, which can now be used to connect to the vtgate.

### Persistence
If you wish to keep the state of the test container across reboots, such as when running the vttestserver container as a database container in local application development environments, you may optionally pass the --persistent_mode flag, along with a --data_dir directory which is bound to a docker volume.  
Due to a bug, the --port argument must also be present for correct operation.

When running in this mode, underlying MySQL table schemas, their data, and the Vitess VSchema objects are persisted under the provided --data_dir.

### Example 1 (single instance)
An example command to run the docker image is as follows :
```
docker run --name=vttestserver \
  -p 33577:33577 \
  --health-cmd="mysqladmin ping -h127.0.0.1 -P33577" \
  --health-interval=5s \
  --health-timeout=2s \
  --health-retries=5 \
  -v vttestserver_data:/vt/vtdataroot/vitess \
  vitess/vttestserver:mysql80
  /vt/bin/vttestserver \
  --alsologtostderr \
  --data_dir=/vt/vtdataroot/vitess \
  --persistent_mode \
  --port=33574 \
  --mysql_bind_host=0.0.0.0 \
  --keyspaces=test,unsharded \
  --num_shards=2,1
```
### Example 2 (with shards)
An example command to run the docker image is as follows:

```
docker run --name=vttestserver \
  -p 33577:33577 \
  -e PORT=33574 \
  -e KEYSPACES=test,unsharded \
  -e NUM_SHARDS=2,1 \
  -e MYSQL_MAX_CONNECTIONS=70000 \
  -e MYSQL_BIND_HOST=0.0.0.0 \
  --health-cmd="mysqladmin ping -h127.0.0.1 -P33577" \
  --health-interval=5s \
  --health-timeout=2s \
  --health-retries=5 \
  vitess/vttestserver:mysql80
```
### Connect Vitess-DB from command line
Now, we can connect to the vtgate from a MySQL client as follows:
`mysql --host 127.0.0.1 --port 33577 --user "root"`

# Simple Test for Vitess Installation
After the connection to Vitess-Server you may test the installation as follows:  
```
CREATE DATABASE <databasename>;
USE <databasename>;
```
```
CREATE TABLE Customers (
    ID int NOT NULL,
    LastName varchar(255) NOT NULL,
    FirstName varchar(255),
    Age int,
    Address varchar(255),
    PostalCode int,
    City varchar(255),
    CONSTRAINT PK_Customers PRIMARY KEY (ID)
);
```
```
SHOW TABLES
```
```
INSERT INTO Customers (ID, LastName, FirstName, Age, Address, PostalCode, City) VALUES (1, 'Cardinal', 'Tom B. Erichsen', 19, 'Skagen 21', 4006, 'Stavanger'); VALUES (2, 'Maker', 'Kim', 23, 'Mainstreet 37', 50012, 'Meansville');
```
```
SELECT * FROM Customers
```

# References
[https://vitess.io/docs/18.0/overview/](https://vitess.io/docs/18.0/overview/)
