FROM java8

COPY h2* /data/
EXPOSE "8082"

VOLUME /db
WORKDIR /db

ENTRYPOINT ["java", "-jar", "/data/h2-1.3.173.jar", "-webAllowOthers"]