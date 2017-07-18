FROM jdk:8u131

RUN apt-get update && \
apt-get install wget && \
wget http://download.jboss.org/wildfly/10.1.0.Final/wildfly-10.1.0.Final.tar.gz && \
tar -xzvf wildfly-10.1.0.Final.tar.gz

ENTRYPOINT ["bash", "wildfly-10.1.0.Final/bin/standalone.sh", "-b=0.0.0.0"]
