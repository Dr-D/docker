FROM ubuntu

RUN apt-get update -y && \
apt-get install -y wget && \
wget -c --header "Cookie: oraclelicense=accept-securebackup-cookie" http://download.oracle.com/otn-pub/java/jdk/8u152-b16/aa0333dd3019491ca4f6ddbe78cdb6d0/jdk-8u152-linux-x64.tar.gz &&\
tar -xzf jdk-8u152-linux-x64.tar.gz && \
rm -rf jdk-8u152-linux-x64.tar.gz && \
mkdir -p /usr/lib/jvm && \
mv jdk1.8.0_152 /usr/lib/jvm && \
update-alternatives --install /usr/bin/java java /usr/lib/jvm/jdk1.8.0_152/jre/bin/java 1 && \
update-alternatives --install /usr/bin/javac javac /usr/lib/jvm/jdk1.8.0_152/bin/javac 1

CMD echo