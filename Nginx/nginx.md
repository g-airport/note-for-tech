proxy_set_header Host $host; \
proxy_set_header Http-Host $http_host; \
proxy_set_header X-Real-IP $remote_addr; \
proxy_set_header REMOTE-HOST $remote_addr; \
proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
