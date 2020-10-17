# BE_A
*api ball check container*

payload that spectator sent : 
{"ball_container":{"id":1, "ball_container_size":4,"current_ball_in_container":3}}

example curl for running locally :
curl -X POST -d '{"ball_container":{"id":1, "ball_container_size":4,"current_ball_in_container":3}}' http://localhost:7789/ball-container-check

curl -X POST -d '{"ball_container":{"id":1, "ball_container_size":3,"current_ball_in_container":1}}' http://localhost:7789/ball-container-check