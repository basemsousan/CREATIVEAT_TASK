# CREATIVEAT_TASK

# Background 
Provides a simple API written in Go. The API is designed to read mocked data (JSON based).

# API Endopoint

-   GET /employees

-   GET /metrics

# API GETs with curl 

-   curl -v http://localhost:8081/employees

-   curl -v http://localhost:8081/metrics

# Deploying the api on top of K8S

Pull the image from Docker Hub using the command bellow

-   docker pull basem85/restapi:firsttry

List the innstalled images on your system using the command bellow

-   docker images

you should see something like bellow 

![image](https://user-images.githubusercontent.com/26410486/151260478-841831ad-75ea-42aa-ab05-5b3e555adbb2.png)


-   kubectl create deployment restapi --image=basem85/restapi:firsttry

-   kubectl expose deployment restapi --type=NodePort --port=8081

-   kubectl port-forward service/restapi 7080:8081

API is now available at 

http://localhost:7080/employees

http://localhost:7080/metrics



