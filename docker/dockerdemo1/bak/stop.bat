# stopped alle container des Benutzers
docker ps -q | xargs -L1 docker stop
#
###### EOF ######