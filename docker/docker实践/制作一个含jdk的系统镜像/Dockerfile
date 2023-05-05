FROM ubuntu
MAINTAINER me<xxx@qq.com>
ENV MYPATH /usr/local
WORKDIR $MYPATH
COPY README.txt $MYPATH
ADD jdk-8u151-linux-x64.tar.gz $MYPATH
RUN apt-get update -y
RUN apt-get install -y net-tools

ENV JAVA_HOME $MYPATH/jdk1.8.0_151
ENV CLASSPATH $JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar
ENV PATH $PATH:$JAVA_HOME/bin

VOLUME ["/usr/local/javafile"]
EXPOSE 8888
CMD /bin/bash
