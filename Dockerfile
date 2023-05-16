# from a basic image
# from an empty image, eg: from scratch
FROM ubuntu

# ENV: environment variables
ENV MY_SERVICE_PORT=80
ENV MY_SERVICE_PORT1=80
ENV MY_SERVICE_PORT2=80
ENV MY_SERVICE_PORT3=80

# organize project by labeling
LABEL multi.label1="value1" multi.label2="value2" other="value3"

# RUN: execute command, like: RUN apt-get update && apt-get install

# CMD: container executed commands with parameters, like: CMD ["executable", "param1", "param2"]

# COPY: copy resource files to target destination( copy support local files only, and do not untar files )
# Eg( multi-stage build ): COPY --from=build /bin/project /bin/project

# ADD: copy resource files to target destination( resources files could be files, directory, or url address )
# wildcard supported, eg: ADD check* /testdir/
# try not to use `ADD URL dest`, it is recommended to use `curl` or `wget && untar`
ADD bin/amd64/httpserver /httpserver

# EXPOSE: release port
EXPOSE 80

# ENTRYPOINT: image entrypoint command
# CMD: specifies what command to run within the container
ENTRYPOINT /httpserver
