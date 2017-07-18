FROM ubuntu

RUN apt-get update -y && \
apt-get install -y wget && \
wget -c --header "Cookie: oraclelicense=accept-securebackup-cookie" http://download.oracle.com/otn-pub/java/jdk/8u131-b11/d54c1d3a095b4ff2b6607d096fa80163/jdk-8u131-linux-x64.tar.gz &&\
tar -xzf jdk-8u131-linux-x64.tar.gz && \
rm -rf jdk-8u131-linux-x64.tar.gz && \
mkdir -p /usr/lib/jvm && \
mv jdk1.8.0_131 /usr/lib/jvm && \
update-alternatives --install /usr/bin/java java /usr/lib/jvm/jdk1.8.0_131/jre/bin/java 1 && \
update-alternatives --install /usr/bin/javac javac /usr/lib/jvm/jdk1.8.0_131/bin/javac 1

CMD echo