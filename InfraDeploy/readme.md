## Infra Service Deploy

[TOC]

------


### Install

#### Mysql 
        1.repo : http://repo.mysql.com
        2.command 
                rpm
                yum
                
#### Redis


#### Neo4j

        cd /tmp
        wget http://debian.neo4j.org/neotechnology.gpg.key
        rpm --import neotechnology.gpg.key
        Then, you'll want to add our yum repo to /etc/yum.repos.d/neo4j.repo:
        
        cat <<EOF>  /etc/yum.repos.d/neo4j.repo
        [neo4j]
        name=Neo4j Yum Repo
        baseurl=http://yum.neo4j.org/stable
        enabled=1
        gpgcheck=1
        EOF
        
        yum install neo4j



### Control

#### Command
        supervisorctl
        service
        

