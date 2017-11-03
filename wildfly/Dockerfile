FROM jdk:8u151

RUN apt-get update && \
apt-get install wget && \
wget http://download.jboss.org/wildfly/11.0.0.Final/wildfly-11.0.0.Final.tar.gz && \
tar -xzvf wildfly-11.0.0.Final.tar.gz

ENTRYPOINT ["bash", "wildfly-11.0.0.Final/bin/standalone.sh", "-b=0.0.0.0"]
